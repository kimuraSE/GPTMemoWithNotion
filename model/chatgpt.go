package model

type ChatGPTRequest struct {
	Question string `json:"question"`
}

type ChatGPTResponse struct {
	Answer string `json:"answer"`
}