package server_test

import (
	"ModanisaBE/server"
	"testing"
)

func TestServer_CreateServer(t *testing.T) {

	testCases := []struct {
		Description string
		port        int
	}{
		{
			Description: "Port number should not be error reason",
			port:        8082,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			go server.CreateServer(testCase.port)
			return
		})
	}
}
