package main

import (
	"net/http"

	"github.com/josuegiron/log"

	"github.com/gorilla/mux"
	apigo "github.com/josuegiron/api-golang"
)

func main() {
	loadConfiguration()
	router := mux.NewRouter()
	log.ChangeCallerSkip(-2)

	if !dbConnect() {
		log.Panic("Error al conectar a la base de datos!")
	}

	middlewares := apigo.MiddlewaresChain(apigo.BasicAuth, apigo.RequestHeaderJson, apigo.GetRequestBodyMiddleware)

	router.HandleFunc("/v1.0/session", middlewares(login)).Methods("POST")
	//
	log.Println("Starting server on port ", config.General.ServerAddress)
	if startServerError := http.ListenAndServe(config.General.ServerAddress, router); startServerError != nil {
		log.Panic(startServerError)
	}
}

func init() {

}
