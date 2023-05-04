package usecase

import (
	"GPTMemoWithNotion/Backend/model"
	"GPTMemoWithNotion/Backend/repository"
)

type INotionUsecase interface {
	CreatePage(notion model.NotionRequest) error
}

type notionUsecase struct {
	nr repository.INotionRepository
}

func NewNotionUsecase(nr repository.INotionRepository) INotionUsecase {
	return &notionUsecase{nr}
}

func (nu *notionUsecase) CreatePage(notion model.NotionRequest) error{
	
	err := nu.nr.CreatePage(notion)
	if err != nil {
		return err
	}
	
	return nil

}
