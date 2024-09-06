package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notice struct {
	Id        string             `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Details   string             `json:"details,omitempty" bson:"details,omitempty"`
	Date      string             `json:"date,omitempty" bson:"date,omitempty"`
	School    primitive.ObjectID `json:"school,omitempty" bson:"school,omitempty"`
	AdminID   string             `json:"adminID,omitempty" bson:"adminID,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
