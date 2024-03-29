package service

import (
	"hellowWorldDeploy/pkg/entity"
)

type EventInterface interface {
	CreateEvent(*entity.Event) error
	CreateEventInterests(*entity.Event) error
	GetEventsByInterests([]string) ([]entity.Event, error)
}

func (s *Service) CreateEvent(event *entity.Event) error {
	err := s.repo.CreateEvent(event)
	if err != nil {
		s.log.Printf("\nError CreateEvent(service): %s\n", err.Error())
		return err
	}
	err = s.repo.CreateEventOrganizers(event.ID, event.OrganizerIDs)
	if err != nil {
		s.log.Printf("\nError CreateEvent(service) during CreateEventOrganizers : %s\n", err.Error())
		return err
	}
	err = s.repo.CreateEventInterests(event.ID, event.EventInterests)
	if err != nil {
		s.log.Printf("\nError CreateEvent(service) during CreateEventInterests : %s\n", err.Error())
		return err
	}
	//log.Println(event.ID, "----Service")
	return nil
}

func (s *Service) CreateEventInterests(event *entity.Event) error {
	//log.Println(event.ID, "----createEventinterests service")
	err := s.repo.CreateEventInterests(event.ID, event.EventInterests)
	if err != nil {
		s.log.Printf("\nError CreateEventInterests(service): %s\n", err.Error())
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
