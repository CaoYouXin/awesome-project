package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSON struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, JSON{
		Code: 200,
		Data: data,
	})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusOK, JSON{
		Code: 500,
		Msg:  err.Error(),
		Data: err,
	})
}
