package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
    
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Err while loading env")
	}

	MongoDB_URL = os.Getenv("MONGO_URL") // we need to Apply this URI to a new client
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB_URL))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Backgound(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo")
	return client

}

 var Client *mongo.Client =DBInstance()

 func OpenCollection(client *mongo.Client ,collectionName string ) *mongo.Collection{
	collection *mongo.Collection := client.Database("cluster0").Collection(collectionName)
	return collection


 }
