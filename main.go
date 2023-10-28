package main

import (
	"github.com/gin-gonic/gin"
	management "github.com/juanmabm/focusio-core/management/http"
)

func SetupHttpServer() *gin.Engine {

	r := gin.Default()

	mh := management.NewHandler()
	management.Routes(r, mh)

	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	return r
}

func main() {
	r := SetupHttpServer()
	r.Run(":8080")
}
