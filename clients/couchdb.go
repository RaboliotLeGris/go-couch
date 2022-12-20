package clients

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

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

func (c CouchDBClient) AddDocument(table string, data io.Reader) error {
	client := &http.Client{}

	requestURI := fmt.Sprintf("%s/%s", c.addr, table)
	log.Info("AddDocument - Request URI: ", requestURI)

	req, err := http.NewRequest(http.MethodPost, requestURI, data)
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
