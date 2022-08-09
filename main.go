package main

import (
	"main/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mongoClient := services.MongoClientGet()
	router := mux.NewRouter()
	router.HandleFunc("/mongo/info", services.MongoData(mongoClient)).Methods("POST")

	panic(http.ListenAndServe(":8090", router))

}
