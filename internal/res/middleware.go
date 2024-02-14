package res

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorHandler(errM ...*ErrorMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastError := c.Errors.Last()
		if lastError == nil {
			return
		}

		for _, err := range errM {
			if err.matchError(lastError.Err) {
				err.response(c, lastError.Err)
			}
		}
	}
}

func ErrorHandler() gin.HandlerFunc {
	err1 := NewErrMap(badRequest).StatusCode(http.StatusOK).Response(func(c *gin.Context, err error) {
		fmt.Println("Wow 400", err.Error())
		c.JSON(http.StatusOK, JSON{
			Code: 400,
			Msg:  err.Error(),
		})
	})

	err2 := NewErrMap(internal).StatusCode(http.StatusOK).Response(func(c *gin.Context, err error) {
		fmt.Println("Wow 500", err.Error())
		c.JSON(http.StatusOK, JSON{
			Code: 500,
			Msg:  err.Error(),
		})
	})

	return errorHandler(err1, err2)
}
