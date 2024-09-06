package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	db "student/database"
	"student/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type TeacherRepository struct{}

func (teacherRepo *TeacherRepository) TeacherRegister(request *model.Teacher) (*model.Teacher, error) {

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 6)
	if err != nil {
		log.Println("bcrypt password error:", err)
	}
	request.BytePassword = pass
	request.Password = ""

	var existingTeacher *model.Teacher

	filter := bson.M{"email": request.Email}
	_ = db.Teachers.FindOne(context.TODO(), filter).Decode(&existingTeacher)
	// if err != nil {
	// 	return nil, fmt.Errorf("error finding Teacher: %v", err)
	// }

	if existingTeacher != nil {
		return nil, errors.New("email already exists'")
	} else {

		request.CreatedAt = time.Now()
		request.UpdatedAt = time.Now()
		result, err := db.Teachers.InsertOne(context.TODO(), request)
		if err != nil {
			return nil, err
		}
		// Initialize the attendance field as an empty array
		request.Attendance = []model.Attendances{}
		request.Id = result.InsertedID.(primitive.ObjectID).Hex()
		request.BytePassword = nil
		return request, nil

	}
}

func (teacherRepo *TeacherRepository) TeacherLogIn(request *model.Teacher) (*model.TeacherResponse, error) {

	var existingTeacher *model.Teacher

	filter := bson.M{"email": request.Email}
	_ = db.Teachers.FindOne(context.TODO(), filter).Decode(&existingTeacher)

	if existingTeacher != nil {

		err := bcrypt.CompareHashAndPassword(existingTeacher.BytePassword, []byte(request.Password))
		if err != nil {
			return nil, errors.New("invalid password")
		}
		existingTeacher.BytePassword = nil

		var school model.School
		filter = bson.M{"_id": existingTeacher.School}
		err = db.Admins.FindOne(context.TODO(), filter).Decode(&school)
		if err != nil {
			return nil, fmt.Errorf("error finding school: %v", err)
		}

		var sclassName model.SclassName
		filter = bson.M{"_id": existingTeacher.TeachSclass}
		err = db.Sclasses.FindOne(context.TODO(), filter).Decode(&sclassName)
		if err != nil {
			return nil, fmt.Errorf("error finding sclassName: %v", err)
		}

		var teachSubject model.TeachSubject
		filter = bson.M{"_id": existingTeacher.TeachSubject}
		err = db.Subjects.FindOne(context.TODO(), filter).Decode(&teachSubject)
		if err != nil {
			return nil, fmt.Errorf("error finding teachSubject: %v", err)
		}

		teacherResp := &model.TeacherResponse{
			Id:           existingTeacher.Id,
			Name:         existingTeacher.Name,
			Email:        existingTeacher.Email,
			Role:         existingTeacher.Role,
			UpdatedAt:    existingTeacher.UpdatedAt,
			CreatedAt:    existingTeacher.CreatedAt,
			Attendance:   existingTeacher.Attendance,
			TeachSclass:  sclassName,
			TeachSubject: teachSubject,
			School:       school,
		}

		return teacherResp, nil

	} else {
		return nil, errors.New("teacher not found")
	}

}

func (teacherRepo *TeacherRepository) GetTeachers(Id string) ([]*model.Teachers, error) {

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	var teachers []*model.Teachers

	filter := bson.M{"school": id}
	cur, _ := db.Teachers.Find(context.TODO(), filter)

	defer cur.Close(context.TODO())

	// Iterate over the cursor to decode documents
	for cur.Next(context.Background()) {

		// Create a new teacher instance to decode the document into
		var teacher model.Teacher
		err := cur.Decode(&teacher)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}

		var sclassName model.SclassName
		filter := bson.M{"_id": teacher.TeachSclass}
		err = db.Sclasses.FindOne(context.TODO(), filter).Decode(&sclassName)
		if err != nil {
			return nil, fmt.Errorf("error finding sclassName: %v", err)
		}

		var teachSubject model.TeachSubject
		filter = bson.M{"_id": teacher.TeachSubject}
		err = db.Subjects.FindOne(context.TODO(), filter).Decode(&teachSubject)
		if err != nil {
			return nil, fmt.Errorf("error finding teachSubject: %v", err)
		}
		//Append the decoded document to the sclasses slice
		teacherResp := &model.Teachers{
			Id:           teacher.Id,
			Name:         teacher.Name,
			Email:        teacher.Email,
			Role:         teacher.Role,
			UpdatedAt:    teacher.UpdatedAt,
			CreatedAt:    teacher.CreatedAt,
			Attendance:   teacher.Attendance,
			TeachSclass:  sclassName,
			TeachSubject: teachSubject,
			School:       teacher.School,
		}
		teachers = append(teachers, teacherResp)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over cursor: %v", err)
	}

	// If no documents were found, return an error
	if len(teachers) == 0 {
		return nil, errors.New("no teachers found")
	}

	return teachers, nil

}

func (teacherRepo *TeacherRepository) GetTeacherDetail(Id string) (*model.TeacherResponse, error) {

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	// Create a new teacher instance to decode the document into
	var teacher *model.Teacher
	filter := bson.M{"_id": id}
	err = db.Teachers.FindOne(context.TODO(), filter).Decode(&teacher)
	if err != nil {
		return nil, fmt.Errorf("error decoding document: %v", err)
	}

	if teacher != nil {

		var sclassName model.SclassName
		filter := bson.M{"_id": teacher.TeachSclass}
		err = db.Sclasses.FindOne(context.TODO(), filter).Decode(&sclassName)
		if err != nil {
			return nil, fmt.Errorf("error finding sclassName: %v", err)
		}

		var teachSubject model.TeachSubject
		filter = bson.M{"_id": teacher.TeachSubject}
		err = db.Subjects.FindOne(context.TODO(), filter).Decode(&teachSubject)
		if err != nil {
			return nil, fmt.Errorf("error finding teachSubject: %v", err)
		}

		var school model.School
		filter = bson.M{"_id": teacher.School}
		err = db.Admins.FindOne(context.TODO(), filter).Decode(&school)
		if err != nil {
			return nil, fmt.Errorf("error finding school: %v", err)
		}

		//Append the decoded document to the sclasses slice
		teacherResp := &model.TeacherResponse{
			Id:           teacher.Id,
			Name:         teacher.Name,
			Email:        teacher.Email,
			Role:         teacher.Role,
			UpdatedAt:    teacher.UpdatedAt,
			CreatedAt:    teacher.CreatedAt,
			Attendance:   teacher.Attendance,
			TeachSclass:  sclassName,
			TeachSubject: teachSubject,
			School:       school,
		}
		return teacherResp, nil
	} else {
		return nil, errors.New("no teacher found")
	}

}

func (teacherRepo *TeacherRepository) DeleteTeachers(request *model.Teacher) (*model.Teacher, error) {

	return nil, nil
}

func (teacherRepo *TeacherRepository) DeleteTeachersByClass(request *model.Teacher) (*model.Teacher, error) {

	return nil, nil
}
func (teacherRepo *TeacherRepository) DeleteTeacher(request *model.Teacher) (*model.Teacher, error) {

	return nil, nil
}

func (teacherRepo *TeacherRepository) UpdateTeacherSubject(request *model.Teacher) (*model.Teacher, error) {

	return nil, nil
}

func (teacherRepo *TeacherRepository) TeacherAttendance(request *model.Teacher) (*model.Teacher, error) {

	return nil, nil
}
