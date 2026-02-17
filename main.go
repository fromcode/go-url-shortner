package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL struct {
	Url      string `bson:"url"`
	ShortUrl string `bson:"shortUrl"`
}

var client *mongo.Client
var urlColl *mongo.Collection

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	uri := os.Getenv("")
	if uri == "" {
		log.Fatal("MONGODB URI is empty, check if its exist or not")
	}
}
