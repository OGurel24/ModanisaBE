package main

import (
	"ModanisaBE/handler"
	"ModanisaBE/repository"
	"ModanisaBE/service"
	"log"
	"net/http"
)

func main() {
	listRepository := repository.CreateNewList()
	listService := service.CreateNewService(listRepository)
	listController := handler.CreateNewController(listService)

	http.HandleFunc("/", listController.MainController)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
