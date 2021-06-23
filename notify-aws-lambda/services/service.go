package services

import (
	"errors"

	"github.com/wilomendez/Onboarding/notify-aws-lambda/models"
)

type DomainNotifyService struct {
	Repo models.IMessageRepo
}

type INotifyService interface {
	Push(message string) error
}

func New(messageRepo models.IMessageRepo) DomainNotifyService {
	return DomainNotifyService{
		Repo: messageRepo,
	}
}

func (ns *DomainNotifyService) Push(message, attrs string) error {
	if message == "" {
		return errors.New("message not found")
	}
	if attrs == "" {
		return errors.New("attr not found")
	}
	return ns.Repo.Push(message, attrs)
}
