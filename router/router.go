package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/go-couch/clients"
	"github.com/RaboliotLeGris/go-couch/controllers"
)

func Create_router(couchDBClient clients.CouchDBClient) *mux.Router {
	log.Info("Creating routers")
	rootHandler := mux.NewRouter()

	v1 := rootHandler.PathPrefix("/api/v1").Subrouter()

	v1.PathPrefix("/documents").Handler(controllers.PostDocuments{CouchDBClient: couchDBClient}).Methods("POST")

	return rootHandler
}

func LaunchServer(router *mux.Router, port int) error {
	serverAddr := fmt.Sprintf("0.0.0.0:%v", port)
	log.Info("Launching HTTP server - available on :", serverAddr)

	srv := &http.Server{
		Handler: router,
		Addr:    serverAddr,
	}

	return srv.ListenAndServe()
}
