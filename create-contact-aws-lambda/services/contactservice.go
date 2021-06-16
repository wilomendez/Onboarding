package services

import (
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/db"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/models"
)

type DomainContactService struct {
	ContactRepo models.IContactsRepo
}

type IContactServices interface {
	Create(contact models.Contacts) (models.Response, error)
}

func New() IContactServices {
	contactRepo := db.New()
	return &DomainContactService{
		ContactRepo: contactRepo,
	}
}

func (cs *DomainContactService) Create(contact models.Contacts) (models.Response, error) {
	contact.Fill_defaults() // Set UUID into ID and CREATED as status
	return cs.ContactRepo.Create(contact)
}
