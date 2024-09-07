package mongodb

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMDB() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("uri"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDBga ulanishda xatolik: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDBga ulanishni tekshirishda xatolik: %v", err)
	}

	collection := client.Database("User_service").Collection("user")

	return collection, nil
}
