package services

import (
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/models"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/utils"
)

type DomainContactService struct {
	ContactRepo models.IContactsRepo
}

type IContactServices interface {
	Create(contact models.Contacts) (models.Response, error)
}

func New(contactRepo models.IContactsRepo) DomainContactService {
	return DomainContactService{
		ContactRepo: contactRepo,
	}
}

func (cs *DomainContactService) Create(contact models.Contacts) (models.Response, error) {
	if contact.FirstName == "" && contact.LastName == "" {
		return models.Response{}, utils.ValidationError("Input validation FAILED, First Name or Last name is empty")
	}
	contact.Fill_defaults() // Set UUID into ID and CREATED as status
	return cs.ContactRepo.Create(contact)
}
