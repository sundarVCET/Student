package handler

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
	db "student/database"
	"student/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type NotificationService struct{}

func AddImage(c *gin.Context) {
	var req model.Image
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !strings.HasSuffix(req.ImageName, ".png") && !strings.HasSuffix(req.ImageName, ".jpeg") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
		return
	}

	if req.Data != "" {
		s := strings.Split(req.Data, ",")
		if len(s) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": ""})
			return
		}

		byteImage := make([]byte, base64.StdEncoding.DecodedLen(len(s[1])))
		_, err := base64.StdEncoding.Decode(byteImage, []byte(s[1]))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ""})
			return
		}
		// string to primitive.ObjectID
		// Id, err := primitive.ObjectIDFromHex(req)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid objectid"})
		// 	return
		//}
		// image := model.Image{
		// 	Content:     byteImage,
		// 	Data:        req.Data,
		// 	ImageName:   strings.TrimSuffix(req.ImageName, ".png"),
		// 	CreatedDate: time.Now(),
		// 	Description: req.Description,
		// 	UserId:      req.UserId,
		// }
		filter := bson.M{"userId": req.UserId}
		update := bson.M{
			"$set": bson.M{
				"imageName":   req.ImageName,
				"data":        req.Data,
				"createdDate": req.CreatedDate,
				"description": req.Description,
				"role":        req.UserRole,
			},
		}

		options := options.Update().SetUpsert(true)

		result, err := db.Image.UpdateOne(context.TODO(), filter, update, options)
		if err != nil {
			log.Println("DB error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		var responseID string
		if result.UpsertedID != nil {
			// A new document was inserted
			responseID = req.UserId
		} else {
			// No new document was inserted; use the filter criteria to get the ID of the updated document
			var updatedImage model.Image
			err := db.Image.FindOne(context.TODO(), filter).Decode(&updatedImage)
			if err != nil {
				log.Println("Error retrieving updated document:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving updated document"})
				return
			}
			responseID = req.UserId
		}

		c.JSON(http.StatusOK, gin.H{"userId": responseID})
	}

	//c.JSON(http.StatusOK, image)
}
func GetImage(c *gin.Context) {

	// Get userId from URL parameters
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID not found"})
		return
	}

	// Define a variable to hold the image
	var image model.Image

	// Filter to find the image by userId
	filter := bson.M{"userId": userId}
	err := db.Image.FindOne(context.TODO(), filter).Decode(&image)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": "Image not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving image"})
		}
		return
	}

	// Respond with the image data
	c.JSON(http.StatusOK, image)
}
