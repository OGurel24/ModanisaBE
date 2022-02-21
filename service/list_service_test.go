package service_test

import (
	"ModanisaBE/repository"
	"ModanisaBE/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_AddItem(t *testing.T) {
	s := service.CreateNewService(repository.CreateNewList())

	initialItemCount := len(s.Data.Items)

	testCases := []struct {
		Description   string
		LastItemCount int
	}{
		{
			Description:   "Item count should be increased by 1 after add operation",
			LastItemCount: initialItemCount + 1,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, 0, initialItemCount)
		s.AddItem("new item")
		t.Run(testCase.Description, func(t *testing.T) {
			assert.Equal(t, len(s.Data.Items), testCase.LastItemCount)
		})
	}
}

func TestService_GetItems(t *testing.T) {

	testCases := []struct {
		Description string
		Items       []string
	}{
		{
			Description: "Empty list should return 0 element string slice",
			Items:       nil,
		},
		{
			Description: "List with elements should return all elements as string slice",
			Items:       []string{"Wake Up", "Work on Modanisa Bootcamp", "Sleep", "Repeat"},
		},
	}

	for _, testCase := range testCases {
		s := service.CreateNewService(repository.CreateNewList())
		s.Data.Items = testCase.Items
		t.Run(testCase.Description, func(t *testing.T) {
			assert.Equal(t, testCase.Items, s.GetItems())
		})
	}
}
