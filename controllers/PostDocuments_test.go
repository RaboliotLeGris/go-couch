package controllers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/RaboliotLeGris/go-couch/router"
)

func Test_PostDocuments(t *testing.T) {
	recorder := httptest.NewRecorder()
	r := router.Create_router()

	req := httptest.NewRequest("POST", "/api/v1/documents", nil)
	r.ServeHTTP(recorder, req)

	require.Equal(t, 200, recorder.Code)
}
