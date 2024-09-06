package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"student/auth"
	db "student/database"
	"student/model"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type AdminRepository struct{}

func (adminRepo *AdminRepository) AdminRegister(request *model.Admin) (*model.Admin, error) {

	var existingEmailAdmin model.Admin
	var existingSchoolAdmin model.Admin

	filter := bson.M{"email": request.Email}
	_ = db.Admins.FindOne(context.TODO(), filter).Decode(&existingEmailAdmin)

	filter = bson.M{"schoolName": request.SchoolName}
	_ = db.Admins.FindOne(context.TODO(), filter).Decode(&existingSchoolAdmin)

	if existingEmailAdmin.Email != "" {
		return nil, errors.New("email already exists")
	}
	if existingSchoolAdmin.SchoolName != "" {
		return nil, errors.New("schoolName already exists")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 6)
	if err != nil {
		log.Println("bcrypt password error:", err)
	}
	request.BytePassword = pass
	request.Password = ""

	result, err := db.Admins.InsertOne(context.TODO(), request)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	// Set the ID in the admin object to include it in the response
	request.ID = result.InsertedID.(primitive.ObjectID).Hex()
	request.BytePassword = nil

	return request, nil
}
func (adminRepo *AdminRepository) AdminLogIn(request *model.Admin) (*model.Admin, error) {

	var admin *model.Admin

	if request.Email != "" || request.Password != "" {

		filter := bson.M{"email": request.Email}
		err := db.Admins.FindOne(context.TODO(), filter).Decode(&admin)
		if err != nil {
			return nil, errors.New("user not found")
		}
		if admin.Email != "" {

			err := bcrypt.CompareHashAndPassword(admin.BytePassword, []byte(request.Password))
			if err != nil {
				return nil, errors.New("invalid password")
			}
			admin.BytePassword = nil
		}

		token, err := auth.GenerateJWT(admin.Email, admin.Name, "ADMIN")
		if err != nil {
			return nil, errors.New("failed to generate token")
		}
		admin.Token = token

		return admin, nil

	} else {

		return nil, errors.New("email and password are required")
	}
}
func (adminRepo *AdminRepository) GetAdminDetail(Id string) (*model.Admin, error) {

	var admin *model.Admin

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId: %v", err)
	}

	filter := bson.M{"_id": id}
	err = db.Admins.FindOne(context.TODO(), filter).Decode(&admin)

	if err != nil {
		return nil, errors.New("no admin found")
	}
	admin.BytePassword = nil
	admin.Password = ""

	return admin, nil
}
