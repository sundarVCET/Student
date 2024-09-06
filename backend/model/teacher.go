package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Teacher struct {
	Id           string             `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required"  bson:"name,omitempty"`
	Email        string             `json:"email,omitempty" validate:"required"  bson:"email,omitempty"`
	BytePassword []byte             `json:"bytepassword,omitempty"  bson:"bytepassword,omitempty"`
	Password     string             `json:"password,omitempty" validate:"required"  bson:"password,omitempty"`
	Role         string             `json:"role,omitempty" validate:"required"  bson:"role,omitempty"`
	School       primitive.ObjectID `json:"school,omitempty" validate:"required"  bson:"school,omitempty"`
	TeachSubject primitive.ObjectID `json:"teachSubject,omitempty" validate:"required"  bson:"teachSubject,omitempty"`
	TeachSclass  primitive.ObjectID `json:"teachSclass,omitempty" validate:"required"  bson:"teachSclass,omitempty"`
	Attendance   []Attendances      `json:"attendance" validate:"required"  bson:"attendance"`
	CreatedAt    time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Attendances struct {
	Date         time.Time `json:"date,omitempty" validate:"required"  bson:"date,omitempty"`
	PresentCount string    `json:"presentCount,omitempty"   bson:"presentCount,omitempty"`
	PbsentCount  string    `json:"absentCount,omitempty"   bson:"absentCount,omitempty"`
}

type TeacherResponse struct {
	Id           string        `json:"_id" bson:"_id"`
	Name         string        `json:"name"   bson:"name"`
	Email        string        `json:"email"   bson:"email"`
	Role         string        `json:"role"  bson:"role"`
	Attendance   []Attendances `json:"attendance"  bson:"attendance"`
	TeachSclass  SclassName    `json:"teachSclass"  bson:"teachSclass"`
	TeachSubject TeachSubject  `json:"teachSubject"  bson:"teachSubject"`
	School       School        `json:"school,omitempty"  bson:"school"`
	CreatedAt    time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt" bson:"updatedAt"`
}
type TeachSubject struct {
	Id       string `json:"_id" bson:"_id"`
	SubName  string `json:"subName" bson:"subName"`
	Sessions string `json:"sessions" bson:"sessions"`
}
type Teachers struct {
	Id           string             `json:"_id" bson:"_id"`
	Name         string             `json:"name"   bson:"name"`
	Email        string             `json:"email"   bson:"email"`
	Role         string             `json:"role"  bson:"role"`
	Attendance   []Attendances      `json:"attendance"  bson:"attendance"`
	TeachSclass  SclassName         `json:"teachSclass"  bson:"teachSclass"`
	TeachSubject TeachSubject       `json:"teachSubject"  bson:"teachSubject"`
	School       primitive.ObjectID `json:"school"  bson:"school"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}
