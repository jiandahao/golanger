package runtime

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiandahao/golanger/pkg/generator/gingen/status"
	grpcStatus "google.golang.org/grpc/status"
)

// HTTPError eplies to the request with an error.
var HTTPError = defaultHTTPErrorHandler

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
	}

	data := responseData{
		Code:    int32(e.Code),
		Msg:     e.Msg,
		Details: e.Details,
	}

	if HideDetails {
		data.Details = ""
	}

	ctx.JSON(e.Status, data)
}

// ForwardResponseMessage forwards the message "resp" from server to REST client.
func ForwardResponseMessage(ctx *gin.Context, resp interface{}) {
	ctx.JSON(http.StatusOK, responseData{
		Code: 0,
		Msg:  "ok",
		Data: resp,
	})
}

type responseData struct {
	Code    int32       `json:"code"`
	Msg     string      `json:"msg"`
	Details string      `json:"details,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
