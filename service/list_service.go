package service

import (
	"ModanisaBE/repository"
)

type Service struct {
	Data repository.DoList
}

func CreateNewService(doList *repository.DoList) *Service {
	service := Service{
		Data: *doList,
	}
	return &service
}

func (s *Service) AddItem(item string) {
	s.Data.AddItem(item)
}

func (s *Service) GetItems() []string {
	return s.Data.Items
}
