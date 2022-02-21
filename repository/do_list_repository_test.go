package repository_test

import (
	"ModanisaBE/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialDoList(t *testing.T) {
	DoList := *repository.getList()

	testCases := []struct {
		description   string
		expectedValue string
	}{
		{
			description:   "there should be 1 member of the initial list",
			expectedValue: "Achieve Modanisa Bootcamp assignment",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, 1, len(DoList))
		})
	}
}
