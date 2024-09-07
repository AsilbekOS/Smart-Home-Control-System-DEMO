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
	clientOptions := options.Client().ApplyURI(os.Getenv("apiurl"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDBga ulanishda xatolik: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDBga ulanishni tekshirishda xatolik: %v", err)
	}

	collection := client.Database("Device_service").Collection("device")
	return collection, nil
}
