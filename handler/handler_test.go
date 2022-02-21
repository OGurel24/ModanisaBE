package handler_test

import (
	"ModanisaBE/handler"
	"ModanisaBE/repository"
	"ModanisaBE/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_MainController(t *testing.T) {

	testCases := []struct {
		description  string
		method       string
		statusCode   int
		responseText string
	}{
		{
			description: "Get method should return HTTP 200 status code",
			method:      http.MethodGet,
			statusCode:  200,
		},
		{
			description: "Post method should return HTTP 417 status code",
			method:      http.MethodPost,
			statusCode:  405,
		},
		{
			description: "Put method should return HTTP 417 status code",
			method:      http.MethodPut,
			statusCode:  200,
		},
		{
			description: "Delete method should return HTTP 200 status code",
			method:      http.MethodDelete,
			statusCode:  405,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(testCase.method, "/", nil)
		w := httptest.NewRecorder()
		listHandler := handler.CreateNewController(service.CreateNewService(repository.CreateNewList()))
		listHandler.MainController(w, req)
		res := w.Result()
		responseCode := res.StatusCode
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.statusCode, responseCode)
		})
	}
}
