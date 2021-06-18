package services

import (
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/models"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/utils"
)

type DomainContactService struct {
	ContactRepo models.IContactsRepo
}

type IContactServices interface {
	Create(contact models.Contacts) (models.Contacts, error)
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
