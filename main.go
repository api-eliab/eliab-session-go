package main

import (
	"database/sql"
	"net/http"

	"github.com/jgolang/log"

	config "github.com/api-eliab/eliab-config-go"
	"github.com/gorilla/mux"
	"github.com/jgolang/apirest"
)

var latestSupportedVersion string

var db *sql.DB

func init() {
	db = config.DB
}

func main() {

	port := config.Get.General["SESSIONS"].PortServer

	err := convertLatestSupportedAppVersionThreePart()
	if err != nil {
		log.Panic(err)
	}

	router := mux.NewRouter()

	middlewares := apirest.MiddlewaresChain(apirest.BasicAuth, apirest.RequestHeaderJSON, apirest.GetRequestBodyMiddleware)

	router.HandleFunc("/v1.0/session", middlewares(login)).Methods("POST")
	router.HandleFunc("/v1.0/messages/users", middlewares(sendBroadcastMessageHandler)).Methods("POST")
	router.HandleFunc("/v1.0/messages/users/{userID}", middlewares(sendOneMessageHandler)).Methods("POST")
	//
	log.Println("Starting server on port ", port)
	if startServerError := http.ListenAndServe(port, router); startServerError != nil {
		log.Panic(startServerError)
	}

}
