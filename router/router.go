package router

import (
	"GPTMemoWithNotion/Backend/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(cc controller.IChatGPTController , nc controller.INotionController,gc controller.IGPTNotionController) *echo.Echo {
	e := echo.New()
	e.POST("/chatgpt", cc.GetAnswer)
	e.POST("/notion", nc.CreatePage)
	e.POST("/createpage", gc.CreateNotionPage)
	return e
}
