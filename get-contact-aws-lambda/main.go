package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/controllers"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/db"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/services"
)

func main() {
	log.Println("Creating Handler!")
	contactService := services.New(db.New())
	getHandler := controllers.New(contactService)
	log.Println("Start exec...")
	lambda.Start(getHandler.Find)
}
