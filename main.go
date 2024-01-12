package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/management/database"
	"github.com/juanmabm/focusio-core/management/focusapp"
	"github.com/juanmabm/focusio-core/management/focuscatalog"
	"go.mongodb.org/mongo-driver/mongo"
)

var hostname = "mongodb://localhost:27017"
var dbName = "focusio"

func SetupHttpServer(dbConnection *mongo.Database) *gin.Engine {

	r := gin.Default()

	far := focusapp.NewFocusAppRepository(dbConnection)
	fcr := focuscatalog.NewFocusCatalogItemRepository(dbConnection)
	focusapp.RegisterHandlers(r, far, fcr)
	focuscatalog.RegisterHandlers(r, fcr)

	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	return r
}

func CreateDatabaseConnection() *mongo.Database {

	dbConnection, err := database.CreateMongoConnection(hostname, dbName)
	if err != nil {
		panic(err.Error())
	}

	database.CreateIndexes(dbConnection)
	return dbConnection
}

func main() {
	dbConnection := CreateDatabaseConnection()
	httpServer := SetupHttpServer(dbConnection)
	httpServer.Run(":8080")
}
