package main

import (
	"main/handlers"
	"main/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mongoClient := services.MongoClientGet()
	router := mux.NewRouter()
	router.HandleFunc("/mongo/info", services.MongoData(mongoClient))
	router.HandleFunc("/users/adduser", handlers.AddUser(mongoClient)).Methods("POST")

	panic(http.ListenAndServe(":8080", router))

}
