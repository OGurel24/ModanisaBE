package server

import (
	"ModanisaBE/handler"
	"ModanisaBE/repository"
	"ModanisaBE/service"
	"fmt"
	"log"
	"net/http"
)

func CreateServer(port int) {
	listRepository := repository.CreateNewList()
	listService := service.CreateNewService(listRepository)
	listController := handler.CreateNewController(listService)

	http.HandleFunc("/", listController.MainController)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
