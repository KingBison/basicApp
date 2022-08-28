package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"main/logger"
	"main/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddUser(client *mongo.Client) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Error("Error Accepting Body")
		}
		var reader models.User
		err = json.Unmarshal(body, &reader)
		if err != nil {
			logger.Error("Error Unmarshalling Body")
		}
		reader.Active = true
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client.Database("basicapp").Collection("users").InsertOne(ctx, reader)

	}
}
