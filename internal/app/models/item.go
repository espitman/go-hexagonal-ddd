package appModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type NewItem struct {
	ListId   string `json:"listId" bson:"listId,omitempty"`
	ItemCode int64  `json:"itemCode" bson:"itemCode"`
}

type Item struct {
	ID        string             `json:"id" bson:"_id"`
	ListId    primitive.ObjectID `json:"listId" bson:"listId"`
	ItemCode  int64              `json:"itemCode" bson:"itemCode"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
}
