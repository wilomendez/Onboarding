package controllers

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/wilomendez/Onboarding/notify-aws-lambda/models"
	"github.com/wilomendez/Onboarding/notify-aws-lambda/services"
)

type NotifyContactHandler struct {
	NotifyService services.DomainNotifyService
}

type INotifyContactHandler interface {
	HandleTrigger(ctx context.Context, evt events.DynamoDBEvent)
}

func New(notifyService services.DomainNotifyService) NotifyContactHandler {
	return NotifyContactHandler{
		NotifyService: notifyService,
	}
}

func (nh NotifyContactHandler) HandleTrigger(ctx context.Context, evt events.DynamoDBEvent) {
	var attr string
	for _, record := range evt.Records {
		log.Println(record.EventID, record.EventName, record.EventSource, record.Change.NewImage)
		if record.EventName == "INSERT" {
			message := models.Message{
				EventID:     record.EventID,
				EventName:   record.EventName,
				EventSource: record.EventSource,
				Item: models.Data{
					Id:        record.Change.NewImage["id"].String(),
					FirstName: record.Change.NewImage["firstname"].String(),
					LastName:  record.Change.NewImage["lastname"].String(),
					Status:    record.Change.NewImage["status"].String(),
				},
			}
			attr = message.Item.Id
			nh.NotifyService.Push(message.String(), attr)
		}
	}
}
