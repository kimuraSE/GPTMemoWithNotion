package usecase

import (
	"GPTMemoWithNotion/Backend/model"
	"GPTMemoWithNotion/Backend/repository"
)

type IGPTNotionUsecase interface {
	CreateNotionPage(gptnotion model.GPTNotionRequest) error
}

type gptNotionUsecase struct {
	gr repository.IGPTNotionRepository
}

func NewGPTNotionUsecase(gr repository.IGPTNotionRepository) IGPTNotionUsecase {
	return &gptNotionUsecase{gr}
}

func (gu *gptNotionUsecase) CreateNotionPage(gptnotion model.GPTNotionRequest) error {
	
	err := gu.gr.CreateNotionPage(gptnotion)
	if err != nil {
		return err
	}

	return nil
}
