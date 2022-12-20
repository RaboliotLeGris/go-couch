package controllers

import "net/http"

type PostDocuments struct {
}

func (d PostDocuments) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
