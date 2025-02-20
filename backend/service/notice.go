package service

import (
	"context"
	"errors"
	"fmt"
	db "student-api/database"
	"student-api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type NoticeRepository struct{}

func (noticeRepo *NoticeRepository) NoticeCreate(request *model.Notice) (*model.Notice, error) {

	// string to primitive.ObjectID
	adminId, err := primitive.ObjectIDFromHex(request.AdminID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	request.AdminID = "" // No need to insert in db , just create for request body this fiedl
	request.School = adminId
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	result, err := db.Notices.InsertOne(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	request.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return request, nil
}

func (noticeRepo *NoticeRepository) NoticeList(id string) ([]*model.Notice, error) {

	// string to primitive.ObjectID
	adminId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	filter := bson.M{"school": adminId}
	cur, _ := db.Notices.Find(context.TODO(), filter)
	defer cur.Close(context.TODO())
	var notices []*model.Notice
	for cur.Next(context.Background()) {
		var notice model.Notice
		err := cur.Decode(&notice)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		notices = append(notices, &notice)
	}
	if len(notices) == 0 {
		return nil, fmt.Errorf("no notices found")
	}

	return notices, nil
}
func (noticeRepo *NoticeRepository) DeleteNotices(id string) (*model.DeleteResponse, error) {
	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	filter := bson.M{"school": Id}
	result, _ := db.Notices.DeleteMany(context.TODO(), filter)
	if result.DeletedCount == 0 {
		return nil, errors.New("no notices found to delete")
	}

	deleteResp := &model.DeleteResponse{
		Acknowledged: result.DeletedCount > 0,
		DeletedCount: result.DeletedCount,
	}

	return deleteResp, nil
}

func (noticeRepo *NoticeRepository) DeleteNotice(id string) (*model.Notice, error) {
	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	var notice *model.Notice

	filter := bson.M{"_id": Id}
	fmt.Println("FILTER", filter)

	_ = db.Notices.FindOneAndDelete(context.TODO(), filter).Decode(&notice)

	if notice == nil {
		return nil, errors.New("no notices found")
	}
	return notice, nil
}

func (noticeRepo *NoticeRepository) UpdateNotice(request *model.Notice, id string) (*model.Notice, error) {

	// string to primitive.ObjectID
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	adminId, err := primitive.ObjectIDFromHex(request.AdminID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	request.School = adminId
	request.AdminID = "" // no need this field ,just handle for request body
	filter := bson.M{"_id": Id}
	update := bson.M{"$set": request}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var notice *model.Notice
	err = db.Notices.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&notice)
	if err != nil {
		return nil, fmt.Errorf("error in update %v", err)
	}
	if notice == nil {
		return nil, errors.New("no notices found")
	}
	return notice, nil

}
