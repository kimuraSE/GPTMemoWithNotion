package repository

import (
	"GPTMemoWithNotion/Backend/model"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type INotionRepository interface {
	CreatePage(notion model.NotionRequest) error
}

type notionRepository struct {
}

func NewNotionRepository() INotionRepository {
	return &notionRepository{}
}

func (nr *notionRepository) CreatePage(notion model.NotionRequest) error {
	
	err := godotenv.Load()
	if err != nil {
		return err
	}
	notionPageId := os.Getenv("NOTION_PAGE_ID")
	notionApiKey := os.Getenv("NOTION_API_KEY")

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

	`,notionPageId,notion.Title,notion.Headline,notion.Content))

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