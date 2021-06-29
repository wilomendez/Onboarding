package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/wilomendez/Onboarding/notify-aws-lambda/controllers"
	"github.com/wilomendez/Onboarding/notify-aws-lambda/notify"
	"github.com/wilomendez/Onboarding/notify-aws-lambda/services"
)

func main() {
	log.Println("Creating Handler!")
	service := services.New(notify.New())
	getHandler := controllers.New(service)
	log.Println("Start exec...")
	lambda.Start(getHandler.HandleTrigger)
}
