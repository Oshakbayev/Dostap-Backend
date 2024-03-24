package service

import "hellowWorldDeploy/pkg/entity"

type EventInterface interface {
	CreateEvent(*entity.Event) error
}

func (s *Service) CreateEvent(event *entity.Event) error {
	err := s.repo.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil
}
