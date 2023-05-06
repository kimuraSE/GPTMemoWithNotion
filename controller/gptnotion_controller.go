package controller

import (
	"Backend/model"
	"Backend/usecase"
	"net/http"
	"github.com/labstack/echo/v4"
)

type IGPTMemoController interface {
	CreateNotionPage(c echo.Context) error
}

type gptNotionController struct {
	gu usecase.IGPTMemoUsecase
}

func NewGPTMemoController(gu usecase.IGPTMemoUsecase) IGPTMemoController {
	return &gptNotionController{gu}
}

func (gc *gptNotionController) CreateNotionPage(c echo.Context) error {
	gptnotion := model.GPTMemoRequest{}
	if err:= c.Bind(&gptnotion); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := gc.gu.CreateNotionPage(gptnotion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
