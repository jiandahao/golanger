package gingen

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// ProtocPlugin is a plugin used to generate service implementation
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
func (p *ProtocPlugin) Execute() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	for _, file := range p.gen.Files {
		p.GenerateFile(NewFileInfo(file))
	}

	return nil
}

// GenerateFile generates file.
func (p *ProtocPlugin) GenerateFile(file *FileInfo) {
	filename := file.GeneratedFilenamePrefix + "_gin.pb.go"
	generatedFile := p.gen.NewGeneratedFile(filename, file.GoImportPath)
	generatedFile.P(`// Code generated. DO NOT EDIT.`)
	generatedFile.P()
	generatedFile.P(`package ` + file.GoPackageName)

	// Using QualifiedGoIdent to make referenced Packages to be automatically imported.
	importedPath := []string{
		"io",
		"fmt",
		"encoding/json",
		"strings",
		"io/ioutil",
		"net/http",
		"bytes",
		"net/url",
		"context",
		"mime/multipart",
		"github.com/gin-gonic/gin",
		"google.golang.org/grpc/codes",
		"google.golang.org/grpc/status",
		"github.com/jiandahao/golanger/pkg/generator/gingen/runtime",
	}

	for _, pkg := range importedPath {
		generatedFile.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(pkg)})
	}

	generatedFile.P(`
		var _ fmt.GoStringer
		var _ json.Marshaler
		var _ strings.Builder
		var _ = ioutil.Discard
		var _ http.RoundTripper
		var _ bytes.Buffer
		var _ url.Values
		var _ multipart.File
		var _ io.Reader
	`)

	for _, item := range file.AllMessages {
		m := NewMessage(generatedFile, item)
		m.P()
	}

	for _, item := range file.AllEnums {
		e := NewEnum(generatedFile, item)
		e.P()
	}

	for _, service := range file.AllService {
		s := NewService(generatedFile, service)
		s.P()
	}
}

// FileInfo file info
type FileInfo struct {
	*protogen.File
	f           *protogen.GeneratedFile
	AllMessages []*protogen.Message
	AllEnums    []*protogen.Enum
	AllService  []*protogen.Service
}

// NewFileInfo constructs then file info.
func NewFileInfo(file *protogen.File) *FileInfo {
	f := &FileInfo{File: file}
	// Collect all enums, messages, and extensions in "flattened ordering".
	var walkMessages func([]*protogen.Message, func(*protogen.Message))
	walkMessages = func(messages []*protogen.Message, f func(*protogen.Message)) {
		for _, m := range messages {
			f(m)
			walkMessages(m.Messages, f)
		}
	}

	f.AllEnums = append(f.AllEnums, f.Enums...)
	f.AllMessages = append(f.AllMessages, f.Messages...)
	f.AllService = append(f.AllService, f.Services...)

	walkMessages(f.Messages, func(m *protogen.Message) {
		f.AllEnums = append(f.AllEnums, m.Enums...)
		f.AllMessages = append(f.AllMessages, m.Messages...)
	})

	return f
}

// Service describes the service .
type Service struct {
	LeadingComments string
	ServiceName     string
	Methods         []*Method

	f *protogen.GeneratedFile
}

// NewService parses protogen.Service and returns it as Service.
func NewService(f *protogen.GeneratedFile, service *protogen.Service) *Service {
	s := &Service{
		ServiceName:     service.GoName,
		LeadingComments: string(appendDeprecationSuffix(service.Comments.Leading, service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated())),
		f:               f,
	}

	for _, m := range service.Methods {
		s.Methods = append(s.Methods, newMethod(f, m))
	}

	return s
}

// P print generated code into genereted file
func (s *Service) P() {
	s.f.P(s.genServiceInterface())
	s.f.P(s.genUnimplementedServer())
	s.f.P(s.genServiceDecorator())
	s.f.P(s.genServiceRegister())
	// s.f.P(s.genClientInterface())
	// s.f.P(s.genDefaultClient())
	s.f.P(s.genEndpointList())
}

var serviceInterfaceTempl = `
	// {{.ServiceName}}Server is the server API for {{.ServiceName}} service.
	type {{.ServiceName}}Server interface {
	{{- range .Methods}}
		{{.Comments -}}
		{{.Name}}(context.Context, *{{.Request.GoName}}) (*{{.Response.GoName}}, error)
	{{- end}}}
`

func (s *Service) genServiceInterface() string {
	return executeTemplate("serviceInterfaceTempl", serviceInterfaceTempl, s)
}

var allEndpointsTempl = `
	// All Endpoints 
	var (
		{{range .Methods -}}
			{{$methodName := .Name}}
			{{- range $index, $rule := .HTTPRules}}
			{{- if eq $index 0}}{{$methodName}}Endpoint{{else}}{{$methodName}}Endpoint_{{$index}}{{end}} = "{{$rule.Path}}"
			{{ end }}
		{{- end}}
	)
`

func (s *Service) genEndpointList() string {
	return executeTemplate("serviceEndpointList", allEndpointsTempl, s)
}

var unimplementedServerTempl = `
	{{$serviceName := .ServiceName}}
	// Unimplemented{{$serviceName}}Server can be embedded to have forward compatible implementations.
	type Unimplemented{{$serviceName}}Server struct {}

	{{- range .Methods}}
	func (s *Unimplemented{{$serviceName}}Server) {{.Name}}(context.Context, *{{.Request.GoName}}) (*{{.Response.GoName}}, error) {
		return nil, status.Errorf(codes.Unimplemented, "method {{.Name}} not implemented")
	}
	{{ end }}		
	`

func (s *Service) genUnimplementedServer() string {
	return executeTemplate("unimplementedServerTempl", unimplementedServerTempl, s)
}

var serviceDecoratorTempl = `
	{{$serviceName := .ServiceName}}
	// Default{{$serviceName}}Decorator the default decorator.
	type Default{{$serviceName}}Decorator struct{
		ss {{$serviceName}}Server
	}

	{{range .Methods}}
		{{$methodName := .Name}}
		{{$requestParamType := .Request.GoName}}
		{{$hasHeaderParam := .Request.HasHeaderParam}}
		{{$hasQueryParam := .Request.HasQueryParam}}
		{{$hasFile := .Request.HasFile}}
		{{range $index, $rule := .HTTPRules}}
			{{- if eq $index 0 -}}
			func (s *Default{{$serviceName}}Decorator) {{$methodName}}(ctx *gin.Context){
			{{- else -}}
			func (s *Default{{$serviceName}}Decorator) {{$methodName}}_{{$index}}(ctx *gin.Context){
			{{- end -}}
				var req {{$requestParamType}}
				
				{{- if eq $rule.Method "GET" "DELETE" }}
				{{- else}}
				shouldBindPayload := func(obj interface{}) error {
					switch ctx.ContentType() {
					case "":
						return ctx.ShouldBindJSON(obj)
					default:
						return ctx.ShouldBind(obj)
					}
				}
				{{- end }}

				bindingHandlers := []func(obj interface{}) error {
					{{- if eq $rule.Method "GET" "DELETE" }}
					{{- else }}
					shouldBindPayload,
					{{- end -}}
				
					{{- if $rule.HasPathParam }}
					ctx.ShouldBindUri,
					{{- end -}}
					{{ if $hasHeaderParam }}
					ctx.ShouldBindHeader,
					{{- end -}}
					{{ if $hasQueryParam }}
					ctx.ShouldBindQuery,
					{{- end }}
				}

				for _, doBinding := range bindingHandlers {
					if err := doBinding(&req); err != nil {
						runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error())) 
						return
					}
				}

				newCtx := runtime.NewContext(ctx)
				resp, err := s.ss.{{$methodName}}(newCtx, &req)
				if err != nil {
					runtime.HTTPError(ctx, err)
					return
				}
		
				runtime.ForwardResponseMessage(newCtx, resp)	
			}
		{{end}}
	{{end}}
`

func (s *Service) genServiceDecorator() string {
	return executeTemplate("serviceDecoratorTempl", serviceDecoratorTempl, s)
}

var registerTempl = `
	{{$serviceName := .ServiceName}}
	// Register{{$serviceName}}Server registers the http handlers for service {{$serviceName}} to "router".
	func Register{{$serviceName}}Server(router gin.IRouter, s {{$serviceName}}Server) {
		d := &Default{{$serviceName}}Decorator{ss: s}
		{{- range .Methods -}}
			{{ $methodName := .Name }}
			{{- range $index, $rule := .HTTPRules}}
				router.Handle("{{$rule.Method}}", "{{$rule.Path}}", {{if eq $index 0}}d.{{$methodName}}{{else}}d.{{$methodName}}_{{$index}}{{end}})
			{{- end}}
		{{- end}}
	}
`

func (s *Service) genServiceRegister() string {
	return executeTemplate("registerTempl", registerTempl, s)
}

var clientInterfaceTempl = `
	// {{.ServiceName}}Client is the client API for for {{.ServiceName}} service.
	type {{.ServiceName}}Client interface {
	{{- range .Methods}}
		{{.Comments -}}
		{{.Name}}(context.Context, *{{.Request.GoName}}) (*{{.Response.GoName}}, error)
	{{- end}}}
`

func (s *Service) genClientInterface() string {
	return executeTemplate("clientInterfaceTempl", clientInterfaceTempl, s)
}

var defaultClientTempl = `
	{{$serviceName := .ServiceName}}
	type default{{$serviceName}}Client struct {
		cc *http.Client
		host string
	}

	// New{{$serviceName}}Client creates a client API for {{$serviceName}} service.
	func New{{$serviceName}}Client(host string, cc *http.Client) {{$serviceName}}Client {
		return &default{{$serviceName}}Client{cc: cc, host: strings.TrimSuffix(host, "/")}
	}
	
	{{range .Methods}}
		{{$methodName := .Name}}
		{{$requestParamType := .Request.GoName}}
		{{$rule := index .HTTPRules 0}}
		{{$header := .Request.FieldWithTag "header"}}
		{{$body := .Request.FieldWithTag "json"}}
		{{$query := .Request.FieldWithTag "query"}}
		{{$hasFiles := .Request.HasFile}}
		{{$hasForm := .Request.HasFormParam}}
		{{$formField := .Request.FieldWithTag "form"}}
		{{$request := .Request}}
		func (c *default{{$serviceName}}Client) {{$methodName}} (ctx context.Context, req *{{.Request.GoName}}) (*{{.Response.GoName}}, error) {
			endpoint := fmt.Sprintf("%s%s", c.host, "{{$rule.Path}}")
			{{- range $key, $value := $rule.PathParams}}
				endpoint = strings.ReplaceAll(endpoint, ":{{$key}}", fmt.Sprint(req.{{$value}}))
			{{- end}}

			{{if eq $rule.Method "GET" "DELETE" }}
				hreq , err := http.NewRequest("{{$rule.Method}}", endpoint, nil)
				if err != nil {
					return nil, fmt.Errorf("failed to create request with error: %s", err)
				}
			{{- else}}
				{{ if or $hasFiles $hasForm }}
					body := &bytes.Buffer{}
					writer := multipart.NewWriter(body)
					{{if $hasForm }}
						{{range $formField }}
						{{$tag := .TagByName "form"}} writer.WriteField("{{$tag.Value}}", req.{{.GoName}})
						{{- end -}}
					{{end}}

					{{if $hasFiles}}
					for filedName, files := range req.MultipartFiles {
						for filename, reader := range files {
							fw, err := writer.CreateFormFile(filedName, filename)
							if err != nil {
								return nil, err
							}

							_, err = io.Copy(fw, reader)
							if err != nil {
								return nil, err
							}
						}
					}
					{{end}}
					if err := writer.Close(); err != nil {
						return nil, err
					}

					hreq, err := http.NewRequest("{{$rule.Method}}", endpoint, body)	
					if err != nil {
						return nil, fmt.Errorf("failed to create request with error: %s", err)
					}
				{{else}}
					data, err := json.Marshal(req)
					if err != nil {
						return nil, fmt.Errorf("failed to marshal request with error: %s", err)
					}
					
					hreq, err := http.NewRequest("{{$rule.Method}}", endpoint, bytes.NewBuffer(data))	
					if err != nil {
						return nil, fmt.Errorf("failed to create request with error: %s", err)
					}	
				{{end}}		

			{{- end}}

			{{if or $hasFiles $hasForm }}
			hreq.Header.Set("Content-Type", writer.FormDataContentType())
			{{else}}
			hreq.Header.Set("Content-Type", "application/json")
			{{end}}

			{{- range $header}}
				{{$tag := .TagByName "header"}} hreq.Header.Add("{{$tag.Value}}", req.{{.GoName}})
			{{- end}}

			{{if $query -}}
				var queries = url.Values{}
				{{- range $query }}
					{{$tag := .TagByName "query"}} queries.Add("{{$tag.Value}}", req.{{.GoName}})
				{{- end}}
				hreq.URL.RawQuery = queries.Encode()
			{{end}}

			res, err := c.cc.Do(hreq)
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()

			respBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}

			var resp {{.Response.GoName}}
			if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
				return nil, err
			}

			return &resp, nil
		}
	{{end}}
`

func (s *Service) genDefaultClient() string {
	return executeTemplate("default_client", defaultClientTempl, s)
}

// Method service's method
type Method struct {
	Name      string
	Request   *Message
	Response  *Message
	HTTPRules []*HTTPRule
	Comments  string
}

func newMethod(f *protogen.GeneratedFile, m *protogen.Method) *Method {
	// try to parse http rules
	rule, ok := proto.GetExtension(m.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
	if rule != nil && ok {
		request := NewMessage(f, m.Input)
		response := NewMessage(f, m.Output)

		method := &Method{
			Name:     m.GoName,
			Request:  request,  //f.QualifiedGoIdent(m.Input.GoIdent),
			Response: response, //f.QualifiedGoIdent(m.Output.GoIdent),
			Comments: m.Comments.Leading.String(),
		}

		method.HTTPRules = append(method.HTTPRules, extractHTTPRule(rule))
		for _, bind := range rule.AdditionalBindings {
			method.HTTPRules = append(method.HTTPRules, extractHTTPRule(bind))
		}

		for idx := range method.HTTPRules {
			for k := range method.HTTPRules[idx].PathParams {
				field, ok := request.QueryFieldByTagKeyAndValue("json", k)
				if ok {
					method.HTTPRules[idx].PathParams[k] = field.GoName
				}
			}
		}

		return method
	}

	panic(fmt.Sprintf("no http rules found for method: %s", m.GoName))
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
	PathParams   map[string]string
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

	var pathParams = make(map[string]string)
	paths := strings.Split(path, "/")
	for idx, item := range paths {
		if len(item) > 0 && (item[0] == '{' && item[len(item)-1] == '}') {
			holder := item[1 : len(item)-1]
			paths[idx] = ":" + holder
			// pathParams = append(pathParams, item[1:len(item)-1])
			pathParams[holder] = toCamel(holder)
		}
	}
	path = strings.Join(paths, "/")

	return &HTTPRule{Method: method, Path: path, HasPathParam: len(pathParams) > 0, PathParams: pathParams}
}

// Enum describes an enum.
type Enum struct {
	GoIdent         string // name of the generated Go type: the enum's name
	LeadingComments string
	Values          []EnumValue

	f *protogen.GeneratedFile
}

// EnumValue represents a enum value.
type EnumValue struct {
	LeadingComments string
	TrailingComment string
	ShortName       string // the short name of the declaration
	GoIdent         string // name of the generated Go type
	Type            string // Go type
	Number          int32  // the enum value as an integer.
}

// NewEnum construct an enum.
func NewEnum(f *protogen.GeneratedFile, enum *protogen.Enum) *Enum {
	leadingComments := appendDeprecationSuffix(enum.Comments.Leading,
		enum.Desc.Options().(*descriptorpb.EnumOptions).GetDeprecated())

	enumInfo := &Enum{
		GoIdent:         f.QualifiedGoIdent(enum.GoIdent),
		LeadingComments: leadingComments.String(),
		f:               f,
	}

	for _, value := range enum.Values {
		// f.Annotate(value.GoIdent.GoName, value.Location)
		leadingComments := appendDeprecationSuffix(value.Comments.Leading,
			value.Desc.Options().(*descriptorpb.EnumValueOptions).GetDeprecated())

		enumInfo.Values = append(enumInfo.Values, EnumValue{
			LeadingComments: leadingComments.String(),
			TrailingComment: trailingComment(value.Comments.Trailing).String(),
			GoIdent:         f.QualifiedGoIdent(value.GoIdent),
			ShortName:       string(value.Desc.Name()),
			Type:            f.QualifiedGoIdent(enum.GoIdent),
			Number:          int32(value.Desc.Number()),
		})
	}

	return enumInfo
}

// P generates and print enum definition code into generated file.
func (e Enum) P() {
	res := executeTemplate("enum", `{{.LeadingComments}} type {{.GoIdent}} int32

	const ({{range .Values}}
		{{.LeadingComments -}}
		{{.GoIdent}} {{.Type}} = {{.Number}} {{.TrailingComment}}
		{{- end}}
	)
	
	var (
		{{.GoIdent}}_name = map[int32]string{
			{{- range .Values}}
			{{.Number}}: "{{.ShortName}}",
			{{- end}}
		}

		{{.GoIdent}}_value = map[string]int32{
			{{- range .Values}}
			"{{- .ShortName}}": {{.Number}},
			{{- end}}
		}
	)

	func (x {{.GoIdent}}) String() string {
		return {{.GoIdent}}_name[int32(x)]
	}
	`, e)

	e.f.P(res)
}

// Message describes the message.
type Message struct {
	GoName          string
	LeadingComments string  // leading comments
	Fields          []Field // all fields
	f               *protogen.GeneratedFile
}

// NewMessage new message info
func NewMessage(f *protogen.GeneratedFile, m *protogen.Message) *Message {
	leadingComments := m.Comments.Leading.String()
	if leadingComments == "" {
		leadingComments = "\n"
	}

	msgInfo := &Message{
		LeadingComments: leadingComments,
		GoName:          f.QualifiedGoIdent(m.GoIdent),
		f:               f,
	}

	var multipartFilesFieldAdded bool
	for _, field := range m.Fields {
		fieldType, isPointer := fieldGoType(f, field)
		if isPointer {
			fieldType = "*" + fieldType
		}

		fieldInfo := Field{
			GoName: field.GoName,
			GoType: fieldType,
			Tags: []*Tag{
				{
					Key:   "json",
					Value: fmt.Sprintf(`%s,omitempty`, field.Desc.TextName()),
				},
			},
			LeadingComments: field.Comments.Leading.String(),
		}

		// var tags []string
		trailingComment := field.Comments.Trailing.String()

		if trailingComment != "" {
			// try to extract user customized tags from trailing comments
			trailingComment = strings.TrimPrefix(trailingComment, "//")
			rawTags := strings.Split(trailingComment, " ")
			for _, rt := range rawTags {
				res := strings.Split(strings.TrimSpace(rt), ":")
				if len(res) == 2 {
					fieldInfo.Tags = append(fieldInfo.Tags, &Tag{
						Key:   res[0],
						Value: strings.Trim(res[1], `"`),
					})

					if res[0] == "file" {
						switch fieldInfo.GoType {
						case "[]byte":
							fieldInfo.GoType = "*multipart.FileHeader"
						case "[][]byte":
							fieldInfo.GoType = "[]*multipart.FileHeader"
						default:
							panic(fmt.Errorf("field %s.%s should be with type `bytes` if it represents a file", msgInfo.GoName, fieldInfo.GoName))
						}

						if fieldInfo.LeadingComments != "" {
							fieldInfo.LeadingComments = fieldInfo.LeadingComments + "Note: Just for Server use only"
						} else {
							fieldInfo.LeadingComments = "// Note: Just for Server use only"
						}

						if !multipartFilesFieldAdded {
							msgInfo.Fields = append(msgInfo.Fields, Field{
								GoName: "MultipartFiles",
								GoType: "map[string]map[string]io.Reader",
								Tags: []*Tag{
									{
										Key:   "json",
										Value: "multipart_files",
									},
									{
										Key:   "file",
										Value: "multipart_files",
									},
								},
								LeadingComments: `// Nested Map for all multipart files. 
							// Keys of the outer map and inner map represent form-data keys and filename, respectively.
							// Note: Just for Client use only`,
							})
							multipartFilesFieldAdded = true
						}

					}
				}
			}
		}

		msgInfo.Fields = append(msgInfo.Fields, fieldInfo)
	}

	return msgInfo
}

// P generates definition code for message and print it into generated file.
func (mi Message) P() {

	fields := []string{}
	for _, field := range mi.Fields {
		fields = append(fields, field.String())
	}

	res := executeTemplate("message_type",
		`{{.LeadingComments -}}
		type {{.GoName}} struct {
			{{.Fields}}
		}`,
		map[string]interface{}{
			"MessageInfo":     mi,
			"GoName":          mi.GoName,
			"LeadingComments": mi.LeadingComments,
			"Fields":          strings.Join(fields, "\n"),
		})

	mi.f.P(res)
}

// HasHeaderParam returns true if there is at least one header param.
func (mi Message) HasHeaderParam() bool {
	return len(mi.FieldWithHeaderTag()) > 0
}

// HasQueryParam returns true if there is at least one query param.
func (mi Message) HasQueryParam() bool {
	return len(mi.FieldWithQueryTag()) > 0
}

// HasFormParam returns true if there is at least one form param.
func (mi Message) HasFormParam() bool {
	return len(mi.FieldWithFormTag()) > 0
}

// HasFile returns true if there is field represents a file.
func (mi Message) HasFile() bool {
	return len(mi.FieldWithFileTag()) > 0
}

// FieldWithTag returns all fileds that has specified tag.
func (mi Message) FieldWithTag(tagName string) []Field {
	var fields []Field
	for idx := range mi.Fields {
		if mi.Fields[idx].TagByName(tagName) != nil {
			fields = append(fields, mi.Fields[idx])
		}
	}

	return fields
}

// FieldWithHeaderTag returns all fields that have tag `header`.
func (mi Message) FieldWithHeaderTag() []Field {
	return mi.FieldWithTag("header")
}

// FieldWithQueryTag returns all fields that have tag `query`.
func (mi Message) FieldWithQueryTag() []Field {
	return mi.FieldWithTag("query")
}

// FieldWithJSONTag returns all fields that have tag `json`.
func (mi Message) FieldWithJSONTag() []Field {
	var fields []Field
	for idx := range mi.Fields {
		value := mi.Fields[idx].TagByName("json")
		if value != nil && value.Value != "-" {
			fields = append(fields, mi.Fields[idx])
		}
	}

	return fields
}

// FieldWithPathTag returns all fields that have tag `path`.
func (mi Message) FieldWithPathTag() []Field {
	return mi.FieldWithTag("path")
}

// FieldWithFormTag returns all fields that have tag `form`.
func (mi Message) FieldWithFormTag() []Field {
	return mi.FieldWithTag("form")
}

// FieldWithFileTag returns all fields that have tag `file`.
func (mi Message) FieldWithFileTag() []Field {
	return mi.FieldWithTag("file")
}

// QueryFieldByTagKeyAndValue returns field that has tag which key_name=`key` and value = `value`
func (mi Message) QueryFieldByTagKeyAndValue(key, value string) (Field, bool) {
	for idx := range mi.Fields {
		if mi.Fields[idx].TagByName(key) != nil && mi.Fields[idx].TagByValue(value) != nil {
			return mi.Fields[idx], true
		}
	}

	return Field{}, false
}

// Field describes field in structure.
type Field struct {
	GoName string // filed name

	GoType string // field type, e.g int32, string, bool
	// field tags

	LeadingComments string
	Tags            []*Tag
}

// TagByName finds tag by name.
func (fi Field) TagByName(tagName string) *Tag {
	for _, item := range fi.Tags {
		if item.Key == tagName {
			return &Tag{
				Key:   item.Key,
				Value: item.Value,
			}
		}
	}
	return nil
}

// TagByValue finds tag by value.
func (fi Field) TagByValue(tagValue string) *Tag {
	for _, item := range fi.Tags {
		if item.Value == tagValue {
			return &Tag{
				Key:   item.Key,
				Value: item.Value,
			}
		}
	}
	return nil
}

// TagStr constructs and returns the formated tag string.
func (fi Field) TagStr() string {
	if len(fi.Tags) == 0 {
		return ""
	}

	var tagStr []string
	for _, item := range fi.Tags {
		tagStr = append(tagStr, item.String())
	}

	return fmt.Sprintf("`%s`", strings.Join(tagStr, " "))
}

// String returns filed info as string.
//
// e.g: Name string `json:"name"`
func (fi Field) String() string {
	if fi.LeadingComments == "" {
		return fmt.Sprintf("%s %s %s", fi.GoName, fi.GoType, fi.TagStr())
	}

	if !strings.HasSuffix(fi.LeadingComments, "\n") {
		fi.LeadingComments = fi.LeadingComments + "\n"
	}
	return fmt.Sprintf("%s%s %s %s", fi.LeadingComments, fi.GoName, fi.GoType, fi.TagStr())
}

// Tag describes tag info for structure filed.
//
// e.g: json:"user_id"
type Tag struct {
	Key   string
	Value string
}

// String returns tag info as string.
func (ti Tag) String() string {
	key := ti.Key
	switch key {
	case "query", "file":
		key = "form" // gin binds query,file payload using form tag
	}
	return fmt.Sprintf(`%s:"%s"`, key, ti.Value)
}
