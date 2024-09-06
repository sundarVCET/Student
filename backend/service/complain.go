package service

import (
	"context"
	"errors"
	"fmt"
	db "student/database"
	"student/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type ComplainRepository struct{}

func (complainRepo *ComplainRepository) ComplainCreate(request *model.Complain) (*model.Complain, error) {
	// string to primitive.ObjectID
	school, err := primitive.ObjectIDFromHex(request.School.Hex())
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	user, err := primitive.ObjectIDFromHex(request.User.Hex())
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	request.School = school
	request.User = user
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	result, err := db.Complains.InsertOne(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	request.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return request, nil
}
func (complainRepo *ComplainRepository) ComplainList(id string) ([]*model.ComplainList, error) {

	// string to primitive.ObjectID
	school, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	filter := bson.M{"school": school}

	var ComplainList []*model.ComplainList
	cur, _ := db.Complains.Find(context.TODO(), filter)
	defer cur.Close(context.TODO())

	for cur.Next(context.Background()) {
		// Create a new complain instance to decode the document into
		var complain model.Complain

		err := cur.Decode(&complain)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		var user model.User
		_ = db.Students.FindOne(context.TODO(), bson.M{"_id": complain.User}).Decode(&user)

		// Create and populate the ComplainList instance
		complainListInstance := &model.ComplainList{
			Id:        complain.Id,
			CreatedAt: complain.CreatedAt,
			UpdatedAt: complain.UpdatedAt,
			Complaint: complain.Complaint,
			Date:      complain.Date,
			School:    id,
			User:      user,
		}

		ComplainList = append(ComplainList, complainListInstance)

	}
	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over cursor: %v", err)
	}

	// If no documents were found, return an error
	if len(ComplainList) == 0 {
		return nil, errors.New("no complains found")
	}

	return ComplainList, nil
}
