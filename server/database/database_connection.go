package database

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


func DBInstance() *mongo.Client {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoDbUrl := os.Getenv("MONGODB_URL")

    if mongoDbUrl == "" {
        log.Fatal("MONGODB_URL not found in .env file")
    }

	fmt.Println("Connecting to MongoDBURL :",mongoDbUrl)

    clientOptions := options.Client().ApplyURI(mongoDbUrl)

	client, err := mongo.Connect( clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	return  client

}


var Client *mongo.Client=DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	 err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
 
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatal("DATABASE_NAME not found in .env file")
	}
	 collection := Client.Database(databaseName).Collection(collectionName)

	 if collection == nil {
		log.Fatal("Error connecting to collection: ", collectionName)
	}

	 return collection
}
