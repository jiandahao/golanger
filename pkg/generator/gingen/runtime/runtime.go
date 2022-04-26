package runtime

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiandahao/golanger/pkg/generator/gingen/status"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

// HTTPError eplies to the request with an error.
var HTTPError = defaultHTTPErrorHandler

// ForwardResponseMessage forwards the message "resp" from server to REST client.
var ForwardResponseMessage = forwardResponseMessage

// BackwardResponseMessage backwards message "resp" from REST client to client API.
var BackwardResponseMessage = backwardResponseMessage

// HideDetails represents whether hide the error details or not.
var HideDetails bool

// defaultHTTPErrorHandler is the default implementation of HTTPError.
func defaultHTTPErrorHandler(ctx *gin.Context, err error) {
	if s, ok := grpcStatus.FromError(err); ok {
		e, ok := status.FromCode(s.Code())
		if !ok {
			e, _ = status.FromCode(status.Internal)
		}

		err = &status.ErrorDetails{Code: e.Code, Msg: e.Msg, Status: e.Status, Details: err.Error()}
	}

	e, ok := err.(*status.ErrorDetails)
	if !ok {
		e, _ = status.FromCode(status.Internal)
		e.Details = err.Error()
	}

	data := responseData{
		Code:    e.Code,
		Msg:     e.Msg,
		Details: e.Details,
	}

	if HideDetails {
		data.Details = ""
	}

	ctx.JSON(e.Status, data)
}

// forwardResponseMessage forwards the message "resp" from server to REST client.
func forwardResponseMessage(ctx *gin.Context, resp interface{}) {
	ctx.JSON(http.StatusOK, responseData{
		Code: codes.OK,
		Msg:  "ok",
		Data: resp,
	})
}

// backwardResponseMessage backwards message "resp" from REST client to client API.
func backwardResponseMessage(body []byte, resp interface{}) error {
	var res responseData
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	if res.Code != codes.OK {
		return &status.ErrorDetails{Code: res.Code, Msg: res.Msg, Details: res.Details}
	}

	data, err := json.Marshal(res.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, resp)
}

type responseData struct {
	Code    codes.Code  `json:"code"`              // error code
	Msg     string      `json:"msg"`               // error short message
	Details string      `json:"details,omitempty"` // error details if possible
	Data    interface{} `json:"data,omitempty"`
}

// runtimeContext represents a runtime context.
type runtimeContext struct {
	ginCtx *gin.Context
}

// NewContext new context
func NewContext(c *gin.Context) context.Context {
	return &runtimeContext{ginCtx: c}
}

func (c *runtimeContext) Deadline() (deadline time.Time, ok bool) {
	return c.ginCtx.Request.Context().Deadline()
}

func (c *runtimeContext) Done() <-chan struct{} {
	return c.ginCtx.Request.Context().Done()
}

func (c *runtimeContext) Err() error {
	return c.ginCtx.Request.Context().Err()
}

func (c *runtimeContext) Value(key interface{}) interface{} {
	val := c.ginCtx.Value(key)
	if val != nil {
		return val
	}

	return c.ginCtx.Request.Context().Value(key)
}
