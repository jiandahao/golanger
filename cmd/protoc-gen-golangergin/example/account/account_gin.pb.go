// Code generated. DO NOT EDIT.

package account

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	fmt "fmt"
	gin "github.com/gin-gonic/gin"
	runtime "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	ioutil "io/ioutil"
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

// register request
type AccountRegister struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

// register status
type RegisterStatus struct {
	Status string `json:"status,omitempty"`
}

// get user profile request
type GetProfileRequest struct {
	UserId     string `json:"user_id,omitempty" uri:"user_id" form:"user_id" validate:"required"`
	CreateTime string `json:"create_time,omitempty" form:"create_time"`
	Token      string `json:"token,omitempty" header:"token"`
}

// user profile
type Profile struct {
	UserId   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Age      int32  `json:"age,omitempty"`
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	// create an account
	CreateAccount(context.Context, *AccountRegister) (*RegisterStatus, error)
	// get user's profile
	GetProfile(context.Context, *GetProfileRequest) (*Profile, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct{}

func (s *UnimplementedAccountServer) CreateAccount(context.Context, *AccountRegister) (*RegisterStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}

func (s *UnimplementedAccountServer) GetProfile(context.Context, *GetProfileRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}

type defaultAccountDecorator struct {
	ss AccountServer
}

func (s defaultAccountDecorator) CreateAccount_0(ctx *gin.Context) {
	var req AccountRegister

	if err := ctx.ShouldBindJSON(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.CreateAccount(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultAccountDecorator) GetProfile_0(ctx *gin.Context) {
	var req GetProfileRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindHeader(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.GetProfile(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

func (s defaultAccountDecorator) GetProfile_1(ctx *gin.Context) {
	var req GetProfileRequest

	if err := ctx.ShouldBindHeader(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		runtime.HTTPError(ctx, status.Errorf(codes.InvalidArgument, err.Error()))
		return
	}

	resp, err := s.ss.GetProfile(ctx, &req)
	if err != nil {
		runtime.HTTPError(ctx, err)
		return
	}

	runtime.ForwardResponseMessage(ctx, resp)
}

// RegisterAccountServer registers the http handlers for service Account to "router".
func RegisterAccountServer(router gin.IRouter, s AccountServer) {
	d := defaultAccountDecorator{ss: s}
	router.Handle("POST", "/v1/auth/signin", d.CreateAccount_0)
	router.Handle("GET", "/v1/user/:user_id/profile", d.GetProfile_0)
	router.Handle("GET", "/v1/user/profiles", d.GetProfile_1)
}

// AccountClient is the client API for for Account service.
type AccountClient interface {
	// create an account
	CreateAccount(context.Context, *AccountRegister) (*RegisterStatus, error)
	// get user's profile
	GetProfile(context.Context, *GetProfileRequest) (*Profile, error)
}

type defaultAccountClient struct {
	cc   *http.Client
	host string
}

// NewAccountClient creates a client API for Account service.
func NewAccountClient(host string, cc *http.Client) AccountClient {
	return &defaultAccountClient{cc: cc, host: strings.TrimSuffix(host, "/")}
}

func (c *defaultAccountClient) CreateAccount(ctx context.Context, req *AccountRegister) (*RegisterStatus, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/auth/signin")

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

	var resp RegisterStatus
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *defaultAccountClient) GetProfile(ctx context.Context, req *GetProfileRequest) (*Profile, error) {
	endpoint := fmt.Sprintf("%s%s", c.host, "/v1/user/:user_id/profile")
	endpoint = strings.ReplaceAll(endpoint, ":user_id", fmt.Sprint(req.UserId))

	hreq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with error: %s", err)
	}

	hreq.Header.Set("Content-Type", "application/json")
	hreq.Header.Add("token", req.Token)

	var queries = url.Values{}
	queries.Add("user_id", req.UserId)
	queries.Add("create_time", req.CreateTime)
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

	var resp Profile
	if err := runtime.BackwardResponseMessage(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
