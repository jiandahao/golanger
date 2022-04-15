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
	FileA *multipart.FileHeader `json:"file_a,omitempty" form:"file_a"`
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

type defaultEchoDecorator struct {
	ss EchoServer
}

func (s defaultEchoDecorator) GetEcho_0(ctx *gin.Context) {
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

	resp, err := s.ss.GetEcho(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultEchoDecorator) GetEcho_1(ctx *gin.Context) {
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

	resp, err := s.ss.GetEcho(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultEchoDecorator) PostEcho_0(ctx *gin.Context) {
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

	resp, err := s.ss.PostEcho(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultEchoDecorator) PostEcho_1(ctx *gin.Context) {
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

	resp, err := s.ss.PostEcho(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultEchoDecorator) PostFormEcho_0(ctx *gin.Context) {
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

	resp, err := s.ss.PostFormEcho(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

// RegisterEchoServer registers the http handlers for service Echo to "router".
func RegisterEchoServer(router gin.IRouter, s EchoServer) {
	d := defaultEchoDecorator{ss: s}
	router.Handle("GET", "/api/v1/echo", d.GetEcho_0)
	router.Handle("GET", "/api/v1/echo/:param_in_uri_or_query", d.GetEcho_1)
	router.Handle("POST", "/api/v1/echo", d.PostEcho_0)
	router.Handle("POST", "/api/v1/echo/:param_in_uri_or_query", d.PostEcho_1)
	router.Handle("POST", "/api/v1/form", d.PostFormEcho_0)
}

// EchoClient is the client API for for Echo service.
type EchoClient interface {
	GetEcho(context.Context, *GetEchoReq) (*GetEchoResp, error)
	PostEcho(context.Context, *PostEchoReq) (*PostEchoResp, error)
	PostFormEcho(context.Context, *PostFormEchoReq) (*PostFormEchoResp, error)
}

type defaultEchoClient struct {
	cc   *http.Client
	host string
}

// NewEchoClient creates a client API for Echo service.
func NewEchoClient(host string, cc *http.Client) EchoClient {
	return &defaultEchoClient{cc: cc, host: strings.TrimSuffix(host, "/")}
}

func (c *defaultEchoClient) GetEcho(ctx context.Context, req *GetEchoReq) (*GetEchoResp, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/api/v1/echo")

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	hreq.Header.Add("param_in_header_or_query", req.ParamInHeaderOrQuery)

	var queries = url.Values{}
	queries.Add("param_in_uri_or_query", req.ParamInUriOrQuery)
	queries.Add("param_in_header_or_query", req.ParamInHeaderOrQuery)
	hreq.URL.RawQuery = queries.Encode()

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp GetEchoResp
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultEchoClient) PostEcho(ctx context.Context, req *PostEchoReq) (*PostEchoResp, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/api/v1/echo")

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request with error: %s", err)
	}

	hreq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")

	hreq.Header.Add("param_in_header", req.ParamInHeader)

	var queries = url.Values{}
	queries.Add("param_in_uri_or_query", req.ParamInUriOrQuery)
	hreq.URL.RawQuery = queries.Encode()

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp PostEchoResp
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultEchoClient) PostFormEcho(ctx context.Context, req *PostFormEchoReq) (*PostFormEchoResp, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/api/v1/form")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("param_in_form_a", req.ParamInFormA)
	writer.WriteField("param_in_form_b", req.ParamInFormB)

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

	if err := writer.Close(); err != nil {
		return nil, err
	}

	hreq, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := c.cc.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp PostFormEchoResp
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
