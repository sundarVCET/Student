package model

import (
	"time"
)

type Image struct {
	ID          string    `json:"ID,omitempty" bson:"_id,omitempty"`
	Content     []byte    `json:"content,omitempty" bson:"content,omitempty"`
	ImageName   string    `json:"imageName,omitempty" bson:"imageName,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty" bson:"createdDate,omitempty"`
	Data        string    `json:"data,omitempty" bson:"data,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	UserId      string    `json:"userId,omitempty" bson:"userId,omitempty"`
	UserRole    string    `json:"userRole,omitempty" bson:"userRole,omitempty"`
}
