package controllers

import (
	"context"
	"log"

	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/models"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/services"
)

type CreateContactHandler struct {
	service services.DomainContactService
}

type ICreateContactHandler interface {
	Create(ctx context.Context, request models.Request) (models.Response, error)
}

func New(contactservice services.DomainContactService) CreateContactHandler {
	return CreateContactHandler{
		service: contactservice,
	}
}

func (cch *CreateContactHandler) Create(ctx context.Context, request models.Request) (models.Response, error) {
	if request.FirstName == "" || request.LastName == "" {
		log.Println("Input validation FAILED, First Name or Last name is empty")
		return models.Response{}, errEmptyData
	}
	contact := new(models.Contacts)
	contact.FirstName = request.FirstName
	contact.LastName = request.LastName
	response, err := cch.service.Create(*contact)
	if err != nil {
		log.Println(err.Error())
		return models.Response{}, HandleCustomValidationError(err.Error())
	}
	return response, nil
}
