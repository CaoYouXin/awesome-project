package res

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type ErrorMap struct {
	errors []error

	statusCode int
	response   func(c *gin.Context, err error)
}

func (e *ErrorMap) StatusCode(statusCode int) *ErrorMap {
	e.statusCode = statusCode
	e.response = func(c *gin.Context, err error) {
		c.Status(statusCode)
	}
	return e
}

func (e *ErrorMap) Response(response func(c *gin.Context, err error)) *ErrorMap {
	e.response = response
	return e
}

func (e *ErrorMap) matchError(actual error) bool {
	for _, expected := range e.errors {
		if errors.Is(actual, expected) {
			return true
		}
	}
	return false
}

func NewErrMap(err ...error) *ErrorMap {
	return &ErrorMap{errors: err}
}
