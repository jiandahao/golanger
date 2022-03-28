// Code generated. DO NOT EDIT.

package descriptorpb

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	runtime "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// The protocol compiler can output a FileDescriptorSet containing the .proto
// files it parses.
type FileDescriptorSet struct {
	File []*FileDescriptorProto `json:"file,omitempty"`
}

// Describes a complete .proto file.
type FileDescriptorProto struct {
	Name             *string                   `json:"name,omitempty"`
	Package          *string                   `json:"package,omitempty"`
	Dependency       []string                  `json:"dependency,omitempty"`
	PublicDependency []int32                   `json:"public_dependency,omitempty"`
	WeakDependency   []int32                   `json:"weak_dependency,omitempty"`
	MessageType      []*DescriptorProto        `json:"message_type,omitempty"`
	EnumType         []*EnumDescriptorProto    `json:"enum_type,omitempty"`
	Service          []*ServiceDescriptorProto `json:"service,omitempty"`
	Extension        []*FieldDescriptorProto   `json:"extension,omitempty"`
	Options          *FileOptions              `json:"options,omitempty"`
	SourceCodeInfo   *SourceCodeInfo           `json:"source_code_info,omitempty"`
	Syntax           *string                   `json:"syntax,omitempty"`
}

// Describes a message type.
type DescriptorProto struct {
	Name           *string                           `json:"name,omitempty"`
	Field          []*FieldDescriptorProto           `json:"field,omitempty"`
	Extension      []*FieldDescriptorProto           `json:"extension,omitempty"`
	NestedType     []*DescriptorProto                `json:"nested_type,omitempty"`
	EnumType       []*EnumDescriptorProto            `json:"enum_type,omitempty"`
	ExtensionRange []*DescriptorProto_ExtensionRange `json:"extension_range,omitempty"`
	OneofDecl      []*OneofDescriptorProto           `json:"oneof_decl,omitempty"`
	Options        *MessageOptions                   `json:"options,omitempty"`
	ReservedRange  []*DescriptorProto_ReservedRange  `json:"reserved_range,omitempty"`
	ReservedName   []string                          `json:"reserved_name,omitempty"`
}

type ExtensionRangeOptions struct {
	UninterpretedOption []*UninterpretedOption `json:"uninterpreted_option,omitempty"`
}

// Describes a field within a message.
type FieldDescriptorProto struct {
	Name           *string                     `json:"name,omitempty"`
	Number         *int32                      `json:"number,omitempty"`
	Label          *FieldDescriptorProto_Label `json:"label,omitempty"`
	Type           *FieldDescriptorProto_Type  `json:"type,omitempty"`
	TypeName       *string                     `json:"type_name,omitempty"`
	Extendee       *string                     `json:"extendee,omitempty"`
	DefaultValue   *string                     `json:"default_value,omitempty"`
	OneofIndex     *int32                      `json:"oneof_index,omitempty"`
	JsonName       *string                     `json:"json_name,omitempty"`
	Options        *FieldOptions               `json:"options,omitempty"`
	Proto3Optional *bool                       `json:"proto3_optional,omitempty"`
}

// Describes a oneof.
type OneofDescriptorProto struct {
	Name    *string       `json:"name,omitempty"`
	Options *OneofOptions `json:"options,omitempty"`
}

// Describes an enum type.
type EnumDescriptorProto struct {
	Name          *string                                  `json:"name,omitempty"`
	Value         []*EnumValueDescriptorProto              `json:"value,omitempty"`
	Options       *EnumOptions                             `json:"options,omitempty"`
	ReservedRange []*EnumDescriptorProto_EnumReservedRange `json:"reserved_range,omitempty"`
	ReservedName  []string                                 `json:"reserved_name,omitempty"`
}

// Describes a value within an enum.
type EnumValueDescriptorProto struct {
	Name    *string           `json:"name,omitempty"`
	Number  *int32            `json:"number,omitempty"`
	Options *EnumValueOptions `json:"options,omitempty"`
}

// Describes a service.
type ServiceDescriptorProto struct {
	Name    *string                  `json:"name,omitempty"`
	Method  []*MethodDescriptorProto `json:"method,omitempty"`
	Options *ServiceOptions          `json:"options,omitempty"`
}

// Describes a method of a service.
type MethodDescriptorProto struct {
	Name            *string        `json:"name,omitempty"`
	InputType       *string        `json:"input_type,omitempty"`
	OutputType      *string        `json:"output_type,omitempty"`
	Options         *MethodOptions `json:"options,omitempty"`
	ClientStreaming *bool          `json:"client_streaming,omitempty"`
	ServerStreaming *bool          `json:"server_streaming,omitempty"`
}

type FileOptions struct {
	JavaPackage               *string                   `json:"java_package,omitempty"`
	JavaOuterClassname        *string                   `json:"java_outer_classname,omitempty"`
	JavaMultipleFiles         *bool                     `json:"java_multiple_files,omitempty"`
	JavaGenerateEqualsAndHash *bool                     `json:"java_generate_equals_and_hash,omitempty"`
	JavaStringCheckUtf8       *bool                     `json:"java_string_check_utf8,omitempty"`
	OptimizeFor               *FileOptions_OptimizeMode `json:"optimize_for,omitempty"`
	GoPackage                 *string                   `json:"go_package,omitempty"`
	CcGenericServices         *bool                     `json:"cc_generic_services,omitempty"`
	JavaGenericServices       *bool                     `json:"java_generic_services,omitempty"`
	PyGenericServices         *bool                     `json:"py_generic_services,omitempty"`
	PhpGenericServices        *bool                     `json:"php_generic_services,omitempty"`
	Deprecated                *bool                     `json:"deprecated,omitempty"`
	CcEnableArenas            *bool                     `json:"cc_enable_arenas,omitempty"`
	ObjcClassPrefix           *string                   `json:"objc_class_prefix,omitempty"`
	CsharpNamespace           *string                   `json:"csharp_namespace,omitempty"`
	SwiftPrefix               *string                   `json:"swift_prefix,omitempty"`
	PhpClassPrefix            *string                   `json:"php_class_prefix,omitempty"`
	PhpNamespace              *string                   `json:"php_namespace,omitempty"`
	PhpMetadataNamespace      *string                   `json:"php_metadata_namespace,omitempty"`
	RubyPackage               *string                   `json:"ruby_package,omitempty"`
	UninterpretedOption       []*UninterpretedOption    `json:"uninterpreted_option,omitempty"`
}

type MessageOptions struct {
	MessageSetWireFormat         *bool                  `json:"message_set_wire_format,omitempty"`
	NoStandardDescriptorAccessor *bool                  `json:"no_standard_descriptor_accessor,omitempty"`
	Deprecated                   *bool                  `json:"deprecated,omitempty"`
	MapEntry                     *bool                  `json:"map_entry,omitempty"`
	UninterpretedOption          []*UninterpretedOption `json:"uninterpreted_option,omitempty"`
}

type FieldOptions struct {
	Ctype               *FieldOptions_CType    `json:"ctype,omitempty"`
	Packed              *bool                  `json:"packed,omitempty"`
	Jstype              *FieldOptions_JSType   `json:"jstype,omitempty"`
	Lazy                *bool                  `json:"lazy,omitempty"`
	Deprecated          *bool                  `json:"deprecated,omitempty"`
	Weak                *bool                  `json:"weak,omitempty"`
	UninterpretedOption []*UninterpretedOption `json:"uninterpreted_option,omitempty"`
}

type OneofOptions struct {
	UninterpretedOption []*UninterpretedOption `json:"uninterpreted_option,omitempty"`
}

type EnumOptions struct {
	AllowAlias          *bool                  `json:"allow_alias,omitempty"`
	Deprecated          *bool                  `json:"deprecated,omitempty"`
	UninterpretedOption []*UninterpretedOption `json:"uninterpreted_option,omitempty"`
}

type EnumValueOptions struct {
	Deprecated          *bool                  `json:"deprecated,omitempty"`
	UninterpretedOption []*UninterpretedOption `json:"uninterpreted_option,omitempty"`
}

type ServiceOptions struct {
	Deprecated          *bool                  `json:"deprecated,omitempty"`
	UninterpretedOption []*UninterpretedOption `json:"uninterpreted_option,omitempty"`
}

type MethodOptions struct {
	Deprecated          *bool                           `json:"deprecated,omitempty"`
	IdempotencyLevel    *MethodOptions_IdempotencyLevel `json:"idempotency_level,omitempty"`
	UninterpretedOption []*UninterpretedOption          `json:"uninterpreted_option,omitempty"`
}

// A message representing a option the parser does not recognize. This only
// appears in options protos created by the compiler::Parser class.
// DescriptorPool resolves these when building Descriptor objects. Therefore,
// options protos in descriptor objects (e.g. returned by Descriptor::options(),
// or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
// in them.
type UninterpretedOption struct {
	Name             []*UninterpretedOption_NamePart `json:"name,omitempty"`
	IdentifierValue  *string                         `json:"identifier_value,omitempty"`
	PositiveIntValue *uint64                         `json:"positive_int_value,omitempty"`
	NegativeIntValue *int64                          `json:"negative_int_value,omitempty"`
	DoubleValue      *float64                        `json:"double_value,omitempty"`
	StringValue      []byte                          `json:"string_value,omitempty"`
	AggregateValue   *string                         `json:"aggregate_value,omitempty"`
}

// Encapsulates information about the original source file from which a
// FileDescriptorProto was generated.
type SourceCodeInfo struct {
	Location []*SourceCodeInfo_Location `json:"location,omitempty"`
}

// Describes the relationship between generated code and its original source
// file. A GeneratedCodeInfo message is associated with only one generated
// source file, but may contain references to different source .proto files.
type GeneratedCodeInfo struct {
	Annotation []*GeneratedCodeInfo_Annotation `json:"annotation,omitempty"`
}

type DescriptorProto_ExtensionRange struct {
	Start   *int32                 `json:"start,omitempty"`
	End     *int32                 `json:"end,omitempty"`
	Options *ExtensionRangeOptions `json:"options,omitempty"`
}

// Range of reserved tag numbers. Reserved tag numbers may not be used by
// fields or extension ranges in the same message. Reserved ranges may
// not overlap.
type DescriptorProto_ReservedRange struct {
	Start *int32 `json:"start,omitempty"`
	End   *int32 `json:"end,omitempty"`
}

// Range of reserved numeric values. Reserved values may not be used by
// entries in the same enum. Reserved ranges may not overlap.
//
// Note that this is distinct from DescriptorProto.ReservedRange in that it
// is inclusive such that it can appropriately represent the entire int32
// domain.
type EnumDescriptorProto_EnumReservedRange struct {
	Start *int32 `json:"start,omitempty"`
	End   *int32 `json:"end,omitempty"`
}

// The name of the uninterpreted option.  Each string represents a segment in
// a dot-separated name.  is_extension is true iff a segment represents an
// extension (denoted with parentheses in options specs in .proto files).
// E.g.,{ ["foo", false], ["bar.baz", true], ["qux", false] } represents
// "foo.(bar.baz).qux".
type UninterpretedOption_NamePart struct {
	NamePart    *string `json:"name_part,omitempty"`
	IsExtension *bool   `json:"is_extension,omitempty"`
}

type SourceCodeInfo_Location struct {
	Path                    []int32  `json:"path,omitempty"`
	Span                    []int32  `json:"span,omitempty"`
	LeadingComments         *string  `json:"leading_comments,omitempty"`
	TrailingComments        *string  `json:"trailing_comments,omitempty"`
	LeadingDetachedComments []string `json:"leading_detached_comments,omitempty"`
}

type GeneratedCodeInfo_Annotation struct {
	Path       []int32 `json:"path,omitempty"`
	SourceFile *string `json:"source_file,omitempty"`
	Begin      *int32  `json:"begin,omitempty"`
	End        *int32  `json:"end,omitempty"`
}
type FieldDescriptorProto_Type int32

const (
	// 0 is reserved for errors.
	// Order is weird for historical reasons.
	FieldDescriptorProto_TYPE_DOUBLE FieldDescriptorProto_Type = 1
	FieldDescriptorProto_TYPE_FLOAT  FieldDescriptorProto_Type = 2
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
	// negative values are likely.
	FieldDescriptorProto_TYPE_INT64  FieldDescriptorProto_Type = 3
	FieldDescriptorProto_TYPE_UINT64 FieldDescriptorProto_Type = 4
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
	// negative values are likely.
	FieldDescriptorProto_TYPE_INT32   FieldDescriptorProto_Type = 5
	FieldDescriptorProto_TYPE_FIXED64 FieldDescriptorProto_Type = 6
	FieldDescriptorProto_TYPE_FIXED32 FieldDescriptorProto_Type = 7
	FieldDescriptorProto_TYPE_BOOL    FieldDescriptorProto_Type = 8
	FieldDescriptorProto_TYPE_STRING  FieldDescriptorProto_Type = 9
	// Tag-delimited aggregate.
	// Group type is deprecated and not supported in proto3. However, Proto3
	// implementations should still be able to parse the group wire format and
	// treat group fields as unknown fields.
	FieldDescriptorProto_TYPE_GROUP   FieldDescriptorProto_Type = 10
	FieldDescriptorProto_TYPE_MESSAGE FieldDescriptorProto_Type = 11 // Length-delimited aggregate.
	// New in version 2.
	FieldDescriptorProto_TYPE_BYTES    FieldDescriptorProto_Type = 12
	FieldDescriptorProto_TYPE_UINT32   FieldDescriptorProto_Type = 13
	FieldDescriptorProto_TYPE_ENUM     FieldDescriptorProto_Type = 14
	FieldDescriptorProto_TYPE_SFIXED32 FieldDescriptorProto_Type = 15
	FieldDescriptorProto_TYPE_SFIXED64 FieldDescriptorProto_Type = 16
	FieldDescriptorProto_TYPE_SINT32   FieldDescriptorProto_Type = 17 // Uses ZigZag encoding.
	FieldDescriptorProto_TYPE_SINT64   FieldDescriptorProto_Type = 18 // Uses ZigZag encoding.
)

type FieldDescriptorProto_Label int32

const (
	// 0 is reserved for errors
	FieldDescriptorProto_LABEL_OPTIONAL FieldDescriptorProto_Label = 1
	FieldDescriptorProto_LABEL_REQUIRED FieldDescriptorProto_Label = 2
	FieldDescriptorProto_LABEL_REPEATED FieldDescriptorProto_Label = 3
)

// Generated classes can be optimized for speed or code size.
type FileOptions_OptimizeMode int32

const (
	FileOptions_SPEED FileOptions_OptimizeMode = 1 // Generate complete code for parsing, serialization,
	// etc.
	FileOptions_CODE_SIZE    FileOptions_OptimizeMode = 2 // Use ReflectionOps to implement these methods.
	FileOptions_LITE_RUNTIME FileOptions_OptimizeMode = 3 // Generate code using MessageLite and the lite runtime.
)

type FieldOptions_CType int32

const (
	// Default mode.
	FieldOptions_STRING       FieldOptions_CType = 0
	FieldOptions_CORD         FieldOptions_CType = 1
	FieldOptions_STRING_PIECE FieldOptions_CType = 2
)

type FieldOptions_JSType int32

const (
	// Use the default type.
	FieldOptions_JS_NORMAL FieldOptions_JSType = 0
	// Use JavaScript strings.
	FieldOptions_JS_STRING FieldOptions_JSType = 1
	// Use JavaScript numbers.
	FieldOptions_JS_NUMBER FieldOptions_JSType = 2
)

// Is this method side-effect-free (or safe in HTTP parlance), or idempotent,
// or neither? HTTP based RPC implementation may choose GET verb for safe
// methods, and PUT verb for idempotent methods instead of the default POST.
type MethodOptions_IdempotencyLevel int32

const (
	MethodOptions_IDEMPOTENCY_UNKNOWN MethodOptions_IdempotencyLevel = 0
	MethodOptions_NO_SIDE_EFFECTS     MethodOptions_IdempotencyLevel = 1 // implies idempotent
	MethodOptions_IDEMPOTENT          MethodOptions_IdempotencyLevel = 2 // idempotent, but may have side effects
)
