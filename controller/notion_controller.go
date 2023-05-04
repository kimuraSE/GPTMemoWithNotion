package controller

import (
	"GPTMemoWithNotion/Backend/model"
	"GPTMemoWithNotion/Backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type INotionController interface {
	CreatePage(e echo.Context) error
}

type notionController struct {
	nu usecase.INotionUsecase
}

func NewNotionController(nu usecase.INotionUsecase) INotionController {
	return &notionController{nu}
}

func (nc *notionController) CreatePage(c echo.Context) error {
	notion := model.NotionRequest{}
	if err:= c.Bind(&notion); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	notionRes,err := nc.nu.CreatePage(notion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notionRes)
}