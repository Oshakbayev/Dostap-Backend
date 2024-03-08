package service

import (
	"hellowWorldDeploy/pkg/entity"
	"hellowWorldDeploy/pkg/repo"
	"log"
)

type SvcInterface interface {
	SignUp(*entity.User) error
	LogIn(string, string) error
	TokenGenerator(string) (string, error)
	VerifyAccount(string) error
}

type Service struct {
	log  *log.Logger
	repo repo.RepInterface
}

func CreateService(repo repo.RepInterface, l *log.Logger) SvcInterface {
	return &Service{repo: repo, log: l}
}
