package models

import "time"

type List struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	UserId    int64     `json:"userId" bson:"userId"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
