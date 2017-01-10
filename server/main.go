package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Main() {

	router := mux.NewRouter()

	// web ui path
	rootSubRouter := router.PathPrefix("/").Subrouter()
	rootSubRouter.HandleFunc("/", IndexHandler)

	// api path
	apiSubRouter := router.PathPrefix("/api").Subrouter()
	apiSubRouter.HandleFunc("/license", LicenseHandler)

	log.Fatal(http.ListenAndServe(":8881", router))
}
