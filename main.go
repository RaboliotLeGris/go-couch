package main

import (
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/go-couch/clients"
	"github.com/RaboliotLeGris/go-couch/config"
	"github.com/RaboliotLeGris/go-couch/router"
)

func main() {
	log.Info("Starting http server")

	// Seed random with time
	rand.Seed(time.Now().UnixNano())

	cfg, err := config.NewConfig()
	if err != nil {
		log.Error("Unable to parse config, exiting")
		os.Exit(1)
	}

	couchDBClient := clients.NewCouchDBClient(cfg.CouchDbAddr, cfg.CouchDBUser, cfg.CouchDBPassword)

	r := router.Create_router(couchDBClient)

	if err := router.LaunchServer(r, cfg.Port); err != nil {
		log.Error("Server stopped with error:", err)
	}
}
