package controllers

import (
	"context"
	"log"

	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/models"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/services"
)

type GetcontactHandler struct {
	service services.DomainContactService
}

type IGetContactHandler interface {
	Find(context.Context, models.Request) (models.Contacts, error)
}

func New(contactservice services.DomainContactService) GetcontactHandler {
	return GetcontactHandler{
		service: contactservice,
	}
}

func (gch *GetcontactHandler) Find(ctx context.Context, request models.Request) (models.Contacts, error) {
	if request.Id == "" || request.Id == " " {
		errMsg := "Input validation FAILED, id field is empty"
		log.Println(errMsg)
		return models.Contacts{}, errEmptyData
	}
	response, err := gch.service.Find(request.Id)
	if err != nil {
		log.Println(err.Error())
		return models.Contacts{}, err
	}
	return response, nil
}
