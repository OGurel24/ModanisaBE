package main

import (
	"ModanisaBE/handler"
	"ModanisaBE/repository"
	"log"
	"net/http"
)

func main() {
	repositoryList := repository.CreateNewList()
	repositoryList = repositoryList.AddItem("add")

	http.HandleFunc("/", handler.MainController)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
