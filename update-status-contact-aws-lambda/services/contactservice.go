package services

import (
	"github.com/wilomendez/Onboarding/update-status-contact-aws-lambda/models"
	"github.com/wilomendez/Onboarding/update-status-contact-aws-lambda/utils"
)

type DomainContactService struct {
	ContactRepo models.IContactsRepo
}

type IContactServices interface {
	Find(id string) (models.Contacts, error)
	ChangeStatus(id string) error
}

func New(contactRepo models.IContactsRepo) DomainContactService {
	return DomainContactService{
		ContactRepo: contactRepo,
	}
}

func (cs *DomainContactService) Find(id string) (models.Contacts, error) {
	if id == "" || id == " " {
		return models.Contacts{}, utils.ValidationError("Id field is Empty")
	}
	return cs.ContactRepo.Find(id)
}

func (cs *DomainContactService) ChangeStatus(id string) error {
	if id == "" {
		return utils.ValidationError("Id field is Empty")
	}
	return cs.ContactRepo.ChangeStatus(id)
}
