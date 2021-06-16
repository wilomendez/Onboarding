package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/controllers"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/db"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/services"
)

func main() {
	log.Println("Creating Handler!")
	contactService := services.New(db.New())
	createHandler := controllers.New(contactService)
	log.Println("Start exec...")
	lambda.Start(createHandler.Create)
}
