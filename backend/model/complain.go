package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Complain struct {
	Id        string             `json:"_id,omitempty" bson:"_id,omitempty"`
	User      primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Date      string             `json:"date,omitempty" bson:"date,omitempty"`
	Complaint string             `json:"complaint,omitempty" bson:"complaint,omitempty"`
	School    primitive.ObjectID `json:"school,omitempty" bson:"school,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
type ComplainList struct {
	Id        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	School    string    `json:"school,omitempty" bson:"school,omitempty"`
	Date      string    `json:"date,omitempty" bson:"date,omitempty"`
	Complaint string    `json:"complaint,omitempty" bson:"complaint,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	User      User      `json:"user,omitempty" bson:"user,omitempty"`
}
type User struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}
