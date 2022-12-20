package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/go-couch/apimodels"
)

type PostDocuments struct {
}

func (d PostDocuments) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var documents apimodels.Documents

	if err := json.NewDecoder(r.Body).Decode(&documents); err != nil {
		log.Debug("Failed to unmarshal body with error: ", err)
		w.WriteHeader(400)
		return
	}

	log.Debug("List of documents received", documents)
}
