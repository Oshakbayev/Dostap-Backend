package service

import "hellowWorldDeploy/pkg/entity"

type InterestInterface interface {
	GetAllInterests() ([]entity.Interest, error)
}

func (s *Service) GetAllInterests() ([]entity.Interest, error) {
	allInterests, err := s.repo.GetAllInterests()
	if err != nil {
		s.log.Printf("error during CreateFriendRequest(service): %v", err)
		return nil, err
	}
	return allInterests, nil
}
