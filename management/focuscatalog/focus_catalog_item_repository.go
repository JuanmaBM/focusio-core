package focuscatalog

import (
	"context"

	"github.com/juanmabm/focusio-core/management/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FocusCatalogItemRepository interface {
	findAll() []entity.FocusCatalogItem
	findByName(name string) (entity.FocusCatalogItem, error)
	insert(catalog entity.FocusCatalogItem) error
	update(name string, catalog entity.FocusCatalogItem) error
	delete(name string) error
}

type focusCatalogItemRepository struct {
	col *mongo.Collection
}

func NewFocusCatalogItemRepository(db *mongo.Database) FocusCatalogItemRepository {
	return focusCatalogItemRepository{
		db.Collection("FocusCatalogItem"),
	}
}

func (r focusCatalogItemRepository) findAll() []entity.FocusCatalogItem {
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

func (r focusCatalogItemRepository) findByName(name string) (entity.FocusCatalogItem, error) {
	var item entity.FocusCatalogItem
	filterByName := bson.D{{"name", name}}

	result := r.col.FindOne(context.Background(), filterByName)
	if result.Err() != nil {
		return item, result.Err()
	}

	if err := result.Decode(&item); err != nil {
		return item, err
	}

	return item, nil
}

func (r focusCatalogItemRepository) insert(item entity.FocusCatalogItem) error {
	_, err := r.col.InsertOne(context.Background(), item)
	return err
}

func (r focusCatalogItemRepository) update(name string, item entity.FocusCatalogItem) error {
	filter := bson.D{{"name", name}}
	_, err := r.col.ReplaceOne(context.Background(), filter, item)
	return err
}

func (r focusCatalogItemRepository) delete(name string) error {
	filter := bson.D{{"name", name}}
	_, err := r.col.DeleteOne(context.Background(), filter)
	return err
}
