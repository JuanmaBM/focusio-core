package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoConnection(hostname string, database string) (*mongo.Database, error) {

	clientOptions := options.Client().ApplyURI(hostname)
	c, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = c.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB [%s]", hostname)

	db := c.Database(database)
	err = db.Client().Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connect to dabase [%s]", database)

	return db, err
}

func CreateIndexes(database *mongo.Database) {
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
