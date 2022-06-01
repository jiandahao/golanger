// Code generated. DO NOT EDIT.

package example

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	fmt "fmt"
	gin "github.com/gin-gonic/gin"
	runtime "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"
	status1 "google.golang.org/genproto/googleapis/rpc/status"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	ioutil "io/ioutil"
	multipart "mime/multipart"
	http "net/http"
	url "net/url"
	strings "strings"
)

var _ fmt.GoStringer
var _ json.Marshaler
var _ strings.Builder
var _ = ioutil.Discard
var _ http.RoundTripper
var _ bytes.Buffer
var _ url.Values
var _ multipart.File
var _ io.Reader

type ErrorResponse struct {
	CorrelationId string       `json:"correlationId,omitempty"`
	Error         *ErrorObject `json:"error,omitempty"`
}

type ErrorObject struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Intentionally complicated message type to cover many features of Protobuf.
type ABitOfEverything struct {
	SingleNested             *ABitOfEverything_Nested            `json:"single_nested,omitempty"`
	Uuid                     string                              `json:"uuid,omitempty"`
	Nested                   []*ABitOfEverything_Nested          `json:"nested,omitempty"`
	FloatValue               float32                             `json:"float_value,omitempty"`
	DoubleValue              float64                             `json:"double_value,omitempty"`
	Int64Value               int64                               `json:"int64_value,omitempty"`
	Uint64Value              uint64                              `json:"uint64_value,omitempty"`
	Int32Value               int32                               `json:"int32_value,omitempty"`
	Fixed64Value             uint64                              `json:"fixed64_value,omitempty"`
	Fixed32Value             uint32                              `json:"fixed32_value,omitempty"`
	BoolValue                bool                                `json:"bool_value,omitempty"`
	StringValue              string                              `json:"string_value,omitempty"`
	BytesValue               []byte                              `json:"bytes_value,omitempty"`
	Uint32Value              uint32                              `json:"uint32_value,omitempty"`
	EnumValue                NumericEnum                         `json:"enum_value,omitempty"`
	PathEnumValue            PathEnum                            `json:"path_enum_value,omitempty"`
	NestedPathEnumValue      MessagePathEnum_NestedPathEnum      `json:"nested_path_enum_value,omitempty"`
	Sfixed32Value            int32                               `json:"sfixed32_value,omitempty"`
	Sfixed64Value            int64                               `json:"sfixed64_value,omitempty"`
	Sint32Value              int32                               `json:"sint32_value,omitempty"`
	Sint64Value              int64                               `json:"sint64_value,omitempty"`
	RepeatedStringValue      []string                            `json:"repeated_string_value,omitempty"`
	OneofEmpty               *emptypb.Empty                      `json:"oneof_empty,omitempty"`
	OneofString              *string                             `json:"oneof_string,omitempty"`
	MapValue                 map[string]NumericEnum              `json:"map_value,omitempty"`
	MappedStringValue        map[string]string                   `json:"mapped_string_value,omitempty"`
	MappedNestedValue        map[string]*ABitOfEverything_Nested `json:"mapped_nested_value,omitempty"`
	NonConventionalNameValue string                              `json:"nonConventionalNameValue,omitempty"`
	TimestampValue           *timestamppb.Timestamp              `json:"timestamp_value,omitempty"`
	// repeated enum value. it is comma-separated in query

	RepeatedEnumValue []NumericEnum `json:"repeated_enum_value,omitempty"`
	// repeated numeric enum comment (This comment is overridden by the field annotation)

	RepeatedEnumAnnotation []NumericEnum `json:"repeated_enum_annotation,omitempty"`
	// numeric enum comment (This comment is overridden by the field annotation)

	EnumValueAnnotation NumericEnum `json:"enum_value_annotation,omitempty"`
	// repeated string comment (This comment is overridden by the field annotation)

	RepeatedStringAnnotation []string `json:"repeated_string_annotation,omitempty"`
	// repeated nested object comment (This comment is overridden by the field annotation)

	RepeatedNestedAnnotation []*ABitOfEverything_Nested `json:"repeated_nested_annotation,omitempty"`
	// nested object comments (This comment is overridden by the field annotation)

	NestedAnnotation  *ABitOfEverything_Nested `json:"nested_annotation,omitempty"`
	Int64OverrideType int64                    `json:"int64_override_type,omitempty"`
	// mark a field as required in Open API definition

	RequiredStringViaFieldBehaviorAnnotation string `json:"required_string_via_field_behavior_annotation,omitempty"`
	// mark a field as readonly in Open API definition

	OutputOnlyStringViaFieldBehaviorAnnotation string `json:"output_only_string_via_field_behavior_annotation,omitempty"`
}

// ABitOfEverythingRepeated is used to validate repeated path parameter functionality
type ABitOfEverythingRepeated struct {
	// repeated values. they are comma-separated in path

	PathRepeatedFloatValue    []float32     `json:"path_repeated_float_value,omitempty"`
	PathRepeatedDoubleValue   []float64     `json:"path_repeated_double_value,omitempty"`
	PathRepeatedInt64Value    []int64       `json:"path_repeated_int64_value,omitempty"`
	PathRepeatedUint64Value   []uint64      `json:"path_repeated_uint64_value,omitempty"`
	PathRepeatedInt32Value    []int32       `json:"path_repeated_int32_value,omitempty"`
	PathRepeatedFixed64Value  []uint64      `json:"path_repeated_fixed64_value,omitempty"`
	PathRepeatedFixed32Value  []uint32      `json:"path_repeated_fixed32_value,omitempty"`
	PathRepeatedBoolValue     []bool        `json:"path_repeated_bool_value,omitempty"`
	PathRepeatedStringValue   []string      `json:"path_repeated_string_value,omitempty"`
	PathRepeatedBytesValue    [][]byte      `json:"path_repeated_bytes_value,omitempty"`
	PathRepeatedUint32Value   []uint32      `json:"path_repeated_uint32_value,omitempty"`
	PathRepeatedEnumValue     []NumericEnum `json:"path_repeated_enum_value,omitempty"`
	PathRepeatedSfixed32Value []int32       `json:"path_repeated_sfixed32_value,omitempty"`
	PathRepeatedSfixed64Value []int64       `json:"path_repeated_sfixed64_value,omitempty"`
	PathRepeatedSint32Value   []int32       `json:"path_repeated_sint32_value,omitempty"`
	PathRepeatedSint64Value   []int64       `json:"path_repeated_sint64_value,omitempty"`
}

type CheckStatusResponse struct {
	Status *status1.Status `json:"status,omitempty"`
}

type Body struct {
	Name string `json:"name,omitempty"`
}

type MessageWithBody struct {
	Id   string `json:"id,omitempty"`
	Data *Body  `json:"data,omitempty"`
}

// UpdateV2Request request for update includes the message and the update mask
type UpdateV2Request struct {
	Abe *ABitOfEverything `json:"abe,omitempty"`
	// The paths to update.

	UpdateMask *fieldmaskpb.FieldMask `json:"update_mask,omitempty"`
}

// An example resource type from AIP-123 used to test the behavior described in
// the CreateBookRequest message.
//
// See: https://google.aip.dev/123
type Book struct {
	// The resource name of the book.
	//
	// Format: `publishers/{publisher}/books/{book}`
	//
	// Example: `publishers/1257894000000000000/books/my-book`

	Name string `json:"name,omitempty"`
	// Output only. The book's ID.

	Id string `json:"id,omitempty"`
	// Output only. Creation time of the book.

	CreateTime *timestamppb.Timestamp `json:"create_time,omitempty"`
}

// A standard Create message from AIP-133 with a user-specified ID.
// The user-specified ID (the `book_id` field in this example) must become a
// query parameter in the OpenAPI spec.
//
// See: https://google.aip.dev/133#user-specified-ids
type CreateBookRequest struct {
	// The publisher in which to create the book.
	//
	// Format: `publishers/{publisher}`
	//
	// Example: `publishers/1257894000000000000`

	Parent string `json:"parent,omitempty"`
	// The book to create.

	Book *Book `json:"book,omitempty"`
	// The ID to use for the book.
	//
	// This must start with an alphanumeric character.

	BookId string `json:"book_id,omitempty"`
}

// A standard Update message from AIP-134
//
// See: https://google.aip.dev/134#request-message
type UpdateBookRequest struct {
	// The book to update.
	//
	// The book's `name` field is used to identify the book to be updated.
	// Format: publishers/{publisher}/books/{book}

	Book *Book `json:"book,omitempty"`
	// The list of fields to be updated.

	UpdateMask *fieldmaskpb.FieldMask `json:"update_mask,omitempty"`
	// If set to true, and the book is not found, a new book will be created.
	// In this situation, `update_mask` is ignored.

	AllowMissing bool `json:"allow_missing,omitempty"`
}

type IdMessage struct {
	Uuid string `json:"uuid,omitempty"`
}

type StringMessage struct {
	Value string `json:"value,omitempty"`
}

type StringValue struct {
	Value string `json:"value,omitempty"`
}

type MessagePathEnum struct {
}

type MessageWithPathEnum struct {
	Value PathEnum `json:"value,omitempty"`
}

type MessageWithNestedPathEnum struct {
	Value MessagePathEnum_NestedPathEnum `json:"value,omitempty"`
}

// Nested is nested type.
type ABitOfEverything_Nested struct {
	// name is nested field.

	Name   string `json:"name,omitempty"`
	Amount uint32 `json:"amount,omitempty"`
	// DeepEnum comment.

	Ok ABitOfEverything_Nested_DeepEnum `json:"ok,omitempty"`
}

type ABitOfEverything_MapValueEntry struct {
	Key   string      `json:"key,omitempty"`
	Value NumericEnum `json:"value,omitempty"`
}

type ABitOfEverything_MappedStringValueEntry struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type ABitOfEverything_MappedNestedValueEntry struct {
	Key   string                   `json:"key,omitempty"`
	Value *ABitOfEverything_Nested `json:"value,omitempty"`
}

// NumericEnum is one or zero.
type NumericEnum int32

const (
	// ZERO means 0
	NumericEnum_ZERO NumericEnum = 0
	// ONE means 1
	NumericEnum_ONE NumericEnum = 1
)

type PathEnum int32

const (
	PathEnum_ABC PathEnum = 0
	PathEnum_DEF PathEnum = 1
)

// DeepEnum is one or zero.
type ABitOfEverything_Nested_DeepEnum int32

const (
	// FALSE is false.
	ABitOfEverything_Nested_FALSE ABitOfEverything_Nested_DeepEnum = 0
	// TRUE is true.
	ABitOfEverything_Nested_TRUE ABitOfEverything_Nested_DeepEnum = 1
)

type MessagePathEnum_NestedPathEnum int32

const (
	MessagePathEnum_GHI MessagePathEnum_NestedPathEnum = 0
	MessagePathEnum_JKL MessagePathEnum_NestedPathEnum = 1
)

// ABitOfEverythingServiceServer is the server API for ABitOfEverythingService service.
type ABitOfEverythingServiceServer interface {
	// Create a new ABitOfEverything
	//
	// This API creates a new ABitOfEverything
	Create(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CreateBody(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	// Create a book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	Lookup(context.Context, *IdMessage) (*ABitOfEverything, error)
	Update(context.Context, *ABitOfEverything) (*emptypb.Empty, error)
	UpdateV2(context.Context, *UpdateV2Request) (*emptypb.Empty, error)
	Delete(context.Context, *IdMessage) (*emptypb.Empty, error)
	GetQuery(context.Context, *ABitOfEverything) (*emptypb.Empty, error)
	GetRepeatedQuery(context.Context, *ABitOfEverythingRepeated) (*ABitOfEverythingRepeated, error)
	// Echo allows posting a StringMessage value.
	//
	// It also exposes multiple bindings.
	//
	// This makes it useful when validating that the OpenAPI v2 API
	// description exposes documentation correctly on all paths
	// defined as additional_bindings in the proto.
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	DeepPathEcho(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	// rpc NoBindings(google.protobuf.Duration) returns (google.protobuf.Empty) {}
	Timeout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ErrorWithDetails(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetMessageWithBody(context.Context, *MessageWithBody) (*emptypb.Empty, error)
	PostWithEmptyBody(context.Context, *Body) (*emptypb.Empty, error)
	CheckGetQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CheckNestedEnumGetQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CheckPostQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	OverwriteResponseContentType(context.Context, *emptypb.Empty) (*StringValue, error)
	CheckExternalPathEnum(context.Context, *MessageWithPathEnum) (*emptypb.Empty, error)
	CheckExternalNestedPathEnum(context.Context, *MessageWithNestedPathEnum) (*emptypb.Empty, error)
	CheckStatus(context.Context, *emptypb.Empty) (*CheckStatusResponse, error)
}

// UnimplementedABitOfEverythingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedABitOfEverythingServiceServer struct{}

func (s *UnimplementedABitOfEverythingServiceServer) Create(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CreateBody(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBody not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CreateBook(context.Context, *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Lookup(context.Context, *IdMessage) (*ABitOfEverything, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Update(context.Context, *ABitOfEverything) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) UpdateV2(context.Context, *UpdateV2Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateV2 not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Delete(context.Context, *IdMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) GetQuery(context.Context, *ABitOfEverything) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuery not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) GetRepeatedQuery(context.Context, *ABitOfEverythingRepeated) (*ABitOfEverythingRepeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepeatedQuery not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) DeepPathEcho(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeepPathEcho not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Timeout(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Timeout not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) ErrorWithDetails(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ErrorWithDetails not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) GetMessageWithBody(context.Context, *MessageWithBody) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageWithBody not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) PostWithEmptyBody(context.Context, *Body) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostWithEmptyBody not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckGetQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckGetQueryParams not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckNestedEnumGetQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNestedEnumGetQueryParams not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckPostQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPostQueryParams not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) OverwriteResponseContentType(context.Context, *emptypb.Empty) (*StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OverwriteResponseContentType not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckExternalPathEnum(context.Context, *MessageWithPathEnum) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExternalPathEnum not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckExternalNestedPathEnum(context.Context, *MessageWithNestedPathEnum) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExternalNestedPathEnum not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckStatus(context.Context, *emptypb.Empty) (*CheckStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}

type defaultABitOfEverythingServiceDecorator struct {
	ss ABitOfEverythingServiceServer
}

func (s defaultABitOfEverythingServiceDecorator) Create(ctx *gin.Context) {
	var req ABitOfEverything
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Create(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CreateBody(ctx *gin.Context) {
	var req ABitOfEverything
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.CreateBody(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CreateBook(ctx *gin.Context) {
	var req CreateBookRequest
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.CreateBook(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) UpdateBook(ctx *gin.Context) {
	var req UpdateBookRequest
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.UpdateBook(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Lookup(ctx *gin.Context) {
	var req IdMessage

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Lookup(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Update(ctx *gin.Context) {
	var req ABitOfEverything
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Update(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) UpdateV2(ctx *gin.Context) {
	var req UpdateV2Request
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.UpdateV2(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}
func (s defaultABitOfEverythingServiceDecorator) UpdateV2_1(ctx *gin.Context) {
	var req UpdateV2Request
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.UpdateV2(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}
func (s defaultABitOfEverythingServiceDecorator) UpdateV2_2(ctx *gin.Context) {
	var req UpdateV2Request
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.UpdateV2(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Delete(ctx *gin.Context) {
	var req IdMessage

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Delete(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) GetQuery(ctx *gin.Context) {
	var req ABitOfEverything

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.GetQuery(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) GetRepeatedQuery(ctx *gin.Context) {
	var req ABitOfEverythingRepeated

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.GetRepeatedQuery(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Echo(ctx *gin.Context) {
	var req StringMessage

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Echo(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}
func (s defaultABitOfEverythingServiceDecorator) Echo_1(ctx *gin.Context) {
	var req StringMessage
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Echo(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}
func (s defaultABitOfEverythingServiceDecorator) Echo_2(ctx *gin.Context) {
	var req StringMessage

	bindingHandlers := []func(obj interface{}) error{}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Echo(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) DeepPathEcho(ctx *gin.Context) {
	var req ABitOfEverything
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.DeepPathEcho(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Timeout(ctx *gin.Context) {
	var req emptypb.Empty

	bindingHandlers := []func(obj interface{}) error{}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Timeout(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) ErrorWithDetails(ctx *gin.Context) {
	var req emptypb.Empty

	bindingHandlers := []func(obj interface{}) error{}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.ErrorWithDetails(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) GetMessageWithBody(ctx *gin.Context) {
	var req MessageWithBody
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.GetMessageWithBody(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) PostWithEmptyBody(ctx *gin.Context) {
	var req Body
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.PostWithEmptyBody(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckGetQueryParams(ctx *gin.Context) {
	var req ABitOfEverything

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.CheckGetQueryParams(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckNestedEnumGetQueryParams(ctx *gin.Context) {
	var req ABitOfEverything

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.CheckNestedEnumGetQueryParams(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckPostQueryParams(ctx *gin.Context) {
	var req ABitOfEverything
	shouldBindPayload := func(obj interface{}) error {
		switch ctx.ContentType() {
		case "":
			return ctx.ShouldBindJSON(obj)
		default:
			return ctx.ShouldBind(obj)
		}
	}

	bindingHandlers := []func(obj interface{}) error{
		shouldBindPayload,
		ctx.ShouldBindUri,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.CheckPostQueryParams(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) OverwriteResponseContentType(ctx *gin.Context) {
	var req emptypb.Empty

	bindingHandlers := []func(obj interface{}) error{}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.OverwriteResponseContentType(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckExternalPathEnum(ctx *gin.Context) {
	var req MessageWithPathEnum

	bindingHandlers := []func(obj interface{}) error{}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.CheckExternalPathEnum(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckExternalNestedPathEnum(ctx *gin.Context) {
	var req MessageWithNestedPathEnum

	bindingHandlers := []func(obj interface{}) error{}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.CheckExternalNestedPathEnum(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckStatus(ctx *gin.Context) {
	var req emptypb.Empty

	bindingHandlers := []func(obj interface{}) error{}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.CheckStatus(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

// RegisterABitOfEverythingServiceServer registers the http handlers for service ABitOfEverythingService to "router".
func RegisterABitOfEverythingServiceServer(router gin.IRouter, s ABitOfEverythingServiceServer) {
	d := defaultABitOfEverythingServiceDecorator{ss: s}
	router.Handle("POST", "/v1/example/a_bit_of_everything/:float_value/:double_value/:int64_value/separator/:uint64_value/:int32_value/:fixed64_value/:fixed32_value/:bool_value/{string_value=strprefix/*}/:uint32_value/:sfixed32_value/:sfixed64_value/:sint32_value/:sint64_value/:nonConventionalNameValue/:enum_value/:path_enum_value/:nested_path_enum_value/:enum_value_annotation", d.Create)
	router.Handle("POST", "/v1/example/a_bit_of_everything", d.CreateBody)
	router.Handle("POST", "/v1/{parent=publishers/*}/books", d.CreateBook)
	router.Handle("PATCH", "/v1/{book.name=publishers/*/books/*}", d.UpdateBook)
	router.Handle("GET", "/v1/example/a_bit_of_everything/:uuid", d.Lookup)
	router.Handle("PUT", "/v1/example/a_bit_of_everything/:uuid", d.Update)
	router.Handle("PUT", "/v2/example/a_bit_of_everything/:abe.uuid", d.UpdateV2)
	router.Handle("PATCH", "/v2/example/a_bit_of_everything/:abe.uuid", d.UpdateV2_1)
	router.Handle("PATCH", "/v2a/example/a_bit_of_everything/:abe.uuid", d.UpdateV2_2)
	router.Handle("DELETE", "/v1/example/a_bit_of_everything/:uuid", d.Delete)
	router.Handle("GET", "/v1/example/a_bit_of_everything/query/:uuid", d.GetQuery)
	router.Handle("GET", "/v1/example/a_bit_of_everything_repeated/:path_repeated_float_value/:path_repeated_double_value/:path_repeated_int64_value/:path_repeated_uint64_value/:path_repeated_int32_value/:path_repeated_fixed64_value/:path_repeated_fixed32_value/:path_repeated_bool_value/:path_repeated_string_value/:path_repeated_bytes_value/:path_repeated_uint32_value/:path_repeated_enum_value/:path_repeated_sfixed32_value/:path_repeated_sfixed64_value/:path_repeated_sint32_value/:path_repeated_sint64_value", d.GetRepeatedQuery)
	router.Handle("GET", "/v1/example/a_bit_of_everything/echo/:value", d.Echo)
	router.Handle("POST", "/v2/example/echo", d.Echo_1)
	router.Handle("GET", "/v2/example/echo", d.Echo_2)
	router.Handle("POST", "/v1/example/deep_path/:single_nested.name", d.DeepPathEcho)
	router.Handle("GET", "/v2/example/timeout", d.Timeout)
	router.Handle("GET", "/v2/example/errorwithdetails", d.ErrorWithDetails)
	router.Handle("POST", "/v2/example/withbody/:id", d.GetMessageWithBody)
	router.Handle("POST", "/v2/example/postwithemptybody/:name", d.PostWithEmptyBody)
	router.Handle("GET", "/v1/example/a_bit_of_everything/params/get/:single_nested.name", d.CheckGetQueryParams)
	router.Handle("GET", "/v1/example/a_bit_of_everything/params/get/nested_enum/:single_nested.ok", d.CheckNestedEnumGetQueryParams)
	router.Handle("POST", "/v1/example/a_bit_of_everything/params/post/:string_value", d.CheckPostQueryParams)
	router.Handle("GET", "/v2/example/overwriteresponsecontenttype", d.OverwriteResponseContentType)
	router.Handle("GET", "/v2/{value}:check", d.CheckExternalPathEnum)
	router.Handle("GET", "/v3/{value}:check", d.CheckExternalNestedPathEnum)
	router.Handle("GET", "/v1/example/checkStatus", d.CheckStatus)
}

// ABitOfEverythingServiceClient is the client API for for ABitOfEverythingService service.
type ABitOfEverythingServiceClient interface {
	// Create a new ABitOfEverything
	//
	// This API creates a new ABitOfEverything
	Create(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CreateBody(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	// Create a book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	Lookup(context.Context, *IdMessage) (*ABitOfEverything, error)
	Update(context.Context, *ABitOfEverything) (*emptypb.Empty, error)
	UpdateV2(context.Context, *UpdateV2Request) (*emptypb.Empty, error)
	Delete(context.Context, *IdMessage) (*emptypb.Empty, error)
	GetQuery(context.Context, *ABitOfEverything) (*emptypb.Empty, error)
	GetRepeatedQuery(context.Context, *ABitOfEverythingRepeated) (*ABitOfEverythingRepeated, error)
	// Echo allows posting a StringMessage value.
	//
	// It also exposes multiple bindings.
	//
	// This makes it useful when validating that the OpenAPI v2 API
	// description exposes documentation correctly on all paths
	// defined as additional_bindings in the proto.
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	DeepPathEcho(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	// rpc NoBindings(google.protobuf.Duration) returns (google.protobuf.Empty) {}
	Timeout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ErrorWithDetails(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetMessageWithBody(context.Context, *MessageWithBody) (*emptypb.Empty, error)
	PostWithEmptyBody(context.Context, *Body) (*emptypb.Empty, error)
	CheckGetQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CheckNestedEnumGetQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CheckPostQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	OverwriteResponseContentType(context.Context, *emptypb.Empty) (*StringValue, error)
	CheckExternalPathEnum(context.Context, *MessageWithPathEnum) (*emptypb.Empty, error)
	CheckExternalNestedPathEnum(context.Context, *MessageWithNestedPathEnum) (*emptypb.Empty, error)
	CheckStatus(context.Context, *emptypb.Empty) (*CheckStatusResponse, error)
}

type defaultABitOfEverythingServiceClient struct {
	cc   *http.Client
	host string
}

// NewABitOfEverythingServiceClient creates a client API for ABitOfEverythingService service.
func NewABitOfEverythingServiceClient(host string, cc *http.Client) ABitOfEverythingServiceClient {
	return &defaultABitOfEverythingServiceClient{cc: cc, host: strings.TrimSuffix(host, "/")}
}

func (c *defaultABitOfEverythingServiceClient) Create(ctx context.Context, req *ABitOfEverything) (*ABitOfEverything, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/:float_value/:double_value/:int64_value/separator/:uint64_value/:int32_value/:fixed64_value/:fixed32_value/:bool_value/{string_value=strprefix/*}/:uint32_value/:sfixed32_value/:sfixed64_value/:sint32_value/:sint64_value/:nonConventionalNameValue/:enum_value/:path_enum_value/:nested_path_enum_value/:enum_value_annotation")
	endpoint = strings.ReplaceAll(endpoint, ":bool_value", fmt.Sprint(req.BoolValue))
	endpoint = strings.ReplaceAll(endpoint, ":double_value", fmt.Sprint(req.DoubleValue))
	endpoint = strings.ReplaceAll(endpoint, ":enum_value", fmt.Sprint(req.EnumValue))
	endpoint = strings.ReplaceAll(endpoint, ":enum_value_annotation", fmt.Sprint(req.EnumValueAnnotation))
	endpoint = strings.ReplaceAll(endpoint, ":fixed32_value", fmt.Sprint(req.Fixed32Value))
	endpoint = strings.ReplaceAll(endpoint, ":fixed64_value", fmt.Sprint(req.Fixed64Value))
	endpoint = strings.ReplaceAll(endpoint, ":float_value", fmt.Sprint(req.FloatValue))
	endpoint = strings.ReplaceAll(endpoint, ":int32_value", fmt.Sprint(req.Int32Value))
	endpoint = strings.ReplaceAll(endpoint, ":int64_value", fmt.Sprint(req.Int64Value))
	endpoint = strings.ReplaceAll(endpoint, ":nested_path_enum_value", fmt.Sprint(req.NestedPathEnumValue))
	endpoint = strings.ReplaceAll(endpoint, ":nonConventionalNameValue", fmt.Sprint(req.NonConventionalNameValue))
	endpoint = strings.ReplaceAll(endpoint, ":path_enum_value", fmt.Sprint(req.PathEnumValue))
	endpoint = strings.ReplaceAll(endpoint, ":sfixed32_value", fmt.Sprint(req.Sfixed32Value))
	endpoint = strings.ReplaceAll(endpoint, ":sfixed64_value", fmt.Sprint(req.Sfixed64Value))
	endpoint = strings.ReplaceAll(endpoint, ":sint32_value", fmt.Sprint(req.Sint32Value))
	endpoint = strings.ReplaceAll(endpoint, ":sint64_value", fmt.Sprint(req.Sint64Value))
	endpoint = strings.ReplaceAll(endpoint, ":uint32_value", fmt.Sprint(req.Uint32Value))
	endpoint = strings.ReplaceAll(endpoint, ":uint64_value", fmt.Sprint(req.Uint64Value))

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp ABitOfEverything
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) CreateBody(ctx context.Context, req *ABitOfEverything) (*ABitOfEverything, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything")

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp ABitOfEverything
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) CreateBook(ctx context.Context, req *CreateBookRequest) (*Book, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/{parent=publishers/*}/books")

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp Book
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) UpdateBook(ctx context.Context, req *UpdateBookRequest) (*Book, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/{book.name=publishers/*/books/*}")

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp Book
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) Lookup(ctx context.Context, req *IdMessage) (*ABitOfEverything, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/:uuid")
	endpoint = strings.ReplaceAll(endpoint, ":uuid", fmt.Sprint(req.Uuid))

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp ABitOfEverything
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) Update(ctx context.Context, req *ABitOfEverything) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/:uuid")
	endpoint = strings.ReplaceAll(endpoint, ":uuid", fmt.Sprint(req.Uuid))

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) UpdateV2(ctx context.Context, req *UpdateV2Request) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v2/example/a_bit_of_everything/:abe.uuid")
	endpoint = strings.ReplaceAll(endpoint, ":abe.uuid", fmt.Sprint(req.Abe.Uuid))

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) Delete(ctx context.Context, req *IdMessage) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/:uuid")
	endpoint = strings.ReplaceAll(endpoint, ":uuid", fmt.Sprint(req.Uuid))

	hreq, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) GetQuery(ctx context.Context, req *ABitOfEverything) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/query/:uuid")
	endpoint = strings.ReplaceAll(endpoint, ":uuid", fmt.Sprint(req.Uuid))

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) GetRepeatedQuery(ctx context.Context, req *ABitOfEverythingRepeated) (*ABitOfEverythingRepeated, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything_repeated/:path_repeated_float_value/:path_repeated_double_value/:path_repeated_int64_value/:path_repeated_uint64_value/:path_repeated_int32_value/:path_repeated_fixed64_value/:path_repeated_fixed32_value/:path_repeated_bool_value/:path_repeated_string_value/:path_repeated_bytes_value/:path_repeated_uint32_value/:path_repeated_enum_value/:path_repeated_sfixed32_value/:path_repeated_sfixed64_value/:path_repeated_sint32_value/:path_repeated_sint64_value")
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_bool_value", fmt.Sprint(req.PathRepeatedBoolValue))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_bytes_value", fmt.Sprint(req.PathRepeatedBytesValue))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_double_value", fmt.Sprint(req.PathRepeatedDoubleValue))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_enum_value", fmt.Sprint(req.PathRepeatedEnumValue))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_fixed32_value", fmt.Sprint(req.PathRepeatedFixed32Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_fixed64_value", fmt.Sprint(req.PathRepeatedFixed64Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_float_value", fmt.Sprint(req.PathRepeatedFloatValue))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_int32_value", fmt.Sprint(req.PathRepeatedInt32Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_int64_value", fmt.Sprint(req.PathRepeatedInt64Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_sfixed32_value", fmt.Sprint(req.PathRepeatedSfixed32Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_sfixed64_value", fmt.Sprint(req.PathRepeatedSfixed64Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_sint32_value", fmt.Sprint(req.PathRepeatedSint32Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_sint64_value", fmt.Sprint(req.PathRepeatedSint64Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_string_value", fmt.Sprint(req.PathRepeatedStringValue))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_uint32_value", fmt.Sprint(req.PathRepeatedUint32Value))
	endpoint = strings.ReplaceAll(endpoint, ":path_repeated_uint64_value", fmt.Sprint(req.PathRepeatedUint64Value))

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp ABitOfEverythingRepeated
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) Echo(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/echo/:value")
	endpoint = strings.ReplaceAll(endpoint, ":value", fmt.Sprint(req.Value))

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp StringMessage
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) DeepPathEcho(ctx context.Context, req *ABitOfEverything) (*ABitOfEverything, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/deep_path/:single_nested.name")
	endpoint = strings.ReplaceAll(endpoint, ":single_nested.name", fmt.Sprint(req.SingleNested.Name))

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp ABitOfEverything
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) Timeout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v2/example/timeout")

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) ErrorWithDetails(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v2/example/errorwithdetails")

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) GetMessageWithBody(ctx context.Context, req *MessageWithBody) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v2/example/withbody/:id")
	endpoint = strings.ReplaceAll(endpoint, ":id", fmt.Sprint(req.Id))

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) PostWithEmptyBody(ctx context.Context, req *Body) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v2/example/postwithemptybody/:name")
	endpoint = strings.ReplaceAll(endpoint, ":name", fmt.Sprint(req.Name))

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) CheckGetQueryParams(ctx context.Context, req *ABitOfEverything) (*ABitOfEverything, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/params/get/:single_nested.name")
	endpoint = strings.ReplaceAll(endpoint, ":single_nested.name", fmt.Sprint(req.SingleNested.Name))

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp ABitOfEverything
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) CheckNestedEnumGetQueryParams(ctx context.Context, req *ABitOfEverything) (*ABitOfEverything, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/params/get/nested_enum/:single_nested.ok")
	endpoint = strings.ReplaceAll(endpoint, ":single_nested.ok", fmt.Sprint(req.SingleNested.Ok))

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp ABitOfEverything
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) CheckPostQueryParams(ctx context.Context, req *ABitOfEverything) (*ABitOfEverything, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/a_bit_of_everything/params/post/:string_value")
	endpoint = strings.ReplaceAll(endpoint, ":string_value", fmt.Sprint(req.StringValue))

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp ABitOfEverything
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) OverwriteResponseContentType(ctx context.Context, req *emptypb.Empty) (*StringValue, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v2/example/overwriteresponsecontenttype")

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp StringValue
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) CheckExternalPathEnum(ctx context.Context, req *MessageWithPathEnum) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v2/{value}:check")

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) CheckExternalNestedPathEnum(ctx context.Context, req *MessageWithNestedPathEnum) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v3/{value}:check")

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultABitOfEverythingServiceClient) CheckStatus(ctx context.Context, req *emptypb.Empty) (*CheckStatusResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/example/checkStatus")

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp CheckStatusResponse
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CamelCaseServiceNameServer is the server API for CamelCaseServiceName service.
type CamelCaseServiceNameServer interface {
	Empty(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

// UnimplementedCamelCaseServiceNameServer can be embedded to have forward compatible implementations.
type UnimplementedCamelCaseServiceNameServer struct{}

func (s *UnimplementedCamelCaseServiceNameServer) Empty(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Empty not implemented")
}

type defaultCamelCaseServiceNameDecorator struct {
	ss CamelCaseServiceNameServer
}

func (s defaultCamelCaseServiceNameDecorator) Empty(ctx *gin.Context) {
	var req emptypb.Empty

	bindingHandlers := []func(obj interface{}) error{}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.Empty(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

// RegisterCamelCaseServiceNameServer registers the http handlers for service CamelCaseServiceName to "router".
func RegisterCamelCaseServiceNameServer(router gin.IRouter, s CamelCaseServiceNameServer) {
	d := defaultCamelCaseServiceNameDecorator{ss: s}
	router.Handle("GET", "/v2/example/empty", d.Empty)
}

// CamelCaseServiceNameClient is the client API for for CamelCaseServiceName service.
type CamelCaseServiceNameClient interface {
	Empty(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

type defaultCamelCaseServiceNameClient struct {
	cc   *http.Client
	host string
}

// NewCamelCaseServiceNameClient creates a client API for CamelCaseServiceName service.
func NewCamelCaseServiceNameClient(host string, cc *http.Client) CamelCaseServiceNameClient {
	return &defaultCamelCaseServiceNameClient{cc: cc, host: strings.TrimSuffix(host, "/")}
}

func (c *defaultCamelCaseServiceNameClient) Empty(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v2/example/empty")

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp emptypb.Empty
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
