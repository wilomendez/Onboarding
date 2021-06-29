package models

import (
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/utils"
)

const statusCreated string = "CREATED"

type Contacts struct {
	Id        string `dynamodbav:"id"`
	FirstName string `dynamodbav:"firstname"`
	LastName  string `dynamodbav:"lastname"`
	Status    string `dynamodbav:"status"`
}

func (c *Contacts) Fill_defaults() {
	if c.Status == "" {
		c.Status = statusCreated
	}

	if c.Id == "" {
		c.Id = utils.UUIDStringGenerator()
	}
}

type IContactsRepo interface {
	Create(contact Contacts) (Response, error)
}
