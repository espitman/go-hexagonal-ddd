package mongodb

// implement item domain repository

import (
	"context"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type itemRepository struct {
	client *mongo.Client
	dbName string
}

func NewItemRepository(client *mongo.Client, dbName string) repositories.ItemRepository {
	return &itemRepository{
		client: client,
		dbName: dbName,
	}
}

func (r *itemRepository) getCollection() *mongo.Collection {
	return r.client.Database(r.dbName).Collection("items")
}

func (r *itemRepository) Create(item *models.Item) (*models.Item, error) {
	item.CreatedAt = time.Now()
	newItem, err := r.getCollection().InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}
	return &models.Item{
		ID:        newItem.InsertedID.(primitive.ObjectID).Hex(),
		ListId:    item.ListId,
		ItemCode:  item.ItemCode,
		CreatedAt: item.CreatedAt,
	}, nil
}

func (r *itemRepository) GetItemsByListID(id string) ([]*models.Item, error) {
	docID, _ := primitive.ObjectIDFromHex(id)
	opts := options.Find()
	cursor, err := r.getCollection().Find(context.Background(), bson.M{"_id": docID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var items []*models.Item
	for cursor.Next(context.Background()) {
		var item models.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *itemRepository) Delete(id string) error {
	docID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.getCollection().DeleteOne(context.Background(), bson.M{"_id": docID})
	return err
}

func (r *itemRepository) IsMemberOfTheList(itemCode int64, listId primitive.ObjectID) bool {
	count, _ := r.getCollection().CountDocuments(context.Background(), bson.M{"itemCode": itemCode, "listId": listId})
	return count != 0
}
