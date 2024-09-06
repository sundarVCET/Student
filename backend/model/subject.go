package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SubName    string             `json:"subName,omitempty" validate:"required"  bson:"subName,omitempty"`
	SubCode    string             `json:"subCode,omitempty" validate:"required"  bson:"subCode,omitempty"`
	Sessions   string             `json:"sessions,omitempty" validate:"required"  bson:"sessions,omitempty"`
	SclassName primitive.ObjectID `json:"sclassName,omitempty" validate:"required"  bson:"sclassName,omitempty"`
	School     primitive.ObjectID `json:"school,omitempty"   bson:"school,omitempty"`
	Teacher    string             `json:"teacher,omitempty"  bson:"teacher,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
type SubjectRequest struct {
	AdminID    string     `json:"adminId,omitempty" validate:"required" bson:"adminId,omitempty"`
	SclassName string     `json:"sclassName,omitempty" validate:"required"  bson:"sclassName,omitempty"`
	Subjects   []Subjects `json:"subjects" validate:"required" bson:"subjects"`
}
type Subjects struct {
	SubName  string `json:"subName,omitempty" validate:"required"  bson:"subName,omitempty"`
	SubCode  string `json:"subCode,omitempty" validate:"required"  bson:"subCode,omitempty"`
	Sessions string `json:"sessions,omitempty" validate:"required"  bson:"sessions,omitempty"`
}
type SubjectResponse struct {
	Subject
	SclassName     *SclassName     `json:"sclassName,omitempty"  bson:"sclassName,omitempty"`
	TeacherDetails *TeacherDetails `json:"teacher,omitempty"  bson:"teacher,omitempty"`
}
type TeacherDetails struct {
	Id   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty"  bson:"name,omitempty"`
}
