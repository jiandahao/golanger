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
func forwardResponseMessage(ctx *Context, resp interface{}) {
	headers, ok := GetResponseHeader(ctx)
	if ok {
		for key, values := range headers {
			for _, value := range values {
				ctx.Writer.Header().Add(key, value)
			}
		}
	}

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

// Context represents a runtime context
type Context struct {
	*gin.Context
}

// NewContext new context
func NewContext(c *gin.Context) *Context {
	var header = http.Header{}
	for key, values := range c.Request.Header {
		for _, value := range values {
			header.Add(key, value)
		}
	}

	ctx := NewRequestHeaderContext(c.Request.Context(), header)
	ctx = NewResponseHeaderContext(ctx, http.Header{})
	ctx = context.WithValue(ctx, clientIPKey{}, c.ClientIP())

	c.Request = c.Request.WithContext(ctx)

	return &Context{Context: c}
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Request.Context().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.Request.Context().Done()
}

func (c *Context) Err() error {
	return c.Request.Context().Err()
}

func (c *Context) Value(key interface{}) interface{} {
	if val := c.Context.Value(key); val != nil {
		return val
	}

	return c.Request.Context().Value(key)
}

type requestHeaderKey struct{}
type reponseHeaderKey struct{}

// NewRequestHeaderContext creates a new context with request header attached.
func NewRequestHeaderContext(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, requestHeaderKey{}, &header)
}

// NewResponseHeaderContext creates a new context with response header attached.
func NewResponseHeaderContext(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, reponseHeaderKey{}, &header)
}

// SendResponseHeader add reponse header into the context.
func SendResponseHeader(ctx context.Context, header http.Header) {
	addHeaderIntoContext(ctx, reponseHeaderKey{}, header)
}

// GetResponseHeader returns response header attached in the context.
func GetResponseHeader(ctx context.Context) (http.Header, bool) {
	return getHeaderFromContext(ctx, reponseHeaderKey{})
}

// GetRequestHeader returns request header attached in the context.
func GetRequestHeader(ctx context.Context) (http.Header, bool) {
	return getHeaderFromContext(ctx, requestHeaderKey{})
}

func addHeaderIntoContext(ctx context.Context, key interface{}, header http.Header) {
	ctxHeader, ok := ctx.Value(key).(*http.Header)
	if !ok || ctxHeader == nil {
		return
	}

	for key, values := range header {
		for _, value := range values {
			ctxHeader.Add(key, value)
		}
	}
}

func getHeaderFromContext(ctx context.Context, key interface{}) (http.Header, bool) {
	header, ok := ctx.Value(reponseHeaderKey{}).(*http.Header)
	if !ok || header == nil {
		return nil, false
	}

	return *header, true
}

type clientIPKey struct{}

// ClientIPFromContext returns the remote client ip in ctx if it exists.
func ClientIPFromContext(ctx context.Context) (string, bool) {
	clientIP, ok := ctx.Value(clientIPKey{}).(string)
	return clientIP, ok
}
