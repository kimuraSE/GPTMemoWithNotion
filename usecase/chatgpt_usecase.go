package usecase

import (
	"GPTMemoWithNotion/Backend/model"
	"GPTMemoWithNotion/Backend/repository"
)

type IChatGPTUsecase interface {
	GetAnswer(chatgpt model.ChatGPTRequest) (model.ChatGPTResponse, error)
}

type chatGPTUsecase struct {
	cr repository.IChatGPTRepository
}

func NewChatGPTUsecase(cr repository.IChatGPTRepository) IChatGPTUsecase {
	return &chatGPTUsecase{cr}
}

func (cu *chatGPTUsecase) GetAnswer(chatgpt model.ChatGPTRequest) (model.ChatGPTResponse, error) {

	asnwer,err := cu.cr.GetAnswer(chatgpt)
	if err != nil {
		return model.ChatGPTResponse{}, err
	}

	return model.ChatGPTResponse{Answer: asnwer}, nil
}
