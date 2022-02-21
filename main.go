package main

import (
	"fmt"
	_ "fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainController)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}

func mainController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	_, err := fmt.Fprintf(w, "User is already exist")
	if err != nil {
		fmt.Println(err)
	}
}
