package focuscatalog

import (
	"context"

	"github.com/juanmabm/focusio-core/management/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FocusCatalogItemRepository interface {
	FindAll() []entity.FocusCatalogItem
	FindByName(name string) (entity.FocusCatalogItem, error)
	Insert(catalog entity.FocusCatalogItem) error
	Update(name string, catalog entity.FocusCatalogItem) error
	Delete(name string) error
}

type focusCatalogItemRepository struct {
	col *mongo.Collection
}

func NewFocusCatalogItemRepository(db *mongo.Database) FocusCatalogItemRepository {
	return focusCatalogItemRepository{
		db.Collection("FocusCatalogItem"),
	}
}

func (r focusCatalogItemRepository) FindAll() []entity.FocusCatalogItem {
	cursor, err := r.col.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	items := []entity.FocusCatalogItem{}
	if err = cursor.All(context.Background(), &items); err != nil {
		panic(err)
	}

	return items
}

func (r focusCatalogItemRepository) FindByName(name string) (entity.FocusCatalogItem, error) {
	var item entity.FocusCatalogItem
	filterByName := bson.D{{Key: "name", Value: name}}

	result := r.col.FindOne(context.Background(), filterByName)
	if result.Err() != nil {
		return item, result.Err()
	}

	if err := result.Decode(&item); err != nil {
		return item, err
	}

	return item, nil
}

func (r focusCatalogItemRepository) Insert(item entity.FocusCatalogItem) error {
	_, err := r.col.InsertOne(context.Background(), item)
	return err
}

func (r focusCatalogItemRepository) Update(name string, item entity.FocusCatalogItem) error {
	filter := bson.D{{Key: "name", Value: name}}
	_, err := r.col.ReplaceOne(context.Background(), filter, item)
	return err
}

func (r focusCatalogItemRepository) Delete(name string) error {
	filter := bson.D{{Key: "name", Value: name}}
	_, err := r.col.DeleteOne(context.Background(), filter)
	return err
}
