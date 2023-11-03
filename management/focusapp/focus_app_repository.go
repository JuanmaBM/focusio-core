package focusapp

import (
	"context"

	"github.com/juanmabm/focusio-core/management/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FocusAppRepository interface {
	insert(fa *entity.FocusApp)
	findByName(n string) entity.FocusApp
	findAll() []entity.FocusApp
	delete(n string)
	update(n string, fa *entity.FocusApp)
}

type focusAppRepository struct {
	col *mongo.Collection
}

func NewFocusAppRepository(c *mongo.Client) focusAppRepository {
	return focusAppRepository{
		c.Database("focusio").Collection("FocusApp"),
	}
}

func (far focusAppRepository) insert(fa *entity.FocusApp) {
	far.col.InsertOne(context.TODO(), fa)
}

func (far focusAppRepository) findByName(n string) entity.FocusApp {
	var fa entity.FocusApp
	filter := bson.D{{"name", n}}
	result := far.col.FindOne(context.TODO(), filter)
	if result != nil {
		result.Decode(&fa)
	}

	return fa
}

func (far focusAppRepository) findAll() []entity.FocusApp {
	cursor, err := far.col.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var fas []entity.FocusApp
	if err = cursor.All(context.TODO(), &fas); err != nil {
		panic(err)
	}

	return fas
}

func (far focusAppRepository) delete(n string) {
	filter := bson.D{{"name", n}}
	far.col.DeleteOne(context.TODO(), filter)
}

func (far focusAppRepository) update(n string, fa *entity.FocusApp) {
	filter := bson.D{{"name", n}}
	far.col.UpdateOne(context.TODO(), filter, fa)
}
