package repository

import (
	"GPTMemoWithNotion/Backend/model"
	"context"
	"os"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

type IChatGPTRepository interface {
	GetAnswer(chatgpt model.ChatGPTRequest) (string, error)
}

type chatGPTRepository struct {
}

func NewChatGPTRepository() IChatGPTRepository {
	return &chatGPTRepository{}
}

func (cr *chatGPTRepository) GetAnswer(chatgpt model.ChatGPTRequest) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	openaiapikey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(openaiapikey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: chatgpt.Question,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	answer:=resp.Choices[0].Message.Content

	return answer, nil
}