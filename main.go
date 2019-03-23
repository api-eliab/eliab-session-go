package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	apigolang "github.com/josuegiron/api-golang"
)

func main() {
	LoadConfiguration()
	router := mux.NewRouter()

	middlewares := apigolang.MiddlewaresChain(apigolang.BasicAuth)

	router.HandleFunc("/v1.0/session", middlewares(login)).Methods("POST")

	log.Println("Starting server on port ", config.General.ServerAddress)
	if startServerError := http.ListenAndServe(config.General.ServerAddress, router); startServerError != nil {
		panic(startServerError)
	}
}
