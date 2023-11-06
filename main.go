package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/management/database"
	"github.com/juanmabm/focusio-core/management/focusapp"
)

func SetupHttpServer() *gin.Engine {

	r := gin.Default()
	mongoConnection, err := database.CreateMongoConnection()
	if err != nil {
		panic(err.Error())
	}

	far := focusapp.NewFocusAppRepository(mongoConnection)
	focusapp.RegisterHandlers(r, far)

	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	return r
}

func main() {
	r := SetupHttpServer()
	r.Run(":8080")
}
