package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/go-couch/apimodels"
	"github.com/RaboliotLeGris/go-couch/clients"
	"github.com/RaboliotLeGris/go-couch/dbmodels"
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

	tableName := generateTableName()

	if err := d.CouchDBClient.CreateTable(tableName); err != nil {
		http.Error(w, fmt.Sprintf("Post Documents - Error while CreateTable - %s", err), http.StatusInternalServerError)
		return
	}

	for _, document := range documents.Items {
		dbDocument := dbmodels.DocumentFromAPI(document)
		encodedDocument, err := json.Marshal(dbDocument)
		if err != nil {
			http.Error(w, fmt.Sprintf("Post Documents - Unable to marshal document - %s", err), http.StatusInternalServerError)
			return
		}

		if err = d.CouchDBClient.AddDocument(tableName, bytes.NewReader(encodedDocument)); err != nil {
			http.Error(w, fmt.Sprintf("Post Documents - Unable to create the document - %s", err), http.StatusInternalServerError)
			return
		}
	}
}

func generateTableName() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
