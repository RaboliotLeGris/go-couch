package apimodels

type Documents struct {
	Items []Document `json:"items"`
}

type Document struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
