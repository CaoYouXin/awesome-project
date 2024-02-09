package auth

import (
	"awesomeProject/internal/res"
	"github.com/gin-gonic/gin"
)

func Controller(c *gin.Context) {
	res.Success(c, "欢迎登录！")
}
