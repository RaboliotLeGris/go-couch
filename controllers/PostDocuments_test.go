package controllers_test

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/RaboliotLeGris/go-couch/apimodels"
	"github.com/RaboliotLeGris/go-couch/clients"
	"github.com/RaboliotLeGris/go-couch/router"
)

func Test_PostDocuments_With_Empty_Body(t *testing.T) {
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")

	recorder := httptest.NewRecorder()
	r := router.Create_router(couchClient)

	req := httptest.NewRequest("POST", "/api/v1/documents", nil)
	r.ServeHTTP(recorder, req)

	require.Equal(t, 400, recorder.Code)
}

func Test_PostDocuments(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")

	giveDocuments := apimodels.Documents{
		Items: []apimodels.Document{
			{
				Title:   "Title",
				Content: "Content",
				Author:  "Author",
			},
			{
				Title:   "Title1",
				Content: "Content1",
				Author:  "Author1",
			},
		},
	}

	rawBuf, err := json.Marshal(giveDocuments)
	require.NoError(t, err)
	body := bytes.NewBuffer(rawBuf)

	recorder := httptest.NewRecorder()
	r := router.Create_router(couchClient)

	req := httptest.NewRequest("POST", "/api/v1/documents", body)
	r.ServeHTTP(recorder, req)

	require.Equal(t, 200, recorder.Code)
}
