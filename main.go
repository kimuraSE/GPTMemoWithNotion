package main

import (
	"GPTMemoWithNotion/Backend/controller"
	"GPTMemoWithNotion/Backend/repository"
	"GPTMemoWithNotion/Backend/router"
	"GPTMemoWithNotion/Backend/usecase"
)

func main(){
	chatgptRepository := repository.NewChatGPTRepository()
	chatgptUsercase := usecase.NewChatGPTUsecase(chatgptRepository)
	chatGPTController := controller.NewChatGPTController(chatgptUsercase)
	e := router.NewRouter(chatGPTController)
	e.Logger.Fatal(e.Start(":8080"))
}