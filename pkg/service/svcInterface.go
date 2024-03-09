package service

import (
	"hellowWorldDeploy/pkg/repo"
	"log"
)

type SvcInterface interface {
	UserServiceInterface
	ProfileServiceInterface
}

type Service struct {
	log  *log.Logger
	repo repo.RepInterface
}

func CreateService(repo repo.RepInterface, l *log.Logger) SvcInterface {
	return &Service{repo: repo, log: l}
}
