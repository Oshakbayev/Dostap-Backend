package service

import (
	"hellowWorldDeploy/bucket"
	"hellowWorldDeploy/pkg/repo"
	"log"
)

type SvcInterface interface {
	UserServiceInterface
	ProfileServiceInterface
	EventInterface
	FriendInterface
	InterestInterface
}

type Service struct {
	log  *log.Logger
	repo repo.RepInterface
	bc bucket.BucketInterface
}

func CreateService(repo repo.RepInterface, l *log.Logger, bc bucket.BucketInterface) SvcInterface {
	return &Service{repo: repo, log: l, bc: bc}
}
