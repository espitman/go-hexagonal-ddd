package appModel

import "time"

type NewList struct {
	Name string `json:"name" bson:"name"`
}

type List struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	UserId    int64     `json:"userId" bson:"userId"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
