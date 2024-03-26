package service

import "hellowWorldDeploy/pkg/entity"

type EventInterface interface {
	CreateEvent(*entity.Event) error
	GetEventsByInterests([]string) ([]entity.Event, error)
}

func (s *Service) CreateEvent(event *entity.Event) error {
	err := s.repo.CreateEvent(event)
	if err != nil {
		s.log.Printf("\nError CreateEvent(service): %s\n", err.Error())
		return err
	}
	return nil
}


func (s *Service) GetEventsByInterests(interests []string) ([]entity.Event, error) {
	events, err := s.repo.GetEventsByInterests(interests)
    if err != nil {
		s.log.Printf("\nError GetEventsByInterests(service): %s\n", err.Error())
        return nil, err
    }
    return events, nil
}