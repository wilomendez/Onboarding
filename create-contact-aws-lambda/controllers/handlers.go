package controllers

import (
	"context"
	"log"

	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/models"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/services"
)

type CreateContactHandler struct{}

type ICreateContactHandler interface {
	Create(ctx context.Context, request models.Request) (models.Response, error)
}

func (cch *CreateContactHandler) Create(ctx context.Context, request models.Request) (models.Response, error) {
	db := services.New()
	if request.FirstName == "" || request.LastName == "" {
		log.Println("Input validation FAILED, First Name or Last name is empty")
		return models.Response{}, errEmptyData
	}
	contact := new(models.Contacts)
	contact.FirstName = request.FirstName
	contact.LastName = request.LastName
	response, err := db.Create(*contact)
	if err != nil {
		log.Println(err.Error())
		return models.Response{}, HandleCustomValidationError(err.Error())
	}
	return response, nil
}
