package services

import (
	"context"

	"log"
	"main/formatter"
	logger "main/logger"
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
	logger.Info("Connected to Database")
	return client
}

func MongoData(client *mongo.Client) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		dbNames, err := client.ListDatabases(ctx, bson.D{})
		if err != nil {
			logger.Error("Failed to get database names")
		}
		rtrn := formatter.FormatDBInfo(dbNames, client)
		rw.Write([]byte(rtrn))
	}
}
