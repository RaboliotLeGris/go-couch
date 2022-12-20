package clients

import "encoding/base64"

type CouchDBClient struct {
	addr string
	auth string
}

func NewCouchDBClient(addr, user, password string) CouchDBClient {
	login := user + ":" + password
	hashed := base64.StdEncoding.EncodeToString([]byte(login))
	return CouchDBClient{
		addr: addr,
		auth: "Basic " + hashed,
	}
}

func (c CouchDBClient) CreateTable(name string) error {
	return nil
}

func (c CouchDBClient) AddDocument(tableName, data any) error {
	return nil
}
