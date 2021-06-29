package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/wilomendez/Onboarding/update-status-contact-aws-lambda/models"
	"github.com/wilomendez/Onboarding/update-status-contact-aws-lambda/services"
)

type UpdateContactHandler struct {
	service services.DomainContactService
}

type IUpdateContactHandler interface {
	HandleTrigger(ctx context.Context, snsEvent events.SNSEvent)
}

func New(contactservice services.DomainContactService) UpdateContactHandler {
	return UpdateContactHandler{
		service: contactservice,
	}
}

func (uh *UpdateContactHandler) HandleTrigger(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		response := models.Message{}
		if err := json.Unmarshal([]byte(snsRecord.Message), &response); err != nil {
			log.Println("Error to unmarshall record")
		}
		uh.service.ChangeStatus(response.Item.Id)
		fmt.Printf("[%s %s] Message = %s id= %v\n", record.EventSource, snsRecord.Timestamp, snsRecord.Message, response.Item.Id)
	}
}
