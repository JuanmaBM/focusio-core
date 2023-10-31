package database

import (
	"context"

	"github.com/juanmabm/focusio-core/management/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var col *mongo.Collection

type FocusAppRepository struct{}

func NewFocusAppRepository(c *mongo.Client) {
	col = c.Database("focusio").Collection("FocusApp")
}

func (*FocusAppRepository) insert(fa *domain.FocusApp) {
	col.InsertOne(context.Background(), fa)
}

func (*FocusAppRepository) findByName(n string) domain.FocusApp {
	var fa domain.FocusApp
	filter := bson.D{{"name", n}}
	result := col.FindOne(context.TODO(), filter)
	if result != nil {
		result.Decode(&fa)
	}

	return fa
}

func (*FocusAppRepository) findAll() []domain.FocusApp {
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var fas []domain.FocusApp
	if err = cursor.All(context.TODO(), &fas); err != nil {
		panic(err)
	}

	return fas
}
