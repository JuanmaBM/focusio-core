package focusapp

import (
	"context"

	"github.com/juanmabm/focusio-core/management/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FocusAppRepository interface {
	Insert(fa entity.FocusApp) error
	FindByName(n string) (entity.FocusApp, error)
	FindAll() []entity.FocusApp
	Delete(n string) error
	Update(n string, fa *entity.FocusApp) error
}

type focusAppRepository struct {
	col *mongo.Collection
}

func NewFocusAppRepository(db *mongo.Database) FocusAppRepository {
	return focusAppRepository{
		db.Collection("FocusApp"),
	}
}

func (far focusAppRepository) Insert(fa entity.FocusApp) error {
	_, err := far.col.InsertOne(context.TODO(), fa)
	return err
}

func (far focusAppRepository) FindByName(n string) (entity.FocusApp, error) {
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

func (far focusAppRepository) FindAll() []entity.FocusApp {
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

func (far focusAppRepository) Delete(n string) error {
	filter := bson.D{{"name", n}}
	_, err := far.col.DeleteOne(context.TODO(), filter)
	return err
}

func (far focusAppRepository) Update(n string, fa *entity.FocusApp) error {
	filter := bson.D{{"name", n}}
	_, err := far.col.ReplaceOne(context.TODO(), filter, fa)
	return err
}
