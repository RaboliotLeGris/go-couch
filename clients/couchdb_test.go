package clients_test

import (
	"fmt"
	"testing"
	"time"

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

	require.NoError(t, couchClient.AddDocument(tableName, document1))
	require.NoError(t, couchClient.AddDocument(tableName, document2))
}

func Test_Create_Documents_Bulk(t *testing.T) {
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")
	tableName := "test_table_bulk"

	_ = couchClient.CreateTable(tableName)

	docs := []dbmodels.Document{}
	for i := 0; i < 201; i++ { // 2 bulk + 1
		docs = append(docs, dbmodels.Document{
			Title:   fmt.Sprintf("title %d", i),
			Content: fmt.Sprintf("content %d", i),
			Author:  fmt.Sprintf("author %d", i),
		})
	}

	require.NoError(t, couchClient.AddDocumentsBulk(tableName, docs))
}

func Test_Create_1995_Documents_Bulk(t *testing.T) {
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")
	tableName := "test_table_bulk_1995"

	_ = couchClient.CreateTable(tableName)

	docs := []dbmodels.Document{}
	for i := 0; i < 1995; i++ {
		docs = append(docs, dbmodels.Document{
			Title:   fmt.Sprintf("title %d", i),
			Content: fmt.Sprintf("content %d", i),
			Author:  fmt.Sprintf("author %d", i),
		})
	}

	require.NoError(t, couchClient.AddDocumentsBulk(tableName, docs))

	// Small race condition here because the test stop before the bulk is processed.
	// An improvement will be to wait until the chan is empty
	time.Sleep(time.Second * 2)
}

func Test_Create_50_Documents_Bulk(t *testing.T) {
	couchClient := clients.NewCouchDBClient("http://127.0.0.1:5984", "admin", "password")
	tableName := "test_table_bulk_50"

	_ = couchClient.CreateTable(tableName)

	docs := []dbmodels.Document{}
	for i := 0; i < 50; i++ { //
		docs = append(docs, dbmodels.Document{
			Title:   fmt.Sprintf("title %d", i),
			Content: fmt.Sprintf("content %d", i),
			Author:  fmt.Sprintf("author %d", i),
		})
	}

	require.NoError(t, couchClient.AddDocumentsBulk(tableName, docs))

	// Small race condition here because the test stop before the bulk is processed.
	// An improvement will be to wait until the chan is empty
	time.Sleep(time.Second * 2)
}
