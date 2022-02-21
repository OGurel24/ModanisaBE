package main

import (
	"ModanisaBE/handler"
	"ModanisaBE/repository"
	"ModanisaBE/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	listRepository := repository.CreateNewList()
	listService := service.CreateNewService(listRepository)
	fmt.Println(listService)

	http.HandleFunc("/", handler.MainController)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
