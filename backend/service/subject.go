package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	db "student-api/database"
	"student-api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type SubjectRepository struct{}

func (subjectRepo *SubjectRepository) SubjectCreate(request *model.SubjectRequest) ([]interface{}, error) {

	// string to primitive.ObjectID
	AdminID, err := primitive.ObjectIDFromHex(request.AdminID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	SclassName, err := primitive.ObjectIDFromHex(request.SclassName)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	//Duplicatemap := make(map[string]string)

	for _, v := range request.Subjects {

		var existingSubject *model.Subject
		filter := bson.M{"subCode": v.SubCode, "school": AdminID}
		_ = db.Subjects.FindOne(context.TODO(), filter).Decode(&existingSubject)

		if existingSubject != nil && existingSubject.SubCode != "" {
			return nil, errors.New("sorry this subcode must be unique as it already exists")
		}

	}
	var Subjects []interface{}

	for _, v := range request.Subjects {

		var subject model.Subject
		subject.Id = primitive.NewObjectID()
		subject.SubName = v.SubName
		subject.SubCode = v.SubCode
		subject.Sessions = v.Sessions
		subject.SclassName = SclassName
		subject.School = AdminID
		subject.CreatedAt = time.Now()
		subject.UpdatedAt = time.Now()
		Subjects = append(Subjects, &subject)
	}
	result, err := db.Subjects.InsertMany(context.TODO(), Subjects)
	if err != nil {
		log.Println("Error in Inserting", err)
	}
	fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))

	return Subjects, nil
}

func (subjectRepo *SubjectRepository) AllSubjects(Id string) ([]*model.SubjectResponse, error) {

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	filter := bson.M{"school": id}

	var SubjectsResp []*model.SubjectResponse
	cur, _ := db.Subjects.Find(context.TODO(), filter)

	defer cur.Close(context.TODO())

	// Iterate over the cursor to decode documents
	for cur.Next(context.Background()) {
		// Create a new Sclass instance to decode the document into
		subjectResp := &model.SubjectResponse{}
		var subject model.Subject
		err := cur.Decode(&subject)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}

		subjectResp.Subject = subject

		classFilter := bson.M{"_id": subject.SclassName}
		fmt.Println("subject.SclassName", subject.SclassName)
		var sclassName model.SclassName
		if err := db.Sclasses.FindOne(context.TODO(), classFilter).Decode(&sclassName); err != nil {
			// You can log the error if you want to track missing sclassName documents
			sclassName = model.SclassName{} // Empty struct if not found
		}

		subjectResp.SclassName = &sclassName
		// Append the decoded document to the sclasses slice
		SubjectsResp = append(SubjectsResp, subjectResp)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over cursor: %v", err)
	}

	// If no documents were found, return an error
	if len(SubjectsResp) == 0 {
		return nil, errors.New("no sclasses found")
	}

	return SubjectsResp, nil
}

func (subjectRepo *SubjectRepository) ClassSubjects(Id string) ([]*model.Subject, error) {

	// string to primitive.ObjectID
	AdminID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	filter := bson.M{"sclassName": AdminID}

	cur, _ := db.Subjects.Find(context.TODO(), filter)

	defer cur.Close(context.TODO())

	var Subjects []*model.Subject
	// Iterate over the cursor to decode documents
	for cur.Next(context.Background()) {
		// Create a new Sclass instance to decode the document into
		var subject model.Subject
		err := cur.Decode(&subject)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		// Append the decoded document to the sclasses slice
		Subjects = append(Subjects, &subject)

	}
	// If no documents were found, return an error
	if len(Subjects) == 0 {
		return nil, errors.New("no subjects found")
	}
	return Subjects, nil
}

func (subjectRepo *SubjectRepository) FreeSubjectList(Id string) ([]*model.Subject, error) {

	// string to primitive.ObjectID
	SubjectId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	filter := bson.M{"sclassName": SubjectId, "teacher": bson.M{"$exists": false}}
	cur, _ := db.Subjects.Find(context.TODO(), filter)

	defer cur.Close(context.TODO())

	var Subjects []*model.Subject
	// Iterate over the cursor to decode documents
	for cur.Next(context.Background()) {
		// Create a new Sclass instance to decode the document into
		var subject model.Subject
		err := cur.Decode(&subject)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		// Append the decoded document to the sclasses slice
		Subjects = append(Subjects, &subject)

	}
	// If no documents were found, return an error
	if len(Subjects) == 0 {
		return nil, errors.New("no subjects found")
	}
	return nil, nil
}

func (subjectRepo *SubjectRepository) GetSubjectDetail(Id string) (*model.SubjectResponse, error) {
	// Convert string to primitive.ObjectID
	SubjectId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	// Initialize the response struct
	subjectRes := &model.SubjectResponse{}

	// Find the subject by its ID
	var subject model.Subject
	filter := bson.M{"_id": SubjectId}
	err = db.Subjects.FindOne(context.TODO(), filter).Decode(&subject)
	if err != nil {
		return nil, fmt.Errorf("error finding subject: %v", err)
	}
	subjectRes.Subject = subject

	// Try to find the corresponding class name for the subject
	var sclassName model.SclassName
	classFilter := bson.M{"_id": subject.SclassName}
	if err = db.Sclasses.FindOne(context.TODO(), classFilter).Decode(&sclassName); err == nil {
		subjectRes.SclassName = &sclassName
	} else if err != mongo.ErrNoDocuments {
		// Handle other potential errors
		return nil, fmt.Errorf("error finding sclassName: %v", err)
	}

	// Try to find the teacher details for the subject
	var teacher model.TeacherDetails
	teacherFilter := bson.M{"teachSubject": SubjectId}
	if err = db.Teachers.FindOne(context.TODO(), teacherFilter).Decode(&teacher); err == nil {
		subjectRes.TeacherDetails = &teacher
	} else if err != mongo.ErrNoDocuments {
		// Handle other potential errors
		return nil, fmt.Errorf("error finding teacher details: %v", err)
	}

	return subjectRes, nil

	// // string to primitive.ObjectID
	// SubjectId, err := primitive.ObjectIDFromHex(Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid ObjectId: %v", err)
	// }

	// subjectRes := &model.SubjectResponse{}

	// var subject model.Subject
	// filter := bson.M{"_id": SubjectId}
	// err = db.Subjects.FindOne(context.TODO(), filter).Decode(&subject)
	// if err != nil {
	// 	return nil, fmt.Errorf("error finding subject: %v", err)
	// }
	// subjectRes.Subject = subject

	// var sclassName *model.SclassName
	// filter = bson.M{"_id": subject.SclassName}
	// _ = db.Sclasses.FindOne(context.TODO(), filter).Decode(&sclassName)
	// // if err != nil {
	// // 	return nil, fmt.Errorf("error finding sclassName: %v", err)
	// // }
	// if sclassName != nil {
	// 	subjectRes.SclassName = sclassName
	// }

	// var teacher model.TeacherDetails
	// filter = bson.M{"teachSubject": SubjectId}
	// _ = db.Teachers.FindOne(context.TODO(), filter).Decode(&teacher)

	// if teacher!=nil{
	// 	subjectRes.TeacherDetails = teacher

	// }
	// if err != nil {
	// 	return nil, fmt.Errorf("error finding teachSubject: %v", err)
	// }
	// if err := db.Teachers.FindOne(context.TODO(), filter).Decode(&teacher); err != nil {
	// 	// You can log the error if you want to track missing sclassName documents
	// 	teacher = model.TeacherDetails{} // Empty struct if not found
	// }

	// subjectRes := &model.SubjectResponse{
	// 	Subject:        subject,
	// 	SclassName:     sclassName,
	// 	TeacherDetails: teacher,
	// }

	//return subjectRes, nil

}
func (subjectRepo *SubjectRepository) DeleteSubject(id string) (*model.Subject, error) {

	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	var deleteSubject *model.Subject

	filter := bson.M{"_id": Id}
	_ = db.Subjects.FindOneAndDelete(context.TODO(), filter).Decode(&deleteSubject)

	if deleteSubject != nil {
		return nil, errors.New("no subject to delete")
	}

	filter = bson.M{"teachSubject": Id}
	update := bson.M{"$sunset": bson.M{"teachSubject": ""}}

	_ = db.Teachers.FindOneAndUpdate(context.TODO(), filter, update)

	filter = bson.M{}
	update = bson.M{"$pull": bson.M{"attendance": bson.M{"subName": Id}}}

	db.Students.UpdateMany(context.TODO(), filter, update)

	filter = bson.M{}
	update = bson.M{"$pull": bson.M{"examResult": bson.M{"subName": Id}}}

	db.Students.UpdateMany(context.TODO(), filter, update)

	return deleteSubject, nil
}
func (subjectRepo *SubjectRepository) DeleteSubjects(id string) (*model.DeleteResponse, error) {

	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	filter := bson.M{"school": Id}
	subjectCursor, err := db.Subjects.Find(context.Background(), filter)
	if err != nil {
		return nil, nil
	}
	defer subjectCursor.Close(context.Background())

	var deletedSubjectIDs []interface{}
	for subjectCursor.Next(context.Background()) {
		var subject bson.M
		if err := subjectCursor.Decode(&subject); err != nil {
			return nil, nil
		}
		deletedSubjectIDs = append(deletedSubjectIDs, subject["_id"])
	}

	if err := subjectCursor.Err(); err != nil {

		return nil, nil
	}

	filter = bson.M{"school": Id}
	deletedSubjects, _ := db.Subjects.DeleteMany(context.TODO(), filter)

	if deletedSubjects.DeletedCount == 0 {
		return nil, errors.New("no subject to delete")
	}

	// Set the teachSubject field to null in teachers

	update := bson.M{"$unset": bson.M{"teachSubject": ""}}
	_, _ = db.Teachers.UpdateMany(context.TODO(), deletedSubjectIDs, update)

	// Set examResult and attendance to null in all students

	update = bson.M{"$unset": bson.M{"examResult": []interface{}{}, "attendance": []interface{}{}}}
	_, _ = db.Teachers.UpdateMany(context.TODO(), deletedSubjectIDs, update)

	deleteResp := &model.DeleteResponse{
		Acknowledged: deletedSubjects.DeletedCount > 0,
		DeletedCount: deletedSubjects.DeletedCount,
	}

	return deleteResp, nil
}

func (subjectRepo *SubjectRepository) DeleteSubjectsByClass(id string) (*model.DeleteResponse, error) {

	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	filter := bson.M{"sclassName": Id}
	subjectCursor, err := db.Subjects.Find(context.Background(), filter)
	if err != nil {
		return nil, nil
	}
	defer subjectCursor.Close(context.Background())

	var deletedSubjectIDs []interface{}
	for subjectCursor.Next(context.Background()) {
		var subject bson.M
		if err := subjectCursor.Decode(&subject); err != nil {
			return nil, nil
		}
		deletedSubjectIDs = append(deletedSubjectIDs, subject["_id"])
	}

	if err := subjectCursor.Err(); err != nil {

		return nil, nil
	}

	filter = bson.M{"school": Id}
	deletedSubjects, _ := db.Subjects.DeleteMany(context.TODO(), filter)

	if deletedSubjects.DeletedCount == 0 {
		return nil, errors.New("no subject to delete")
	}

	// Set the teachSubject field to null in teachers

	update := bson.M{"$unset": bson.M{"teachSubject": ""}}
	_, _ = db.Teachers.UpdateMany(context.TODO(), deletedSubjectIDs, update)

	// Set examResult and attendance to null in all students

	update = bson.M{"$unset": bson.M{"examResult": []interface{}{}, "attendance": []interface{}{}}}
	_, _ = db.Teachers.UpdateMany(context.TODO(), deletedSubjectIDs, update)

	deleteResp := &model.DeleteResponse{
		Acknowledged: deletedSubjects.DeletedCount > 0,
		DeletedCount: deletedSubjects.DeletedCount,
	}

	return deleteResp, nil
}
