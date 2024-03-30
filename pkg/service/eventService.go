package service

import (
	"hellowWorldDeploy/pkg/entity"
)

type EventInterface interface {
	CreateEvent(*entity.Event) error
	GetEventsByInterests([]int) ([]entity.Event, error)
	GetAllEvents() ([]entity.Event, error)
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
	err = s.repo.CreateEventInterests(event.ID, event.InterestIDs)
	if err != nil {
		s.log.Printf("\nError CreateEvent(service) during CreateEventInterests : %s\n", err.Error())
		return err
	}
	//log.Println(event.ID, "----Service")
	return nil
}

func (s *Service) GetEventsByInterests(interests []int) ([]entity.Event, error) {
	events, err := s.repo.GetEventsByInterests(interests)
	if err != nil {
		s.log.Printf("\nError GetEventsByInterests(service): %s\n", err.Error())
		return nil, err
	}
	return events, nil
}

func (s *Service) GetAllEvents() ([]entity.Event, error) {
	events, err := s.repo.GetAllEvents()
	if err != nil {
		s.log.Printf("\nError GetAllEvents(service): %s\n", err.Error())
		return nil, err
	}
	return events, nil
}
