package repository_test

import (
	"ModanisaBE/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialDoList(t *testing.T) {
	DoList := repository.CreateNewList()

	testCases := []struct {
		description   string
		expectedValue string
		length        int
	}{
		{
			description:   "There should be 1 member of the initial list",
			expectedValue: "Achieve Modanisa Bootcamp assignment",
			length:        1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.length, len(DoList.Items))
			assert.Equal(t, testCase.expectedValue, DoList.Items[0])
		})
	}
}
