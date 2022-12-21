package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RaboliotLeGris/go-couch/dbmodels"
	log "github.com/sirupsen/logrus"
)

type CouchDBClient struct {
	addr     string
	user     string
	password string
	bulkChan chan BulkJob
}

func NewCouchDBClient(addr, user, password string) CouchDBClient {
	client := CouchDBClient{
		addr:     addr,
		user:     user,
		password: password,
	}

	client.bulkChan = client.StartBulkPool()

	return client
}

func (c CouchDBClient) CreateTable(table string) error {
	if table == "" {
		return fmt.Errorf("empty table name is not allowed")
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
		return fmt.Errorf("database '%s' already exists", table)
	}

	if resp.StatusCode != http.StatusCreated {
		bodyContent := "N/A"
		if body, err := io.ReadAll(resp.Body); err != nil {
			bodyContent = string(body)
		}
		return fmt.Errorf("CreateTable failed with status %d - %s", resp.StatusCode, bodyContent)
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
		if body, err := io.ReadAll(resp.Body); err != nil {
			bodyContent = string(body)
		}
		return fmt.Errorf("AddDocument failed with status %d - %s", resp.StatusCode, bodyContent)
	}

	return nil
}

func (c CouchDBClient) AddDocumentsBulk(table string, docs []dbmodels.Document) error {
	if len(docs) == 0 {
		return fmt.Errorf("empty list of docs to add")
	}

	docsCount := len(docs)
	lowerBound := 0
	upperBound := 100
	if docsCount < 100 {
		upperBound = docsCount
	}

	// The downside of this implementation it is that the user won't know when his documents are inserted
	// But as long as we use one client, we are garanted to only have at most 8 parallel requests
	for upperBound <= docsCount {
		log.Debug("Bulk bounds: ", lowerBound, upperBound)
		slice := docs[lowerBound:upperBound]

		c.bulkChan <- BulkJob{Table: table, Documents: slice}

		lowerBound += 100
		if upperBound == docsCount {
			break
		} else if upperBound+100 > docsCount {
			upperBound = docsCount
		} else {
			upperBound += 100
		}
	}

	return nil
}

func (c CouchDBClient) bulkRequest(table string, bulk io.Reader) error {
	client := &http.Client{}

	requestURI := fmt.Sprintf("%s/%s/_bulk_docs", c.addr, table)

	req, err := http.NewRequest(http.MethodPost, requestURI, bulk)
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
		if body, err := io.ReadAll(resp.Body); err != nil {
			bodyContent = string(body)
		}
		return fmt.Errorf("AddDocumentBulk failed with status %d - %s", resp.StatusCode, bodyContent)
	}
	return nil
}

type BulkJob struct {
	Table     string
	Documents []dbmodels.Document
}

func (c CouchDBClient) StartBulkPool() chan BulkJob {
	ch := make(chan BulkJob, 8)
	for i := 0; i < 8; i++ {
		go c.BulkInsertWorker(ch)
	}
	return ch
}

type addDocumentBulk struct {
	Docs []dbmodels.Document `json:"docs"`
}

func (c CouchDBClient) BulkInsertWorker(ch chan BulkJob) {
	for bulk := range ch {
		log.Debug("Worker bulk length", len(bulk.Documents))

		bulkToInsert := addDocumentBulk{
			Docs: bulk.Documents,
		}
		encodedBulk, err := json.Marshal(bulkToInsert)
		if err != nil {
			log.Error("unable to marshal bulk - ", err)
			continue
		}

		err = c.bulkRequest(bulk.Table, bytes.NewReader(encodedBulk))
		if err != nil {
			log.Error("unable to set bulk to database - ", err)
		}
	}
}
