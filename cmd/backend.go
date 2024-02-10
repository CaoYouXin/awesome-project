package cmd

import (
	"awesomeProject/internal/auth"
	"awesomeProject/internal/jgg"
	"awesomeProject/internal/ping"
	"awesomeProject/internal/res"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(res.ErrorHandler())

	// Ping test
	r.GET("/ping", ping.Controller)

	authorized := r.Group("/")
	authorized.Use(auth.BasicAuth(&gin.Accounts{
		"foo": "bar",
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'

		go:
		user := c.MustGet(gin.AuthUserKey).(string)
	*/
	authorized.POST("auth", auth.Controller)

	// 九宫格
	authorized.POST("/birthday", jgg.SetBirthDay)

	return r
}

func StartServer(port string) net.Listener {
	ln, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err = http.Serve(ln, setupRouter()); err != nil {
			log.Fatal(err)
		}
	}()

	return ln
}
