package controller

import (
	"GPTMemoWithNotion/Backend/model"
	"GPTMemoWithNotion/Backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)


type IChatGPTController interface {
	GetAnswer(e echo.Context) error
}

type chatGPTController struct {
	cu usecase.IChatGPTUsecase
}

func NewChatGPTController(cu usecase.IChatGPTUsecase) IChatGPTController {
	return &chatGPTController{cu}
}

func (cc *chatGPTController) GetAnswer(c echo.Context) error {
	chatgpt := model.ChatGPTRequest{}
	if err:= c.Bind(&chatgpt); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	chatgptRes,err := cc.cu.GetAnswer(chatgpt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, chatgptRes)
}




