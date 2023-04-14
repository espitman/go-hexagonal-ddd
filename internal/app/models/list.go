package appModel

type NewList struct {
	Name   string `json:"name" bson:"name"`
	UserId int64  `json:"userId" bson:"userId"`
}

type List struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name,omitempty"`
	UserId int64  `json:"userId" bson:"userId"`
}
