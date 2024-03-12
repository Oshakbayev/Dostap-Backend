package service

import (
	"database/sql"
	"errors"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

type ProfileServiceInterface interface {
	UpdateUserProfileInfo(*entity.User) (int, error)
}

func (s *Service) UpdateUserProfileInfo(updatedUser *entity.User) (int, error) {
	err := s.repo.UpdateUser(updatedUser)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, errors.New("user does not exist")
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
