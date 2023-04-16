package appModel

type Team struct {
	Id    int64  `json:"id,omitempty" bson:"id"`
	Name  string `json:"name,omitempty" bson:"name"`
	Crest string `json:"crest,omitempty" bson:"crest"`
}
