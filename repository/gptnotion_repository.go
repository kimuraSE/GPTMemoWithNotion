package repository

import (
	"GPTMemoWithNotion/Backend/model"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

type IGPTNotionRepository interface {
	CreateNotionPage(gptnotion model.GPTNotionRequest) error
}

type gptNotionRepository struct {
}

func NewGPTNotionRepository() IGPTNotionRepository {
	return &gptNotionRepository{}
}

func (gr *gptNotionRepository) CreateNotionPage(gptnotion model.GPTNotionRequest) error {
	
	err := godotenv.Load()
	if err != nil {
		return err
	}

	openaiapikey := os.Getenv("OPENAI_API_KEY")
	notionPageId := os.Getenv("NOTION_PAGE_ID")
	notionApiKey := os.Getenv("NOTION_API_KEY")

	client := openai.NewClient(openaiapikey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(`
					"%s"について結論→なぜ→例→まとめの順に1文でマークアップを用いらずに教えてください
					`, gptnotion.Question),
				},
			},
		},
	)

	if err != nil {
		return err
	}

	answer:=resp.Choices[0].Message.Content


	body := strings.NewReader(fmt.Sprintf(`
	{
		"parent": {
		  "page_id": "%s"
		},
		"properties": {
		  "title": {
			"title": [
				{
				  "text": {
					"content": "%s"
				  }
				}
			]
		  }
		},
		"children": [
		  {
			"object": "block",
			"type": "heading_2",
			"heading_2": {
			  "text": [
				{
				  "type": "text",
				  "text": {
					"content": "%s"
				  }
				}
			  ]
			}
		  },
		  {
			"object": "block",
			"type": "paragraph",
			"paragraph": {
			  "text": [
				{
				  "type": "text",
				  "text": {
					"content": "%s"
				  }
				}
			  ]
			}
		  }
		]
	  }

	`,notionPageId,gptnotion.Title,gptnotion.Headline,answer))

	req, err := http.NewRequest("POST", "https://api.notion.com/v1/pages", body)
	if err != nil {
		return err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + notionApiKey)
	req.Header.Add("Notion-Version", "2021-05-13")
	req.Header.Add("content-Type", "application/json")
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	return nil
}