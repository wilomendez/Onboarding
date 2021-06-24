package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/wilomendez/Onboarding/update-status-contact-aws-lambda/controllers"
	"github.com/wilomendez/Onboarding/update-status-contact-aws-lambda/db"
	"github.com/wilomendez/Onboarding/update-status-contact-aws-lambda/services"
)

func main() {
	log.Println("Creating Handler!")
	service := services.New(db.New())
	getHandler := controllers.New(service)
	log.Println("Start exec...")
	lambda.Start(getHandler.HandleTrigger)
}
