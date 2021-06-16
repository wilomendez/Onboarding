package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/controllers"
)

func main() {
	log.Println("Creating Handler!")
	createHandler := controllers.CreateContactHandler{}
	log.Println("Start exec...")
	lambda.Start(createHandler.Create)
}
