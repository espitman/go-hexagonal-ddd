package appModel

type Team struct {
	Id     string `json:"id,omitempty" bson:"id"`
	TeamId int64  `json:"teamId,omitempty" bson:"teamId"`
	Name   string `json:"name,omitempty" bson:"name"`
	Crest  string `json:"crest,omitempty" bson:"crest"`
}
