package main

import (
	"Backend/controller"
	"Backend/repository"
	"Backend/router"
	"Backend/usecase"
)

func main(){
	
	gptMemoRepository := repository.NewGPTMemoRepository()
	gptMemoUsecase := usecase.NewGPTMemoUsecase(gptMemoRepository)
	gptMemoController := controller.NewGPTMemoController(gptMemoUsecase)

	e := router.NewRouter(gptMemoController)
	e.Logger.Fatal(e.Start(":8080"))
}