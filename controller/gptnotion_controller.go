package controller

import (
	"GPTMemoWithNotion/Backend/model"
	"GPTMemoWithNotion/Backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IGPTNotionController interface {
	CreateNotionPage(c echo.Context) error
}

type gptNotionController struct {
	gu usecase.IGPTNotionUsecase
}

func NewGPTNotionController(gu usecase.IGPTNotionUsecase) IGPTNotionController {
	return &gptNotionController{gu}
}

func (gc *gptNotionController) CreateNotionPage(c echo.Context) error {
	gptnotion := model.GPTNotionRequest{}
	if err:= c.Bind(&gptnotion); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := gc.gu.CreateNotionPage(gptnotion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
