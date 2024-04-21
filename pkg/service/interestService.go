package service

import (
	"hellowWorldDeploy/pkg/entity"
)

type InterestInterface interface {
	GetAllInterests() ([]entity.Interest, error)
	GetUserInterests(int) ([]entity.Interest, error)
	CreateUserInterests(userId int, interests []int) error
}

func (s *Service) GetAllInterests() ([]entity.Interest, error) {
	allInterests, err := s.repo.GetAllInterests()
	if err != nil {
		s.log.Printf("error during CreateFriendRequest(service): %v", err)
		return nil, err
	}
	return allInterests, nil
}
func (s *Service) GetUserInterests(userId int) ([]entity.Interest, error) {
	return s.repo.GetUserInterests(userId)
}

func (s *Service) CreateUserInterests(userId int, interests []int) error {
	return s.repo.CreateUserInterests(userId, interests)
}

func (s *Service) UpdateUserInterests(userId int, interests []int) error {
	err := s.repo.DeleteUserInterests(userId)
	if err != nil {
		s.log.Printf("Error in UpdateUserProfileInfo %s", err.Error())
		return err
	}
	err = s.repo.CreateUserInterests(userId, interests)
	if err != nil {
		s.log.Printf("Error in UpdateUserProfileInfo %s", err.Error())

	}
	return err
}
