package handler

import (
	"ModanisaBE/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	service service.Service
}

func CreateNewController(service *service.Service) *Controller {
	c := Controller{service: *service}
	return &c
}

func (c *Controller) MainController(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	header.Add("Access-Control-Allow-Origin", "http://localhost:8080")

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)

		list, _ := json.Marshal(c.service.GetItems())
		_, err := w.Write(list)
		if err != nil {
			log.Print(err)
		}
		return
	}

	if r.Method == http.MethodPut {
		w.WriteHeader(http.StatusCreated)
		item, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		c.service.AddItem(string(item))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
