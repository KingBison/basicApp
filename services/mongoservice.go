package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClientGet() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://rufus2012:Sgra0606@cluster0.sxu6rhe.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

func MongoData(client *mongo.Client) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		testdata := bson.M{
			"title":  "The Polyglot Developer Podcast",
			"author": "Nic Raboy",
			"tags":   bson.A{"development", "programming", "coding"},
		}
		testcoll := client.Database("testdb").Collection("test")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err := testcoll.InsertOne(ctx, testdata)
		if err != nil {
			panic("insert failed")
		}
	}
}
