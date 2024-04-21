package service

import (
	"errors"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

type ProfileServiceInterface interface {
	UpdateUserProfileInfo(*entity.User) (int, error)
}

func (s *Service) UpdateUserProfileInfo(updatedUser *entity.User) (int, error) {
	err := s.repo.UpdateUserByID(updatedUser)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			s.log.Printf("ErrNoRowsErrNoRows(UpdateUserByID) in UpdateUserProfileInfo")
			return http.StatusBadRequest, errors.New("user does not exist")
		}
		s.log.Printf("Error in UpdateUserProfileInfo %s", err.Error())
		return http.StatusInternalServerError, err
	}
	err = s.UpdateUserInterests(updatedUser.ID, updatedUser.Interests)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
