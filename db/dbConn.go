package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//connect to mongodb cloud for db

func DBInstance() *mongo.Client {
	//get url from os.GetEnv
	mongoDB := os.Getenv("mongo")
	if mongoDB == "" {
		fmt.Printf("please provide mongodb url to connect")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDB))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//everyhting goes well, connect to mongodb
	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	//create a new table -> collection in mongodb
	var collection *mongo.Collection = client.Database("romeoDB").Collection(collectionName)
	return collection
}
