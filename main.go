package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/go-couch/router"
)

func main() {
	log.Info("Starting http server")

	defaultPort := 7777

	r := router.Create_router()

	if err := router.LaunchServer(r, defaultPort); err != nil {
		log.Error("Server stopped with error:", err)
	}
}
