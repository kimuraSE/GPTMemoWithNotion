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

	chatgptUsercase := usecase.NewChatGPTUsecase(chatgptRepository)
	notionUsecase := usecase.NewNotionUsecase(notionRepository)

	chatGPTController := controller.NewChatGPTController(chatgptUsercase)
	notionController := controller.NewNotionController(notionUsecase)

	e := router.NewRouter(chatGPTController, notionController)
	e.Logger.Fatal(e.Start(":8080"))
}