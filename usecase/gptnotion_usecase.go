package usecase

import (
	"Backend/model"
	"Backend/repository"
)

type IGPTMemoUsecase interface {
	CreateNotionPage(gptnotion model.GPTMemoRequest) error
}

type gptMemoUsecase struct {
	gr repository.IGPTMemoRepository
}

func NewGPTMemoUsecase(gr repository.IGPTMemoRepository) IGPTMemoUsecase {
	return &gptMemoUsecase{gr}
}

func (gu *gptMemoUsecase) CreateNotionPage(gptnotion model.GPTMemoRequest) error {
	
	err := gu.gr.CreateNotionPage(gptnotion)
	if err != nil {
		return err
	}

	return nil
}
