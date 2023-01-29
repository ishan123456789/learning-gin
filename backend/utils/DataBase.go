package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var DataBase *mongo.Database

// Connect : Connect to MongoDB
func Connect() (*mongo.Client, error) {
	if client == nil {
		// Get MongoDB URL from environment variable
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		mongoURL := os.Getenv("MONGO_URL")

		if mongoURL == "" {
			return nil, fmt.Errorf("MONGO_URL environment variable not set")
		}

		// Get MongoDB Database Name from environment variable
		mongoDB := os.Getenv("MONGO_DB")
		if mongoDB == "" {
			return nil, fmt.Errorf("MONGO_DB environment variable not set")
		}

		// Connect to MongoDB
		client, err = mongo.NewClient(options.Client().ApplyURI(mongoURL))
		if err != nil {
			return nil, err
		}
		err = client.Connect(context.TODO())
		if err != nil {
			return nil, err
		}
		// Check the connection
		pingError := client.Ping(context.TODO(), nil)
		if pingError != nil {
			log.Fatal(err)
		}
	}

	// Return MongoDB client
	return client, nil
}

func GetDB() (*mongo.Database, error) {
	localClient, err := Connect()
	if err != nil {
		return nil, err
	}
	// Get MongoDB Database Name from environment variable
	mongoDB := os.Getenv("MONGO_DB")
	if mongoDB == "" {
		return nil, fmt.Errorf("MONGO_DB environment variable not set")
	}
	DataBase = localClient.Database(mongoDB)
	return DataBase, nil
}

func GymCollection() *mongo.Collection {
	DB, err := GetDB()
	collection := DB.Collection("gyms")
	if err != nil {
		log.Fatal(err)
	}
	return collection
}
