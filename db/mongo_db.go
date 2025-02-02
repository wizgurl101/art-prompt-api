package db

import (
	"art-prompt-api/models"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	mongo_db_uri := os.Getenv("MONGODB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongo_db_uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, err
	}

	return client, nil
}

func GetCollection(collectionName string) *mongo.Collection {
	client, err := getClient()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return client.Database("art-prompt").Collection(collectionName)
}

func GetUser(email string) (models.User, error) {
	filter := bson.D{{Key: "email", Value: email}}

	collection := GetCollection("users")
	var result models.User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.User{}, err
	}

	return result, nil
}
