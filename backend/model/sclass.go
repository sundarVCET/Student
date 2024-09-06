package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SclassRequest struct {
	SclassName string `json:"sclassName" validate:"required"  bson:"sclassName"`
	AdminId    string `json:"adminID" bson:"school"`
}

type Sclass struct {
	SclassName string             `json:"sclassName,omitempty" validate:"required"  bson:"sclassName"`
	School     primitive.ObjectID `json:"school,omitempty" bson:"school,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type SclassRes struct {
	Id         string             `json:"_id" bson:"_id"`
	SclassName string             `json:"sclassName,omitempty" validate:"required"  bson:"sclassName"`
	School     primitive.ObjectID `json:"school,omitempty" bson:"school,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
type School struct {
	Id         string `json:"_id" bson:"_id"`
	SchoolName string `json:"schoolName" bson:"schoolName"`
}

type SclassResult struct {
	Id         string    `json:"_id,omitempty" bson:"_id"`
	SclassName string    `json:"sclassName,omitempty" validate:"required"  bson:"sclassName,omitempty"`
	School     School    `json:"school,omitempty" bson:"school,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
type SclassName struct {
	Id         string `json:"_id" bson:"_id"`
	SclassName string `json:"sclassName" bson:"sclassName"`
}
type DeleteResponse struct {
	Acknowledged bool  `json:"acknowledged"`
	DeletedCount int64 `json:"deletedCount"`
}
