package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Item struct {
	ID        string             `json:"id" bson:"_id,omitempty"`
	ListId    primitive.ObjectID `json:"listId" bson:"listId"`
	ItemCode  int64              `json:"itemCode" bson:"itemCode"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
}
