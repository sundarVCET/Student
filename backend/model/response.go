package model

type Response struct {
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}
