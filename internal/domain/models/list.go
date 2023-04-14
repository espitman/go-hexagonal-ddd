package models

type List struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Name      string `json:"name" bson:"name"`
	UserId    int64  `json:"userId" bson:"userId"`
	CreatedAt int64  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at,omitempty"`
}
