package router

import (
	"GPTMemoWithNotion/Backend/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(cc controller.IChatGPTController) *echo.Echo {
	e := echo.New()
	e.POST("/chatgpt", cc.GetAnswer)
	return e
}
