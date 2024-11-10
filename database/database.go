package database

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MONGO_URI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}
	return client
}

var Client *mongo.Client = NewClient()

func OpenConnection(c *mongo.Client, collection string) *mongo.Collection {
	MONGO_URI := os.Getenv("MONGO_URI")
	return c.Database(filepath.Base(MONGO_URI)).Collection(collection)
}
