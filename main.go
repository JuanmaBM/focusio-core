package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/internal/database"
	"github.com/juanmabm/focusio-core/internal/focusapp"
	"github.com/juanmabm/focusio-core/internal/focuscatalog"
	"go.mongodb.org/mongo-driver/mongo"
)

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

	var hostname = os.Getenv("MONGO_HOSTNAME")
	var username = os.Getenv("MONGO_USERNAME")
	var password = os.Getenv("MONGO_PASSWORD")
	var dbName = "focusio"

	dbConnection, err := database.CreateMongoConnection(hostname, username, password, dbName)
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
