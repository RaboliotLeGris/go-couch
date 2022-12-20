package clients_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/RaboliotLeGris/go-couch/clients"
	"github.com/RaboliotLeGris/go-couch/dbmodels"
)

func Test_Create_Database_With_Empty_Table_Name(t *testing.T) {
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")

	require.Error(t, couchClient.CreateTable(""))
}

func Test_Create_Database(t *testing.T) {
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")

	require.NoError(t, couchClient.CreateTable("a_name"))
}

func Test_Create_Database_(t *testing.T) {
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")

	require.NoError(t, couchClient.CreateTable("collision"))
	require.Error(t, couchClient.CreateTable("collision"))
}

func Test_Create_Document(t *testing.T) {
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")

	tableName := "test_table"
	document1 := dbmodels.Document{
		Title:   "title1",
		Content: "content1",
		Author:  "author1",
	}
	document2 := dbmodels.Document{
		Title:   "title2",
		Content: "content2",
		Author:  "author2",
	}

	_ = couchClient.CreateTable(tableName)

	encodedDocument, err := json.Marshal(document1)
	require.NoError(t, err)
	require.NoError(t, couchClient.AddDocument(tableName, bytes.NewReader(encodedDocument)))

	encodedDocument, err = json.Marshal(document2)
	require.NoError(t, err)
	require.NoError(t, couchClient.AddDocument(tableName, bytes.NewReader(encodedDocument)))
}
