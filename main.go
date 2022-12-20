package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/go-couch/clients"
	"github.com/RaboliotLeGris/go-couch/router"
)

func main() {
	log.Info("Starting http server")

	// Seed random with time
	rand.Seed(time.Now().UnixNano())

	port := 7777
	couchAddr := "http://127.0.0.1:5984"
	couchUser := "admin"
	couchPassword := "password"

	couchDBClient := clients.NewCouchDBClient(couchAddr, couchUser, couchPassword)

	r := router.Create_router(couchDBClient)

	if err := router.LaunchServer(r, port); err != nil {
		log.Error("Server stopped with error:", err)
	}
}
