package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RaboliotLeGris/go-couch/dbmodels"
	log "github.com/sirupsen/logrus"
)

type CouchDBClient struct {
	addr     string
	user     string
	password string
}

func NewCouchDBClient(addr, user, password string) CouchDBClient {
	return CouchDBClient{
		addr:     addr,
		user:     user,
		password: password,
	}
}

func (c CouchDBClient) CreateTable(table string) error {
	if table == "" {
		return errors.New("empty table name is not allowed")
	}

	client := &http.Client{}

	requestURI := fmt.Sprintf("%s/%s", c.addr, table)
	log.Debug("CreateTable - Request URI: ", requestURI)

	req, err := http.NewRequest(http.MethodPut, requestURI, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.user, c.password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusPreconditionFailed {
		return errors.New("database '" + table + "' already exists")
	}

	if resp.StatusCode != http.StatusCreated {
		bodyContent := "N/A"
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			bodyContent = string(body)
		}
		return errors.New(fmt.Sprintf("CreateTable failed with status %d - %s", resp.StatusCode, bodyContent))
	}

	return nil
}

func (c CouchDBClient) AddDocument(table string, document dbmodels.Document) error {
	client := &http.Client{}

	requestURI := fmt.Sprintf("%s/%s", c.addr, table)
	log.Info("AddDocument - Request URI: ", requestURI)

	encodedDocument, err := json.Marshal(document)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, requestURI, bytes.NewReader(encodedDocument))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.user, c.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		bodyContent := "N/A"
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			bodyContent = string(body)
		}
		return errors.New(fmt.Sprintf("AddDocument failed with status %d - %s", resp.StatusCode, bodyContent))
	}

	return nil
}

type addDocumentBulk struct {
	Docs []dbmodels.Document `json:"docs"`
}

func (c CouchDBClient) AddDocumentBulk(table string, docs []dbmodels.Document) error {
	client := &http.Client{}

	requestURI := fmt.Sprintf("%s/%s/_bulk_docs", c.addr, table)
	log.Info("AddDocumentBulk - Request URI: ", requestURI)

	bulkToInsert := addDocumentBulk{
		Docs: docs,
	}
	encodedBulk, err := json.Marshal(bulkToInsert)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, requestURI, bytes.NewReader(encodedBulk))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.user, c.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		bodyContent := "N/A"
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			bodyContent = string(body)
		}
		return errors.New(fmt.Sprintf("AddDocumentBulk failed with status %d - %s", resp.StatusCode, bodyContent))
	}

	return nil
}
