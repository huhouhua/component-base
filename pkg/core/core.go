package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huhouhua/errors"
	"github.com/huhouhua/log"
)

// SvcResponse defines the return messages when an error occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type SvcResponse[T any] struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	// Reference returns the reference document which maybe useful to solve this error.
	Reference string `json:"reference,omitempty"`

	Data T `json:"data"`
}

func newSvcResponse[T any](code int, message string, reference string, data T) *SvcResponse[T] {
	resp := SvcResponse[T]{
		code,
		message,
		reference,
		data,
	}
	return &resp
}

func Success(data any) *SvcResponse[any] {
	return newSvcResponse[any](http.StatusOK, "操作成功！", "", data)
}
func SuccessWithGeneric[T any](data T) *SvcResponse[T] {
	return newSvcResponse(http.StatusOK, "操作成功！", "", data)
}

func Error(code int, message string, reference string) *SvcResponse[any] {
	return newSvcResponse[any](code, message, reference, "")
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), Error(coder.Code(), coder.String(), coder.Reference()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func WriteResponseSuccess(c *gin.Context, resp *SvcResponse[any]) {
	WriteResponse(c, nil, resp)
}
func WriteResponseError(c *gin.Context, err error) {
	WriteResponse(c, err, nil)
}
