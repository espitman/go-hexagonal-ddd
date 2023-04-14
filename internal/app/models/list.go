package appModel

type NewList struct {
	Name string `json:"name" bson:"name,omitempty"`
}

type List struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
}
