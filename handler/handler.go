package handler

import (
	"ModanisaBE/service"
	"net/http"
)

func MainController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		service.GetItems()
		return
	}
	if r.Method == http.MethodPut {
		w.WriteHeader(http.StatusOK)
		service.AddItem(r.Body)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
