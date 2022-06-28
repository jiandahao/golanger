// Code generated. DO NOT EDIT.

package echo

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	fmt "fmt"
	gin "github.com/gin-gonic/gin"
	runtime "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetEchoReq struct {
	ParamInUriOrQuery    string `json:"param_in_uri_or_query,omitempty" uri:"param_in_uri_or_query" form:"param_in_uri_or_query"`
	ParamInHeaderOrQuery string `json:"param_in_header_or_query,omitempty" header:"param_in_header_or_query" form:"param_in_header_or_query"`
}

type GetEchoResp struct {
	ParamInUriOrQuery    string `json:"param_in_uri_or_query,omitempty"`
	ParamInHeaderOrQuery string `json:"param_in_header_or_query,omitempty"`
}

type PostEchoReq struct {
	ParamInUriOrQuery string `json:"param_in_uri_or_query,omitempty" uri:"param_in_uri_or_query" form:"param_in_uri_or_query"`
	ParamInHeader     string `json:"param_in_header,omitempty" header:"param_in_header"`
	ParamInBody       string `json:"param_in_body,omitempty"`
}

type PostEchoResp struct {
	ParamInUriOrQuery string `json:"param_in_uri_or_query,omitempty"`
	ParamInHeader     string `json:"param_in_header,omitempty"`
	ParamInBody       string `json:"param_in_body,omitempty"`
}

type PostFormEchoReq struct {
	ParamInFormA string `json:"param_in_form_a,omitempty" form:"param_in_form_a"`
	ParamInFormB string `json:"param_in_form_b,omitempty" form:"param_in_form_b"`
	// Nested Map for all multipart files.
	// Keys of the outer map and inner map represent form-data keys and filename, respectively.
	// Note: Just for Client use only
	MultipartFiles map[string]map[string]io.Reader `json:"multipart_files" form:"multipart_files"`
	// Note: Just for Server use only
	FilesA []*multipart.FileHeader `json:"files_a,omitempty" form:"files_a"`
	// Note: Just for Server use only
	FileB *multipart.FileHeader `json:"file_b,omitempty" form:"file_b"`
}

type PostFormEchoResp struct {
	ParamInFormA string `json:"param_in_form_a,omitempty"`
	ParamInFormB string `json:"param_in_form_b,omitempty"`
	FilenameA    string `json:"filename_a,omitempty"`
	FilenameB    string `json:"filename_b,omitempty"`
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	GetEcho(context.Context, *GetEchoReq) (*GetEchoResp, error)
	PostEcho(context.Context, *PostEchoReq) (*PostEchoResp, error)
	PostFormEcho(context.Context, *PostFormEchoReq) (*PostFormEchoResp, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct{}

func (s *UnimplementedEchoServer) GetEcho(context.Context, *GetEchoReq) (*GetEchoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcho not implemented")
}

func (s *UnimplementedEchoServer) PostEcho(context.Context, *PostEchoReq) (*PostEchoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostEcho not implemented")
}

func (s *UnimplementedEchoServer) PostFormEcho(context.Context, *PostFormEchoReq) (*PostFormEchoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostFormEcho not implemented")
}

// DefaultEchoDecorator the default decorator.
type DefaultEchoDecorator struct {
	ss EchoServer
}

// NewDefaultEchoDecorator constructs a new default Echo decorator
func NewDefaultEchoDecorator(ss EchoServer) *DefaultEchoDecorator {
	return &DefaultEchoDecorator{ss: ss}
}

func (s *DefaultEchoDecorator) GetEcho(ctx *gin.Context) {
	var req GetEchoReq

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindHeader,
		ctx.ShouldBindQuery,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.GetEcho(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}
func (s *DefaultEchoDecorator) GetEcho_1(ctx *gin.Context) {
	var req GetEchoReq

	bindingHandlers := []func(obj interface{}) error{
		ctx.ShouldBindUri,
		ctx.ShouldBindHeader,
		ctx.ShouldBindQuery,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.GetEcho(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s *DefaultEchoDecorator) PostEcho(ctx *gin.Context) {
	var req PostEchoReq
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
		ctx.ShouldBindHeader,
		ctx.ShouldBindQuery,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.PostEcho(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}
func (s *DefaultEchoDecorator) PostEcho_1(ctx *gin.Context) {
	var req PostEchoReq
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
		ctx.ShouldBindHeader,
		ctx.ShouldBindQuery,
	}

	for _, doBinding := range bindingHandlers {
		if err := doBinding(&req); err != nil {
			runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
			return
		}
	}

	newCtx := runtime.NewContext(ctx)
	resp, err := s.ss.PostEcho(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

func (s *DefaultEchoDecorator) PostFormEcho(ctx *gin.Context) {
	var req PostFormEchoReq
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
	resp, err := s.ss.PostFormEcho(newCtx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(newCtx, resp)
}

// RegisterEchoServer registers the http handlers for service Echo to "router".
func RegisterEchoServer(router gin.IRouter, s EchoServer) {
	d := &DefaultEchoDecorator{ss: s}
	router.Handle("GET", "/api/v1/echo", d.GetEcho)
	router.Handle("GET", "/api/v1/echo/:param_in_uri_or_query", d.GetEcho_1)
	router.Handle("POST", "/api/v1/echo", d.PostEcho)
	router.Handle("POST", "/api/v1/echo/:param_in_uri_or_query", d.PostEcho_1)
	router.Handle("POST", "/api/v1/form", d.PostFormEcho)
}

// All Endpoints
var (
	GetEchoEndpoint      = "/api/v1/echo"
	GetEchoEndpoint_1    = "/api/v1/echo/:param_in_uri_or_query"
	PostEchoEndpoint     = "/api/v1/echo"
	PostEchoEndpoint_1   = "/api/v1/echo/:param_in_uri_or_query"
	PostFormEchoEndpoint = "/api/v1/form"
)
