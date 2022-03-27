package runtime

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPError eplies to the request with an error.
var HTTPError = defaultHTTPErrorHandler

// defaultHTTPErrorHandler // DefaultHTTPError is the default implementation of HTTPError.
func defaultHTTPErrorHandler(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, err.Error())
}

// ForwardResponseMessage forwards the message "resp" from server to REST client.
func ForwardResponseMessage(ctx *gin.Context, resp interface{}) {
	ctx.JSON(http.StatusOK, resp)
}
