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

func CreateNewController(service service.Service) Controller {
	c := Controller{service: service}
	return c
}

func (c *Controller) MainController(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusOK)
		item, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		c.service = c.service.AddItem(string(item))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
