package dbmodels

import "github.com/RaboliotLeGris/go-couch/apimodels"

type Document struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func DocumentFromAPI(apiDocument apimodels.Document) Document {
	return Document{
		Title:   apiDocument.Title,
		Content: apiDocument.Content,
		Author:  apiDocument.Author,
	}
}
