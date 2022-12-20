package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/RaboliotLeGris/go-couch/apimodels"
	"github.com/RaboliotLeGris/go-couch/router"
)

func Test_PostDocuments_With_Empty_Body(t *testing.T) {

	recorder := httptest.NewRecorder()
	r := router.Create_router()

	req := httptest.NewRequest("POST", "/api/v1/documents", nil)
	r.ServeHTTP(recorder, req)

	require.Equal(t, 200, recorder.Code)
}

func Test_PostDocuments(t *testing.T) {
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
	r := router.Create_router()

	req := httptest.NewRequest("POST", "/api/v1/documents", body)
	r.ServeHTTP(recorder, req)

	require.Equal(t, 200, recorder.Code)
}
