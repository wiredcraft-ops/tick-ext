package client

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Main() {

	router := mux.NewRouter()

	// web ui path
	rootSubRouter := router.PathPrefix("/").Subrouter()
	rootSubRouter.HandleFunc("/", IndexHandler)

	// api path
	apiSubRouter := router.PathPrefix("/api").Subrouter()
	apiSubRouter.HandleFunc("/license", LicenseHandler)

	log.Fatal(http.ListenAndServe(":8880", router))
}
