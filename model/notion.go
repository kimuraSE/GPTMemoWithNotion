package model

type NotionRequest struct {
	Title string `json:"title"`
	Headline string `json:"headline"`
	Content string `json:"content"`
}

type NotionResponse struct {
	Title string `json:"title"`
}
