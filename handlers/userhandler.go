package handlers

import (
	"main/logger"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddUser(client *mongo.Client) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		logger.Debug(r.GetBody)
	}
}
