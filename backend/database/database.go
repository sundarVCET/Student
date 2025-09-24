package db

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Admins, Complains, Notices, Students, Subjects, Sclasses, Teachers, Image *mongo.Collection
var Bucket *gridfs.Bucket

var (
	client *mongo.Client
	ctx    = context.TODO()
)

func Init() {
	clientOptions := options.Client().ApplyURI(viper.GetString("MongoURL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	// Check if the connection is successful
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}
	log.Println("Connected to MongoDB")

	dbName := viper.GetString("MongoDBName")
	log.Println("Using database:", dbName)

	// Initialize collections
	Admins = client.Database(dbName).Collection("admins")
	Complains = client.Database(dbName).Collection("complains")
	Notices = client.Database(dbName).Collection("notices")
	Students = client.Database(dbName).Collection("students")
	Sclasses = client.Database(dbName).Collection("sclasses")
	Subjects = client.Database(dbName).Collection("subjects")
	Teachers = client.Database(dbName).Collection("teachers")
	Image = client.Database(dbName).Collection("image")

	Bucket, err = gridfs.NewBucket(client.Database(dbName))
	if err != nil {
		panic(err)
	}

}

func CloseClientDB() error {
	if client == nil {
		log.Println("MongoDB client is already nil, skipping disconnect.")
		return nil
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal("Error disconnecting MongoDB client: ", err)
		return err
	}
	log.Println("MongoDB connection closed.")
	return nil
}
