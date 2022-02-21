package service

import "ModanisaBE/repository"

type Service struct {
	Data repository.DoList
}

func CreateNewService(doList repository.DoList) Service {
	service := Service{
		Data: doList,
	}
	return service
}

func (s *Service) AddItem(item string) Service {
	s.Data = s.Data.AddItem(item)
	return *s
}

func (s *Service) GetItems() []string {
	return s.Data.Items
}
