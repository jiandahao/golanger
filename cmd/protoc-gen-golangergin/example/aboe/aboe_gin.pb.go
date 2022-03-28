// Code generated. DO NOT EDIT.

package example

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	runtime "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"
	status "github.com/jiandahao/golanger/pkg/generator/gingen/status"
	status1 "google.golang.org/genproto/googleapis/rpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

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
	SingleNested                               *ABitOfEverything_Nested            `json:"single_nested,omitempty"`
	Uuid                                       string                              `json:"uuid,omitempty"`
	Nested                                     []*ABitOfEverything_Nested          `json:"nested,omitempty"`
	FloatValue                                 float32                             `json:"float_value,omitempty"`
	DoubleValue                                float64                             `json:"double_value,omitempty"`
	Int64Value                                 int64                               `json:"int64_value,omitempty"`
	Uint64Value                                uint64                              `json:"uint64_value,omitempty"`
	Int32Value                                 int32                               `json:"int32_value,omitempty"`
	Fixed64Value                               uint64                              `json:"fixed64_value,omitempty"`
	Fixed32Value                               uint32                              `json:"fixed32_value,omitempty"`
	BoolValue                                  bool                                `json:"bool_value,omitempty"`
	StringValue                                string                              `json:"string_value,omitempty"`
	BytesValue                                 []byte                              `json:"bytes_value,omitempty"`
	Uint32Value                                uint32                              `json:"uint32_value,omitempty"`
	EnumValue                                  NumericEnum                         `json:"enum_value,omitempty"`
	PathEnumValue                              PathEnum                            `json:"path_enum_value,omitempty"`
	NestedPathEnumValue                        MessagePathEnum_NestedPathEnum      `json:"nested_path_enum_value,omitempty"`
	Sfixed32Value                              int32                               `json:"sfixed32_value,omitempty"`
	Sfixed64Value                              int64                               `json:"sfixed64_value,omitempty"`
	Sint32Value                                int32                               `json:"sint32_value,omitempty"`
	Sint64Value                                int64                               `json:"sint64_value,omitempty"`
	RepeatedStringValue                        []string                            `json:"repeated_string_value,omitempty"`
	OneofEmpty                                 *emptypb.Empty                      `json:"oneof_empty,omitempty"`
	OneofString                                *string                             `json:"oneof_string,omitempty"`
	MapValue                                   map[string]NumericEnum              `json:"map_value,omitempty"`
	MappedStringValue                          map[string]string                   `json:"mapped_string_value,omitempty"`
	MappedNestedValue                          map[string]*ABitOfEverything_Nested `json:"mapped_nested_value,omitempty"`
	NonConventionalNameValue                   string                              `json:"nonConventionalNameValue,omitempty"`
	TimestampValue                             *timestamppb.Timestamp              `json:"timestamp_value,omitempty"`
	RepeatedEnumValue                          []NumericEnum                       `json:"repeated_enum_value,omitempty"`
	RepeatedEnumAnnotation                     []NumericEnum                       `json:"repeated_enum_annotation,omitempty"`
	EnumValueAnnotation                        NumericEnum                         `json:"enum_value_annotation,omitempty"`
	RepeatedStringAnnotation                   []string                            `json:"repeated_string_annotation,omitempty"`
	RepeatedNestedAnnotation                   []*ABitOfEverything_Nested          `json:"repeated_nested_annotation,omitempty"`
	NestedAnnotation                           *ABitOfEverything_Nested            `json:"nested_annotation,omitempty"`
	Int64OverrideType                          int64                               `json:"int64_override_type,omitempty"`
	RequiredStringViaFieldBehaviorAnnotation   string                              `json:"required_string_via_field_behavior_annotation,omitempty"`
	OutputOnlyStringViaFieldBehaviorAnnotation string                              `json:"output_only_string_via_field_behavior_annotation,omitempty"`
}

// ABitOfEverythingRepeated is used to validate repeated path parameter functionality
type ABitOfEverythingRepeated struct {
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
	Abe        *ABitOfEverything      `json:"abe,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `json:"update_mask,omitempty"`
}

// An example resource type from AIP-123 used to test the behavior described in
// the CreateBookRequest message.
//
// See: https://google.aip.dev/123
type Book struct {
	Name       string                 `json:"name,omitempty"`
	Id         string                 `json:"id,omitempty"`
	CreateTime *timestamppb.Timestamp `json:"create_time,omitempty"`
}

// A standard Create message from AIP-133 with a user-specified ID.
// The user-specified ID (the `book_id` field in this example) must become a
// query parameter in the OpenAPI spec.
//
// See: https://google.aip.dev/133#user-specified-ids
type CreateBookRequest struct {
	Parent string `json:"parent,omitempty"`
	Book   *Book  `json:"book,omitempty"`
	BookId string `json:"book_id,omitempty"`
}

// A standard Update message from AIP-134
//
// See: https://google.aip.dev/134#request-message
type UpdateBookRequest struct {
	Book         *Book                  `json:"book,omitempty"`
	UpdateMask   *fieldmaskpb.FieldMask `json:"update_mask,omitempty"`
	AllowMissing bool                   `json:"allow_missing,omitempty"`
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
	Name   string                           `json:"name,omitempty"`
	Amount uint32                           `json:"amount,omitempty"`
	Ok     ABitOfEverything_Nested_DeepEnum `json:"ok,omitempty"`
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
	return nil, status.Errorf(status.Unimplemented, "method Create not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CreateBody(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(status.Unimplemented, "method CreateBody not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CreateBook(context.Context, *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(status.Unimplemented, "method CreateBook not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*Book, error) {
	return nil, status.Errorf(status.Unimplemented, "method UpdateBook not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Lookup(context.Context, *IdMessage) (*ABitOfEverything, error) {
	return nil, status.Errorf(status.Unimplemented, "method Lookup not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Update(context.Context, *ABitOfEverything) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method Update not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) UpdateV2(context.Context, *UpdateV2Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method UpdateV2 not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Delete(context.Context, *IdMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method Delete not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) GetQuery(context.Context, *ABitOfEverything) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method GetQuery not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) GetRepeatedQuery(context.Context, *ABitOfEverythingRepeated) (*ABitOfEverythingRepeated, error) {
	return nil, status.Errorf(status.Unimplemented, "method GetRepeatedQuery not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(status.Unimplemented, "method Echo not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) DeepPathEcho(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(status.Unimplemented, "method DeepPathEcho not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) Timeout(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method Timeout not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) ErrorWithDetails(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method ErrorWithDetails not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) GetMessageWithBody(context.Context, *MessageWithBody) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method GetMessageWithBody not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) PostWithEmptyBody(context.Context, *Body) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method PostWithEmptyBody not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckGetQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(status.Unimplemented, "method CheckGetQueryParams not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckNestedEnumGetQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(status.Unimplemented, "method CheckNestedEnumGetQueryParams not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckPostQueryParams(context.Context, *ABitOfEverything) (*ABitOfEverything, error) {
	return nil, status.Errorf(status.Unimplemented, "method CheckPostQueryParams not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) OverwriteResponseContentType(context.Context, *emptypb.Empty) (*StringValue, error) {
	return nil, status.Errorf(status.Unimplemented, "method OverwriteResponseContentType not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckExternalPathEnum(context.Context, *MessageWithPathEnum) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method CheckExternalPathEnum not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckExternalNestedPathEnum(context.Context, *MessageWithNestedPathEnum) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method CheckExternalNestedPathEnum not implemented")
}

func (s *UnimplementedABitOfEverythingServiceServer) CheckStatus(context.Context, *emptypb.Empty) (*CheckStatusResponse, error) {
	return nil, status.Errorf(status.Unimplemented, "method CheckStatus not implemented")
}

type defaultABitOfEverythingServiceDecorator struct {
	ss ABitOfEverythingServiceServer
}

func (s defaultABitOfEverythingServiceDecorator) Create_0(ctx *gin.Context) {
	var req ABitOfEverything

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Create(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CreateBody_0(ctx *gin.Context) {
	var req ABitOfEverything

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CreateBody(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CreateBook_0(ctx *gin.Context) {
	var req CreateBookRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CreateBook(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) UpdateBook_0(ctx *gin.Context) {
	var req UpdateBookRequest

	if err := ctx.ShouldBind(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.UpdateBook(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Lookup_0(ctx *gin.Context) {
	var req IdMessage

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Lookup(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Update_0(ctx *gin.Context) {
	var req ABitOfEverything

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Update(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) UpdateV2_0(ctx *gin.Context) {
	var req UpdateV2Request

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.UpdateV2(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) UpdateV2_1(ctx *gin.Context) {
	var req UpdateV2Request

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.UpdateV2(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) UpdateV2_2(ctx *gin.Context) {
	var req UpdateV2Request

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.UpdateV2(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Delete_0(ctx *gin.Context) {
	var req IdMessage

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Delete(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) GetQuery_0(ctx *gin.Context) {
	var req ABitOfEverything

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.GetQuery(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) GetRepeatedQuery_0(ctx *gin.Context) {
	var req ABitOfEverythingRepeated

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.GetRepeatedQuery(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Echo_0(ctx *gin.Context) {
	var req StringMessage

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Echo(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Echo_1(ctx *gin.Context) {
	var req StringMessage

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Echo(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Echo_2(ctx *gin.Context) {
	var req StringMessage

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Echo(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) DeepPathEcho_0(ctx *gin.Context) {
	var req ABitOfEverything

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.DeepPathEcho(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) Timeout_0(ctx *gin.Context) {
	var req emptypb.Empty

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Timeout(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) ErrorWithDetails_0(ctx *gin.Context) {
	var req emptypb.Empty

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.ErrorWithDetails(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) GetMessageWithBody_0(ctx *gin.Context) {
	var req MessageWithBody

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.GetMessageWithBody(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) PostWithEmptyBody_0(ctx *gin.Context) {
	var req Body

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.PostWithEmptyBody(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckGetQueryParams_0(ctx *gin.Context) {
	var req ABitOfEverything

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CheckGetQueryParams(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckNestedEnumGetQueryParams_0(ctx *gin.Context) {
	var req ABitOfEverything

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CheckNestedEnumGetQueryParams(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckPostQueryParams_0(ctx *gin.Context) {
	var req ABitOfEverything

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CheckPostQueryParams(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) OverwriteResponseContentType_0(ctx *gin.Context) {
	var req emptypb.Empty

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.OverwriteResponseContentType(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckExternalPathEnum_0(ctx *gin.Context) {
	var req MessageWithPathEnum

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CheckExternalPathEnum(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckExternalNestedPathEnum_0(ctx *gin.Context) {
	var req MessageWithNestedPathEnum

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CheckExternalNestedPathEnum(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultABitOfEverythingServiceDecorator) CheckStatus_0(ctx *gin.Context) {
	var req emptypb.Empty

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CheckStatus(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

// RegisterABitOfEverythingServiceServer registers the http handlers for service ABitOfEverythingService to "router".
func RegisterABitOfEverythingServiceServer(router gin.IRouter, s ABitOfEverythingServiceServer) {
	d := defaultABitOfEverythingServiceDecorator{ss: s}
	router.Handle("POST", "/v1/example/a_bit_of_everything/:float_value/:double_value/:int64_value/separator/:uint64_value/:int32_value/:fixed64_value/:fixed32_value/:bool_value/{string_value=strprefix/*}/:uint32_value/:sfixed32_value/:sfixed64_value/:sint32_value/:sint64_value/:nonConventionalNameValue/:enum_value/:path_enum_value/:nested_path_enum_value/:enum_value_annotation", d.Create_0)
	router.Handle("POST", "/v1/example/a_bit_of_everything", d.CreateBody_0)
	router.Handle("POST", "/v1/{parent=publishers/*}/books", d.CreateBook_0)
	router.Handle("PATCH", "/v1/{book.name=publishers/*/books/*}", d.UpdateBook_0)
	router.Handle("GET", "/v1/example/a_bit_of_everything/:uuid", d.Lookup_0)
	router.Handle("PUT", "/v1/example/a_bit_of_everything/:uuid", d.Update_0)
	router.Handle("PUT", "/v2/example/a_bit_of_everything/:abe.uuid", d.UpdateV2_0)
	router.Handle("PATCH", "/v2/example/a_bit_of_everything/:abe.uuid", d.UpdateV2_1)
	router.Handle("PATCH", "/v2a/example/a_bit_of_everything/:abe.uuid", d.UpdateV2_2)
	router.Handle("DELETE", "/v1/example/a_bit_of_everything/:uuid", d.Delete_0)
	router.Handle("GET", "/v1/example/a_bit_of_everything/query/:uuid", d.GetQuery_0)
	router.Handle("GET", "/v1/example/a_bit_of_everything_repeated/:path_repeated_float_value/:path_repeated_double_value/:path_repeated_int64_value/:path_repeated_uint64_value/:path_repeated_int32_value/:path_repeated_fixed64_value/:path_repeated_fixed32_value/:path_repeated_bool_value/:path_repeated_string_value/:path_repeated_bytes_value/:path_repeated_uint32_value/:path_repeated_enum_value/:path_repeated_sfixed32_value/:path_repeated_sfixed64_value/:path_repeated_sint32_value/:path_repeated_sint64_value", d.GetRepeatedQuery_0)
	router.Handle("GET", "/v1/example/a_bit_of_everything/echo/:value", d.Echo_0)
	router.Handle("POST", "/v2/example/echo", d.Echo_1)
	router.Handle("GET", "/v2/example/echo", d.Echo_2)
	router.Handle("POST", "/v1/example/deep_path/:single_nested.name", d.DeepPathEcho_0)
	router.Handle("GET", "/v2/example/timeout", d.Timeout_0)
	router.Handle("GET", "/v2/example/errorwithdetails", d.ErrorWithDetails_0)
	router.Handle("POST", "/v2/example/withbody/:id", d.GetMessageWithBody_0)
	router.Handle("POST", "/v2/example/postwithemptybody/:name", d.PostWithEmptyBody_0)
	router.Handle("GET", "/v1/example/a_bit_of_everything/params/get/:single_nested.name", d.CheckGetQueryParams_0)
	router.Handle("GET", "/v1/example/a_bit_of_everything/params/get/nested_enum/:single_nested.ok", d.CheckNestedEnumGetQueryParams_0)
	router.Handle("POST", "/v1/example/a_bit_of_everything/params/post/:string_value", d.CheckPostQueryParams_0)
	router.Handle("GET", "/v2/example/overwriteresponsecontenttype", d.OverwriteResponseContentType_0)
	router.Handle("GET", "/v2/{value}:check", d.CheckExternalPathEnum_0)
	router.Handle("GET", "/v3/{value}:check", d.CheckExternalNestedPathEnum_0)
	router.Handle("GET", "/v1/example/checkStatus", d.CheckStatus_0)
}

// CamelCaseServiceNameServer is the server API for CamelCaseServiceName service.
type CamelCaseServiceNameServer interface {
	Empty(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

// UnimplementedCamelCaseServiceNameServer can be embedded to have forward compatible implementations.
type UnimplementedCamelCaseServiceNameServer struct{}

func (s *UnimplementedCamelCaseServiceNameServer) Empty(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(status.Unimplemented, "method Empty not implemented")
}

type defaultCamelCaseServiceNameDecorator struct {
	ss CamelCaseServiceNameServer
}

func (s defaultCamelCaseServiceNameDecorator) Empty_0(ctx *gin.Context) {
	var req emptypb.Empty

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(status.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.Empty(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

// RegisterCamelCaseServiceNameServer registers the http handlers for service CamelCaseServiceName to "router".
func RegisterCamelCaseServiceNameServer(router gin.IRouter, s CamelCaseServiceNameServer) {
	d := defaultCamelCaseServiceNameDecorator{ss: s}
	router.Handle("GET", "/v2/example/empty", d.Empty_0)
}
