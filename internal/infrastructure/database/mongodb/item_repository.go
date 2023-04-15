package mongodb

// implement item domain repository

import (
	"context"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *itemRepository) GetByListID(id string) (*models.Item, error) {
	var item models.Item
	docID, _ := primitive.ObjectIDFromHex(id)
	err := r.getCollection().FindOne(context.Background(), bson.M{"_id": docID}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
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
