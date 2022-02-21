package handler_test

import (
	"ModanisaBE/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainController(t *testing.T) {

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
			statusCode:  417,
		},
		{
			description: "Put method should return HTTP 417 status code",
			method:      http.MethodPut,
			statusCode:  200,
		},
		{
			description: "Delete method should return HTTP 200 status code",
			method:      http.MethodDelete,
			statusCode:  417,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(testCase.method, "/", nil)
		w := httptest.NewRecorder()
		handler.MainController(w, req)
		res := w.Result()
		responseCode := res.StatusCode
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.statusCode, responseCode)
		})
	}
}
