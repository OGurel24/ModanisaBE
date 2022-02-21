package service_test

import (
	"ModanisaBE/repository"
	"ModanisaBE/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddItem(t *testing.T) {
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
		assert.Equal(t, 1, initialItemCount)
		s = s.AddItem("new item")
		t.Run(testCase.Description, func(t *testing.T) {
			assert.Equal(t, len(s.Data.Items), testCase.LastItemCount)
		})
	}
}
