package ping

import (
	"awesomeProject/internal/res"
	"github.com/gin-gonic/gin"
)

func Controller(c *gin.Context) {
	// model计算
	counter++

	// 创建view
	var data = Pong{
		Msg:   "Pong",
		Count: counter,
	}

	// 框架返回结果
	res.Success(c, data)
}
