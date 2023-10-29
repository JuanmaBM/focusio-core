package main

import (
	"github.com/gin-gonic/gin"
	mng "github.com/juanmabm/focusio-core/management/http"
)

func SetupHttpServer() *gin.Engine {

	r := gin.Default()

	fah := mng.NewFocusAppHandler()
	fch := mng.NewFocusCatalogItemHandler()
	mng.Routes(r, fah, fch)

	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	return r
}

func main() {
	r := SetupHttpServer()
	r.Run(":8080")
}
