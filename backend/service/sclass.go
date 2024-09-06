package service

import (
	"context"
	"errors"
	"fmt"
	db "student/database"
	"student/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type SClassRepository struct{}

func (classRepo *SClassRepository) SclassCreate(request *model.SclassRequest) (*model.SclassRes, error) {

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(request.AdminId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	var sclass model.Sclass
	filter := bson.M{"school": id, "sclassName": request.SclassName}
	err = db.Sclasses.FindOne(context.TODO(), filter).Decode(&sclass)
	if err == nil {
		// If no error, it means the class already exists
		return nil, errors.New("sorry this class name already exists")
	} else if err != mongo.ErrNoDocuments {
		// Error other than not finding a document, return it
		return nil, fmt.Errorf("error checking for existing class: %v", err)
	}

	// Populate sclass fields
	sclass.SclassName = request.SclassName
	sclass.CreatedAt = time.Now()
	sclass.UpdatedAt = time.Now()
	sclass.School = id
	// Insert sclass into the database
	result, err := db.Sclasses.InsertOne(context.TODO(), sclass)
	if err != nil {
		return nil, fmt.Errorf("error inserting class into database: %v", err)
	}
	var sclassRes model.SclassRes

	sclassRes.SclassName = request.SclassName
	sclassRes.CreatedAt = time.Now()
	sclassRes.UpdatedAt = time.Now()
	sclassRes.School = id
	sclassRes.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return &sclassRes, nil

}

func (classRepo *SClassRepository) SclassList(Id string) ([]*model.SclassRes, error) {

	var sclasses []*model.SclassRes

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	filter := bson.M{"school": id}
	cur, _ := db.Sclasses.Find(context.TODO(), filter)

	defer cur.Close(context.TODO())

	// Iterate over the cursor to decode documents
	for cur.Next(context.Background()) {
		// Create a new Sclass instance to decode the document into
		var sclass model.SclassRes
		err := cur.Decode(&sclass)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		// sclass.School =
		// Append the decoded document to the sclasses slice
		sclasses = append(sclasses, &sclass)
	}
	fmt.Println("Sclasses", sclasses)

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over cursor: %v", err)
	}

	// If no documents were found, return an error
	if len(sclasses) == 0 {
		return nil, errors.New("no sclasses found")
	}

	fmt.Printf("Found multiple documents: %+v\n", sclasses)

	return sclasses, nil
}
func (classRepo *SClassRepository) GetSclassDetail(Id string) (*model.SclassResult, error) {

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	var existingSClass *model.SclassRes

	filter := bson.M{"_id": id}
	_ = db.Sclasses.FindOne(context.TODO(), filter).Decode(&existingSClass)

	if existingSClass == nil {
		// If no error, it means the class already exists
		return nil, errors.New("no class found")
	}

	filter = bson.M{"_id": existingSClass.School}
	var school model.School
	_ = db.Admins.FindOne(context.TODO(), filter).Decode(&school)

	schoolResult := &model.SclassResult{}

	schoolResult.School = school
	schoolResult.CreatedAt = existingSClass.CreatedAt
	schoolResult.UpdatedAt = existingSClass.UpdatedAt
	schoolResult.SclassName = existingSClass.SclassName
	schoolResult.Id = existingSClass.Id

	return schoolResult, nil
}
func (classRepo *SClassRepository) GetSclassStudents(Id string) ([]*model.Student, error) {

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	var students []*model.Student

	filter := bson.M{"sclassName": id}
	cur, _ := db.Students.Find(context.TODO(), filter)

	defer cur.Close(context.TODO())

	// Iterate over the cursor to decode documents
	for cur.Next(context.Background()) {
		// Create a new Sclass instance to decode the document into
		var student model.Student
		err := cur.Decode(&student)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		student.BytePassword = nil
		// Append the decoded document to the sclasses slice
		students = append(students, &student)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over cursor: %v", err)
	}

	// If no documents were found, return an error
	if len(students) == 0 {
		return nil, errors.New("no sclasses found")
	}

	fmt.Printf("Found multiple documents: %+v\n", students)

	return students, nil

}
func (classRepo *SClassRepository) DeleteSclasses(Id string) (*model.DeleteResponse, error) {
	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	filter := bson.M{"school": id}
	result, _ := db.Sclasses.DeleteMany(context.TODO(), filter)
	if result.DeletedCount == 0 {
		return nil, errors.New("no classes found to delete")
	}
	filter = bson.M{"school": id}

	db.Students.DeleteMany(context.TODO(), filter)
	db.Subjects.DeleteMany(context.TODO(), filter)
	db.Teachers.DeleteMany(context.TODO(), filter)

	deleteResp := &model.DeleteResponse{
		Acknowledged: result.DeletedCount > 0,
		DeletedCount: result.DeletedCount,
	}

	return deleteResp, nil
}

func (classRepo *SClassRepository) DeleteSclass(Id string) (*model.Sclass, error) {

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	var deleteSclass *model.Sclass
	filter := bson.M{"_id": id}
	_ = db.Sclasses.FindOneAndDelete(context.TODO(), filter).Decode(&deleteSclass)
	if deleteSclass != nil {
		return nil, errors.New("no classes found to delete")
	}
	filter = bson.M{"sclassName": id}

	db.Students.DeleteMany(context.TODO(), filter)
	db.Subjects.DeleteMany(context.TODO(), filter)
	db.Teachers.DeleteMany(context.TODO(), filter)

	return deleteSclass, nil
}
