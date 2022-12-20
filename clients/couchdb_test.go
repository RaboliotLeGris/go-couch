package clients_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/RaboliotLeGris/go-couch/clients"
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
