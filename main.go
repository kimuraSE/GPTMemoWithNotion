package main

import (
	"GPTMemoWithNotion/Backend/controller"
	"GPTMemoWithNotion/Backend/repository"
	"GPTMemoWithNotion/Backend/router"
	"GPTMemoWithNotion/Backend/usecase"
)

func main(){
	chatgptRepository := repository.NewChatGPTRepository()
	notionRepository := repository.NewNotionRepository()
	gptNotionRepository := repository.NewGPTNotionRepository()
	chatgptUsercase := usecase.NewChatGPTUsecase(chatgptRepository)
	notionUsecase := usecase.NewNotionUsecase(notionRepository)
	gptNotionUsecase := usecase.NewGPTNotionUsecase(gptNotionRepository)

	chatGPTController := controller.NewChatGPTController(chatgptUsercase)
	notionController := controller.NewNotionController(notionUsecase)
	gptNotionController := controller.NewGPTNotionController(gptNotionUsecase)

	e := router.NewRouter(chatGPTController, notionController, gptNotionController)
	e.Logger.Fatal(e.Start(":8080"))
}