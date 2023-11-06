package focusapp

import (
	"context"

	"github.com/juanmabm/focusio-core/management/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FocusAppRepository interface {
	insert(fa entity.FocusApp) error
	findByName(n string) (entity.FocusApp, error)
	findAll() []entity.FocusApp
	delete(n string) error
	update(n string, fa *entity.FocusApp) error
}

type focusAppRepository struct {
	col *mongo.Collection
}

func NewFocusAppRepository(c *mongo.Client) FocusAppRepository {
	return focusAppRepository{
		c.Database("focusio").Collection("FocusApp"),
	}
}

func (far focusAppRepository) insert(fa entity.FocusApp) error {
	_, err := far.col.InsertOne(context.TODO(), fa)
	return err
}

func (far focusAppRepository) findByName(n string) (entity.FocusApp, error) {
	var fa entity.FocusApp
	filter := bson.D{{"name", n}}

	result := far.col.FindOne(context.TODO(), filter)
	if result.Err() != nil {
		return fa, result.Err()
	}

	if err := result.Decode(&fa); err != nil {
		return fa, err
	}

	return fa, nil
}

func (far focusAppRepository) findAll() []entity.FocusApp {
	cursor, err := far.col.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var fas []entity.FocusApp = []entity.FocusApp{}
	if err = cursor.All(context.TODO(), &fas); err != nil {
		panic(err)
	}

	return fas
}

func (far focusAppRepository) delete(n string) error {
	filter := bson.D{{"name", n}}
	_, err := far.col.DeleteOne(context.TODO(), filter)
	return err
}

func (far focusAppRepository) update(n string, fa *entity.FocusApp) error {
	filter := bson.D{{"name", n}}
	_, err := far.col.ReplaceOne(context.TODO(), filter, fa)
	return err
}
