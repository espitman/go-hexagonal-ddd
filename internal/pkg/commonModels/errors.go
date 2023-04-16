package commonModels

type ErrorResponse struct {
	Error string `json:"error" bson:"error"`
}
