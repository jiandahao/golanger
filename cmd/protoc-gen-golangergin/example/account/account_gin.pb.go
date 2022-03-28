// Code generated. DO NOT EDIT.

package account

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	runtime "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

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
