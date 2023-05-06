package router

import (
	"Backend/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(gc controller.IGPTMemoController) *echo.Echo {
	e := echo.New()
	e.POST("/createpage", gc.CreateNotionPage)
	return e
}
