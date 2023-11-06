package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoConnection() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = c.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB [mongodb://localhost:27017]")
	return c, err
}

func CreateIndexes(c *mongo.Client) {
	database := c.Database("focusio")
	createIndex(database, "FocusApp", "name")
	createIndex(database, "FocusCatalogItem", "name")
}

func createIndex(database *mongo.Database, collection string, index string) {
	_, err := database.Collection(collection).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: index, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		panic(err)
	}
}
