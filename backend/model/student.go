package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID           string             `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required" bson:"name,omitempty"`
	AdminID      string             `json:"adminId,omitempty" validate:"required" bson:"adminId,omitempty"`
	RollNum      string             `json:"rollNum,omitempty" validate:"required" bson:"rollNum,omitempty"`
	BytePassword []byte             `json:"bytePassword,omitempty"  bson:"bytePassword,omitempty"`
	Password     string             `json:"password,omitempty" validate:"required" bson:"password,omitempty"`
	SclassName   primitive.ObjectID `json:"sclassName,omitempty" validate:"required" bson:"sclassName,omitempty"`
	School       primitive.ObjectID `json:"school,omitempty" validate:"required" bson:"school,omitempty"`
	Role         string             `json:"role,omitempty" bson:"role,omitempty"`
	ExamResult   []ExamResult       `json:"examResult" bson:"examResult"`
	Attendance   []Attendance       `json:"attendance" validate:"required" bson:"attendance"`
}
type ExamResult struct {
	SubName       string `json:"subName,omitempty" bson:"subName,omitempty"`
	MarksObtained int    `json:"marksObtained,omitempty" bson:"marksObtained,omitempty"`
}
type Attendance struct {
	Date    time.Time          `json:"date,omitempty" validate:"required" bson:"date,omitempty"`
	Status  string             `json:"status,omitempty" validate:"required" bson:"status,omitempty"`
	SubName primitive.ObjectID `json:"subName,omitempty" validate:"required" bson:"subName,omitempty"`
}
type StudentLoginResponse struct {
	ID         string       `json:"_id" bson:"_id"`
	Name       string       `json:"name" bson:"name"`
	Role       string       `json:"role"  bson:"role"`
	RollNum    string       `json:"rollNum"  bson:"rollNum"`
	School     School       `json:"school" bson:"school"`
	SclassName SclassName   `json:"sclassName" bson:"sclassName"`
	ExamResult []ExamResult `json:"examResult,omitempty" bson:"examResult,omitempty"`
	Attendance []Attendance `json:"attendance,omitempty" bson:"attendance,omitempty"`
}

type StudentLoginRequest struct {
	StudentName string `json:"studentName,omitempty" validate:"required" bson:"studentName,omitempty"`
	RollNum     string `json:"rollNum,omitempty" validate:"required" bson:"rollNum,omitempty"`
	Password    string `json:"password,omitempty" validate:"required" bson:"password,omitempty"`
}
type Students struct {
	ID         string       `json:"_id" bson:"_id"`
	Name       string       `json:"name" bson:"name"`
	Role       string       `json:"role"  bson:"role"`
	RollNum    string       `json:"rollNum"  bson:"rollNum"`
	School     string       `json:"school" bson:"school"`
	SclassName SclassName   `json:"sclassName" bson:"sclassName"`
	ExamResult []ExamResult `json:"examResult" bson:"examResult"`
	Attendance []Attendance `json:"attendance" bson:"attendance"`
}
