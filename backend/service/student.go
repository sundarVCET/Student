package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	db "student-api/database"
	"student-api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type StudentRepository struct{}

func (studentRepo *StudentRepository) StudentRegister(request *model.Student) (*model.Student, error) {

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 6)
	if err != nil {
		log.Println("bcrypt password error:", err)
	}
	request.BytePassword = pass
	request.Password = ""

	var existingStudent *model.Student

	// string to primitive.ObjectID
	adminId, err := primitive.ObjectIDFromHex(request.AdminID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	filter := bson.M{"rollNum": request.RollNum, "school": adminId, "sclassName": request.SclassName}
	_ = db.Students.FindOne(context.TODO(), filter).Decode(&existingStudent)

	if existingStudent != nil {
		return nil, errors.New("roll Number already exists")
	} else {
		request.School = adminId
		request.AdminID = ""

		result, err := db.Students.InsertOne(context.TODO(), request)
		if err != nil {
			return nil, err
		}
		request.ID = result.InsertedID.(primitive.ObjectID).Hex()
		request.BytePassword = nil
		return request, nil

	}
}

func (studentRepo *StudentRepository) StudentLogIn(request *model.StudentLoginRequest) (*model.StudentLoginResponse, error) {

	var student *model.Student

	filter := bson.M{"rollNum": request.RollNum, "name": request.StudentName}
	_ = db.Students.FindOne(context.TODO(), filter).Decode(&student)

	if student == nil {
		return nil, errors.New("student not found")
	}

	if student != nil {

		err := bcrypt.CompareHashAndPassword(student.BytePassword, []byte(request.Password))
		if err != nil {
			return nil, errors.New("invalid password")
		} else {

			var admin model.Admin

			adminId, err := primitive.ObjectIDFromHex(student.School.Hex())
			if err != nil {
				return nil, fmt.Errorf("invalid ObjectId: %v", err)
			}

			filter = bson.M{"_id": adminId}
			_ = db.Admins.FindOne(context.TODO(), filter).Decode(&admin)

			sclassName, err := primitive.ObjectIDFromHex(student.SclassName.Hex())
			if err != nil {
				return nil, fmt.Errorf("invalid ObjectId: %v", err)
			}
			var sClassName model.SclassName

			filter = bson.M{"_id": sclassName}
			_ = db.Sclasses.FindOne(context.TODO(), filter).Decode(&sClassName)

			studentLogin := &model.StudentLoginResponse{}
			studentLogin.ID = student.ID
			studentLogin.Name = student.Name
			studentLogin.RollNum = student.RollNum
			studentLogin.Role = student.Role

			studentLogin.School = model.School{}
			studentLogin.School.Id = admin.ID
			studentLogin.School.SchoolName = admin.SchoolName

			studentLogin.SclassName = model.SclassName{}

			studentLogin.SclassName.Id = sClassName.Id
			studentLogin.SclassName.SclassName = sClassName.SclassName

			return studentLogin, nil

		}

	}

	return nil, nil
}
func (studentRepo *StudentRepository) GetStudents(Id string) ([]*model.Students, error) {

	var students []*model.Students
	// string to primitive.ObjectID
	SchoolId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	filter := bson.M{"school": SchoolId}
	cur, _ := db.Students.Find(context.TODO(), filter)

	defer cur.Close(context.TODO())
	fmt.Println("Students", students)
	// Iterate over the cursor to decode documents
	for cur.Next(context.Background()) {
		// Create a new Sclass instance to decode the document into
		var student model.Student
		err := cur.Decode(&student)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}

		var sclassName model.SclassName

		filter = bson.M{"school": SchoolId}
		err = db.Sclasses.FindOne(context.TODO(), filter).Decode(&sclassName)
		if err != nil {
			return nil, fmt.Errorf("error finding sclassName: %v", err)
		}

		studentResp := &model.Students{
			ID:         student.ID,
			Name:       student.Name,
			Role:       student.Role,
			RollNum:    student.RollNum,
			Attendance: student.Attendance,
			SclassName: sclassName,
			School:     Id,
		}
		// Append the decoded document to the sclasses slice
		students = append(students, studentResp)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over cursor: %v", err)
	}

	// If no documents were found, return an error
	if len(students) == 0 {
		return nil, errors.New("no students found")
	}

	fmt.Printf("Found multiple documents: %+v\n", students)

	return students, nil

}
func (studentRepo *StudentRepository) GetStudentDetail(id string) (*model.StudentLoginResponse, error) {

	// string to primitive.ObjectID
	adminId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	var student *model.Student
	filter := bson.M{"_id": adminId}
	_ = db.Students.FindOne(context.TODO(), filter).Decode(&student)

	if student != nil {
		var sclassName model.SclassName

		filter = bson.M{"school": student.School}
		err = db.Sclasses.FindOne(context.TODO(), filter).Decode(&sclassName)
		if err != nil {
			return nil, fmt.Errorf("error finding sclassName: %v", err)
		}
		var school model.School
		filter = bson.M{"_id": student.School}
		err = db.Admins.FindOne(context.TODO(), filter).Decode(&school)
		if err != nil {
			return nil, fmt.Errorf("error finding school: %v", err)
		}

		studentResp := &model.StudentLoginResponse{
			ID:         student.ID,
			Name:       student.Name,
			Role:       student.Role,
			RollNum:    student.RollNum,
			School:     school,
			SclassName: sclassName,
			ExamResult: student.ExamResult,
			Attendance: student.Attendance,
		}
		return studentResp, nil
	} else {
		return nil, errors.New("no students found")
	}
}
func (studentRepo *StudentRepository) DeleteStudents(id string) (*model.DeleteResponse, error) {

	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	filter := bson.M{"school": Id}
	result, _ := db.Students.DeleteMany(context.TODO(), filter)
	if result.DeletedCount == 0 {
		return nil, errors.New("no students found to delete")
	}

	deleteResp := &model.DeleteResponse{
		Acknowledged: result.DeletedCount > 0,
		DeletedCount: result.DeletedCount,
	}

	return deleteResp, nil
}

func (studentRepo *StudentRepository) DeleteStudentsByClass(id string) (*model.DeleteResponse, error) {

	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	filter := bson.M{"sclassName": Id}
	result, _ := db.Students.DeleteMany(context.TODO(), filter)
	if result.DeletedCount == 0 {
		return nil, errors.New("no students found to delete")
	}

	deleteResp := &model.DeleteResponse{
		Acknowledged: result.DeletedCount > 0,
		DeletedCount: result.DeletedCount,
	}

	return deleteResp, nil

}
func (studentRepo *StudentRepository) DeleteStudent(id string) (*model.Student, error) {

	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	var deleteStudent *model.Student

	filter := bson.M{"_id": Id}
	_ = db.Students.FindOneAndDelete(context.TODO(), filter).Decode(&deleteStudent)

	if deleteStudent != nil {
		return nil, errors.New("no students found to delete")
	}
	return deleteStudent, nil
}
func (studentRepo *StudentRepository) UpdateStudent(request *model.Student) (*model.Student, error) {

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 6)
	if err != nil {
		log.Println("bcrypt password error:", err)
	}
	request.BytePassword = pass
	request.Password = ""

	var existingStudent *model.Student

	// string to primitive.ObjectID
	adminId, err := primitive.ObjectIDFromHex(request.AdminID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(request.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	request.School = adminId
	request.AdminID = ""
	filter := bson.M{"_id": Id}
	update := bson.M{"$set": request}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	_ = db.Students.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&existingStudent)

	request.BytePassword = nil
	return request, nil

}

func (studentRepo *StudentRepository) UpdateExamResult(id string) (*model.Student, error) {
	return nil, nil
}
func (studentRepo *StudentRepository) StudentAttendance(id string) (*model.Student, error) {
	return nil, nil
}

func (studentRepo *StudentRepository) ClearAllStudentsAttendanceBySubject(id string) (*model.Student, error) {
	return nil, nil
}
func (studentRepo *StudentRepository) ClearAllStudentsAttendance(id string) (*model.Student, error) {
	return nil, nil
}
func (studentRepo *StudentRepository) RemoveStudentAttendanceBySubject(id string) (*model.Student, error) {
	return nil, nil
}
func (studentRepo *StudentRepository) RemoveStudentAttendance(id string) (*model.Student, error) {
	return nil, nil
}
