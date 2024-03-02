package service

import (
	"fmt"
	"hellowWorldDeploy/pkg/entity"
)

func (s Service) SignUp(user *entity.User) error {
	err := s.repo.CreateUser(user)
	if err != nil {
		s.log.Printf("Error while Inserting user into the table at the service level")
		return fmt.Errorf("error while insert new user: %v, error: %s", user, err)
	}
	return nil
}
