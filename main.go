package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type URL struct {
	Url      string `bson:"url"`
	ShortUrl string `bson:"shortUrl"`
}

var client *mongo.Client
var urlColl *mongo.Collection

func main() {
	r := gin.Default()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load .env")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB URI is empty, check if its exist or not")
	}

	var err error
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	urlColl = client.Database("url-shortner").Collection("urls")
}
