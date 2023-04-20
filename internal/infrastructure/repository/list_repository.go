package repository

// implement list domain repository

import (
	"context"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type listRepository struct {
	client *mongo.Client
	dbName string
}

func NewListRepository(client *mongo.Client, dbName string) repositories.ListRepository {
	return &listRepository{
		client: client,
		dbName: dbName,
	}
}

func (r *listRepository) getCollection() *mongo.Collection {
	return r.client.Database(r.dbName).Collection("lists")
}

func (r *listRepository) Create(list *models.List) (*models.List, error) {
	list.CreatedAt = time.Now()
	list.UpdatedAt = time.Now()
	newList, err := r.getCollection().InsertOne(context.Background(), list)
	if err != nil {
		return nil, err
	}
	return &models.List{
		ID:        newList.InsertedID.(primitive.ObjectID).Hex(),
		Name:      list.Name,
		UserId:    list.UserId,
		CreatedAt: list.CreatedAt,
		UpdatedAt: list.UpdatedAt,
	}, nil
}

func (r *listRepository) GetByID(id string) (*models.List, error) {
	var list models.List
	docID, _ := primitive.ObjectIDFromHex(id)
	err := r.getCollection().FindOne(context.Background(), bson.M{"_id": docID}).Decode(&list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *listRepository) GetByName(name string) (*models.List, error) {
	var list models.List
	err := r.getCollection().FindOne(context.Background(), bson.M{"name": name}).Decode(&list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *listRepository) Update(id string, list *models.List) (*models.List, error) {
	docID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": list}
	_, err := r.getCollection().UpdateOne(context.Background(), bson.M{"_id": docID}, update)
	if err != nil {
		return nil, err
	}
	return &models.List{
		ID:   id,
		Name: list.Name,
	}, nil
}

func (r *listRepository) Delete(id string) error {
	docID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.getCollection().DeleteOne(context.Background(), bson.M{"_id": docID})
	return err
}

func (r *listRepository) GetAll() ([]*models.List, error) {
	opts := options.Find()
	cursor, err := r.getCollection().Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var lists []*models.List
	for cursor.Next(context.Background()) {
		var list models.List
		if err := cursor.Decode(&list); err != nil {
			return nil, err
		}
		lists = append(lists, &list)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *listRepository) GetAllByUserId(userId int64) ([]*models.List, error) {
	opts := options.Find()
	cursor, err := r.getCollection().Find(context.Background(), bson.M{"userId": userId}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var lists []*models.List
	for cursor.Next(context.Background()) {
		var list models.List
		if err := cursor.Decode(&list); err != nil {
			return nil, err
		}
		lists = append(lists, &list)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return lists, nil
}
