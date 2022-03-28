package gingen

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"google.golang.org/genproto/googleapis/api/annotations"
	//gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/reflect/protoreflect"

	// gengo "github.com/jiandahao/golanger/pkg/generator/gingen/gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// ProtocPlugin is a plugin used to generate service implementaion
type ProtocPlugin struct {
	gen *protogen.Plugin
}

// NewProtocPlugin new protoc plugin
func NewProtocPlugin() func(gen *protogen.Plugin) error {
	return func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		plugin := &ProtocPlugin{gen: gen}

		return plugin.Execute()
	}
}

// Execute starts to generate file.
func (p *ProtocPlugin) Execute() error {
	for _, file := range p.gen.Files {
		p.generateFile(p.newFileInfo(file))
		//	gengo.GenerateFile(p.gen, file)
	}

	return nil
}

type fileInfo struct {
	*protogen.File

	allEnums    []*protogen.Enum
	allMessages []*protogen.Message
	// allExtensions []*extensionInfo
}

func (p *ProtocPlugin) newFileInfo(file *protogen.File) *fileInfo {
	f := &fileInfo{File: file}
	// Collect all enums, messages, and extensions in "flattened ordering".
	var walkMessages func([]*protogen.Message, func(*protogen.Message))
	walkMessages = func(messages []*protogen.Message, f func(*protogen.Message)) {
		for _, m := range messages {
			f(m)
			walkMessages(m.Messages, f)
		}
	}

	f.allEnums = append(f.allEnums, f.Enums...)
	f.allMessages = append(f.allMessages, f.Messages...)

	walkMessages(f.Messages, func(m *protogen.Message) {
		f.allEnums = append(f.allEnums, m.Enums...)
		f.allMessages = append(f.allMessages, m.Messages...)
	})
	return f
}

func (p *ProtocPlugin) generateFile(file *fileInfo) {
	filename := file.GeneratedFilenamePrefix + "_gin.pb.go"
	generatedFile := p.gen.NewGeneratedFile(filename, file.GoImportPath)
	generatedFile.P(`// Code generated. DO NOT EDIT.

	package ` + file.GoPackageName)

	// Using QualifiedGoIdent to make referenced Packages to be automatically imported.
	generatedFile.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context"})
	generatedFile.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "github.com/gin-gonic/gin"})
	generatedFile.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes"})
	generatedFile.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status"})
	generatedFile.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"})

	p.genMessage(generatedFile, file.allMessages)
	p.genEnum(generatedFile, file.allEnums)

	for _, service := range file.Services {
		p.generateService(generatedFile, service)
	}
}

func (p *ProtocPlugin) generateService(generatedFile *protogen.GeneratedFile, service *protogen.Service) {
	opts := service.Desc.Options().(*descriptorpb.ServiceOptions)
	if opts.GetDeprecated() {
		generatedFile.P("Deprecated: Do not use.")
	}

	s := &ServiceInfo{
		ServiceName: service.GoName,
	}
	for _, method := range service.Methods {
		s.Methods = append(s.Methods, genMethod(generatedFile, method))
	}

	generatedFile.P(executeTemplate("serviceInterfaceTempl", serviceInterfaceTempl, s))
	generatedFile.P(executeTemplate("unimplementServerTempl", unimplementServerTempl, s))
	generatedFile.P(executeTemplate("serviceDecoratorTempl", serviceDecoratorTempl, s))
	generatedFile.P(executeTemplate("registerTempl", registerTempl, s))
}

func (p *ProtocPlugin) genEnum(generatedFile *protogen.GeneratedFile, enums []*protogen.Enum) {
	for _, enum := range enums {
		leadingComments := appendDeprecationSuffix(enum.Comments.Leading,
			enum.Desc.Options().(*descriptorpb.EnumOptions).GetDeprecated())
		generatedFile.P(leadingComments, "type ", enum.GoIdent, " int32")

		// Enum value constants.
		generatedFile.P("const (")
		for _, value := range enum.Values {
			generatedFile.Annotate(value.GoIdent.GoName, value.Location)
			leadingComments := appendDeprecationSuffix(value.Comments.Leading,
				value.Desc.Options().(*descriptorpb.EnumValueOptions).GetDeprecated())
			generatedFile.P(leadingComments,
				value.GoIdent, " ", enum.GoIdent, " = ", value.Desc.Number(),
				trailingComment(value.Comments.Trailing))
		}
		generatedFile.P(")")
		generatedFile.P()
	}
}

// appendDeprecationSuffix optionally appends a deprecation notice as a suffix.
func appendDeprecationSuffix(prefix protogen.Comments, deprecated bool) protogen.Comments {
	if !deprecated {
		return prefix
	}
	if prefix != "" {
		prefix += "\n"
	}
	return prefix + " Deprecated: Do not use.\n"
}

// trailingComment is like protogen.Comments, but lacks a trailing newline.
type trailingComment protogen.Comments

func (c trailingComment) String() string {
	s := strings.TrimSuffix(protogen.Comments(c).String(), "\n")
	if strings.Contains(s, "\n") {
		// We don't support multi-lined trailing comments as it is unclear
		// how to best render them in the generated code.
		return ""
	}
	return s
}

func (p *ProtocPlugin) genMessage(generatedFile *protogen.GeneratedFile, messages []*protogen.Message) {
	type structTag [][2]string
	for _, msg := range messages {
		leadingComments := msg.Comments.Leading.String()
		if leadingComments == "" {
			leadingComments = "\n"
		}
		generatedFile.P(leadingComments, "type ", msg.GoIdent.GoName, " struct {\n")
		for _, field := range msg.Fields {
			fieldName := field.GoName
			fieldType, isPointer := fieldGoType(generatedFile, field)
			if isPointer {
				fieldType = "*" + fieldType
			}

			var tags []string
			trailingComment := field.Comments.Trailing.String()
			tags = append(tags, fmt.Sprintf(`json:"%s,omitempty"`, field.Desc.TextName()))

			if trailingComment != "" {
				trailingComment = strings.TrimPrefix(trailingComment, "//")
				rawTags := strings.Split(trailingComment, " ")
				for _, rt := range rawTags {
					res := strings.Split(strings.TrimSpace(rt), ":")
					if len(res) == 2 {
						tags = append(tags, fmt.Sprintf(`%s:"%s"`, res[0], strings.TrimFunc(res[1], func(r rune) bool {
							return r == '"'
						})))
					}
				}
			}

			generatedFile.P(fmt.Sprintf("%s %s %s", fieldName, fieldType, fmt.Sprintf(" `%s`", strings.Join(tags, " "))))
		}
		generatedFile.P("}")
	}
}

// fieldGoType returns the Go type used for a field.
//
// If it returns pointer=true, the struct field is a pointer to the type.
func fieldGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := fieldGoType(g, field.Message.Fields[0])
		valType, _ := fieldGoType(g, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}

func genMethod(g *protogen.GeneratedFile, m *protogen.Method) *Method {
	// try to parse http rules
	rule, ok := proto.GetExtension(m.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
	if rule != nil && ok {
		method := &Method{
			Name:     m.GoName,
			Request:  g.QualifiedGoIdent(m.Input.GoIdent),  //m.Input.GoIdent.GoName,
			Response: g.QualifiedGoIdent(m.Output.GoIdent), //m.Output.GoIdent.GoName,
			Comments: m.Comments.Leading.String(),
		}

		method.HTTPRules = append(method.HTTPRules, extractHTTPRule(rule))
		for _, bind := range rule.AdditionalBindings {
			method.HTTPRules = append(method.HTTPRules, extractHTTPRule(bind))
		}

		return method
	}

	panic(fmt.Sprintf("no http rules found for method: %s", m.GoName))
}

func extractHTTPRule(rule *annotations.HttpRule) *HTTPRule {
	var method, path string
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		method = http.MethodGet
	case *annotations.HttpRule_Put:
		path = pattern.Put
		method = http.MethodPut
	case *annotations.HttpRule_Post:
		path = pattern.Post
		method = http.MethodPost
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		method = http.MethodDelete
	case *annotations.HttpRule_Patch:
		path = pattern.Patch
		method = http.MethodPatch
	case *annotations.HttpRule_Custom:
		path = pattern.Custom.Path
		method = pattern.Custom.Kind
	}

	var hasPathParam bool
	paths := strings.Split(path, "/")
	for idx, item := range paths {
		if len(item) > 0 && (item[0] == '{' && item[len(item)-1] == '}') {
			paths[idx] = ":" + item[1:len(item)-1]
			hasPathParam = true
		}
	}
	path = strings.Join(paths, "/")

	return &HTTPRule{Method: method, Path: path, HasPathParam: hasPathParam}
}

// ServiceInfo describes the service info.
type ServiceInfo struct {
	ServiceName string
	Methods     []*Method
}

// Method service's method
type Method struct {
	Name      string
	Request   string
	Response  string
	HTTPRules []*HTTPRule
	Comments  string
}

// HasPathParam returns true if there is at least one http rule contains path param.
func (m Method) HasPathParam() bool {
	for _, r := range m.HTTPRules {
		if r.HasPathParam {
			return true
		}
	}

	return false
}

// HTTPRule http rule
type HTTPRule struct {
	Method       string
	Path         string
	HasPathParam bool
}

// Message message
type Message struct {
	*protogen.Message
}

// Name returns the message's name
func (m *Message) Name() string {
	return m.GoIdent.GoName
}

func executeTemplate(name string, tmpl string, data interface{}) string {
	t := template.Must(template.New(name).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		panic(fmt.Sprintf("failed to render template: %v", err))
	}
	return buf.String()
}

var serviceInterfaceTempl = `
	// {{.ServiceName}}Server is the server API for {{.ServiceName}} service.
	type {{.ServiceName}}Server interface {
	{{- range .Methods}}
		{{.Comments -}}
		{{.Name}}(context.Context, *{{.Request}}) (*{{.Response}}, error)
	{{- end}}}
`

var unimplementServerTempl = `
	{{$serviceName := .ServiceName}}
	// Unimplemented{{$serviceName}}Server can be embedded to have forward compatible implementations.
	type Unimplemented{{$serviceName}}Server struct {}

	{{- range .Methods}}
	func (s *Unimplemented{{$serviceName}}Server) {{.Name}}(context.Context, *{{.Request}}) (*{{.Response}}, error) {
		return nil, status.Errorf(codes.Unimplemented, "method {{.Name}} not implemented")
	}
	{{ end }}		
	`

var serviceDecoratorTempl = `
	{{$serviceName := .ServiceName}}
	type default{{$serviceName}}Decorator struct{
		ss {{$serviceName}}Server
	}

	{{range .Methods}}
		{{$methodName := .Name}}
		{{$requestParamType := .Request}}
		{{range $index, $rule := .HTTPRules}}
			func (s default{{$serviceName}}Decorator) {{$methodName}}_{{$index}}(ctx *gin.Context){
				var req {{$requestParamType}}
				{{ if $rule.HasPathParam }}
				if err := ctx.ShouldBindUri(&req); err != nil {
					runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error())) 
					return
				}
				{{ end }}

				{{ if eq $rule.Method "GET" "DELETE" }}
				if err := ctx.ShouldBindQuery(&req); err != nil {
					runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error())) 
					return
				}
				{{else if eq $rule.Method "POST" "PUT" }}
				if err := ctx.ShouldBindJSON(&req); err != nil {
					runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error())) 
					return
				}
				{{else}}
				if err := ctx.ShouldBind(&req); err != nil {
					runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error())) 
					return
				}
				{{end}}

				resp, err := s.ss.{{$methodName}}(ctx, &req)
				if err != nil {
					runtime.HTTPError(ctx, err)
					return
				}
		
				runtime.ForwardResponseMessage(ctx, resp)	
			}
		{{end}}
	{{end}}
`

var registerTempl = `
	{{$serviceName := .ServiceName}}
	// Register{{$serviceName}}Server registers the http handlers for service {{$serviceName}} to "router".
	func Register{{$serviceName}}Server(router gin.IRouter, s {{$serviceName}}Server) {
		d := default{{$serviceName}}Decorator{ss: s}
		{{- range .Methods -}}
			{{ $methodName := .Name }}
			{{- range $index, $rule := .HTTPRules}}
				router.Handle("{{$rule.Method}}", "{{$rule.Path}}", d.{{$methodName}}_{{$index}})
			{{- end}}
		{{- end}}
	}
`

var clientTempl = `
	type {{$serviceName}}Client interface {
		{{- range .Methods}}
			{{.Comments -}}
			{{.Name}}(context.Context, *{{.Request}}) (*{{.Response}}, error)
		{{end}}
	}

	type {{$serviceName}}Client struct {
		cc *grpc.ClientConn
	}

	func NewAccountCreatorClient(cc *grpc.ClientConn) AccountCreatorClient {
		return &accountCreatorClient{cc}
	}
`
