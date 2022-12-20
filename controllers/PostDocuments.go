package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/go-couch/apimodels"
	"github.com/RaboliotLeGris/go-couch/clients"
)

type PostDocuments struct {
	CouchDBClient clients.CouchDBClient
}

func (d PostDocuments) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var documents apimodels.Documents

	if err := json.NewDecoder(r.Body).Decode(&documents); err != nil {
		log.Debug("Failed to unmarshal body with error: ", err)
		w.WriteHeader(400)
		return
	}

	log.Debug("List of documents received", documents)
	tableName := uuid.New().String()

	if err := d.CouchDBClient.CreateTable(tableName); err != nil {
		http.Error(w, fmt.Sprintf("Post Documents - Error while CreateTable - %s", err), http.StatusInternalServerError)
		return
	}
}
