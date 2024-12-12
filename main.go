package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/internal/database"
	"github.com/juanmabm/focusio-core/internal/focusapp"
	"github.com/juanmabm/focusio-core/internal/focuscatalog"
	"github.com/juanmabm/focusio-core/pkg/argocdclient"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupHttpServer() *gin.Engine {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	return r
}

func createDatabaseConnection() *mongo.Database {

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

func createArgocdConnection() *argocdclient.ArgoCDClient {

	var serverAddr = os.Getenv("ARGOCD_SERVER_ADDR")
	var port = os.Getenv("ARGOCD_SERVER_PORT")
	var username = os.Getenv("ARGOCD_USERNAME")
	var password = os.Getenv("ARGOCD_PASSWORD")
	var insecure bool = os.Getenv("ARGOCD_CONNECTION_INSECURE") == "true"

	ac, err := argocdclient.NewArgoCDClient(serverAddr, port, username, password, insecure)
	if err != nil {
		panic(err)
	}
	return ac
}

func main() {
	dbConnection := createDatabaseConnection()
	// TODO: Add Argocd client to handlers
	// ac := createArgocdConnection()
	r := setupHttpServer()

	far := focusapp.NewFocusAppRepository(dbConnection)
	fcr := focuscatalog.NewFocusCatalogItemRepository(dbConnection)
	focusapp.RegisterHandlers(r, far, fcr)
	focuscatalog.RegisterHandlers(r, fcr)

	r.Run(":8080")
}
