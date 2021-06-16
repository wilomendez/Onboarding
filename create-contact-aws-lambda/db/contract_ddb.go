package db

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/wilomendez/Onboarding/create-contract-aws-lambda/models"
)

const (
	REGION    = "us-east-1"
	TABLENAME = "Contacts_WM"
)

type ContactService struct {
	TableName string
	ddb       *dynamodb.DynamoDB
}

func New() models.IContactsRepo {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(REGION)}))
	return &ContactService{
		TableName: TABLENAME,
		ddb:       dynamodb.New(sess),
	}
}

func (cs *ContactService) Create(contact models.Contacts) (models.Response, error) {
	ctItem, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		return models.Response{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      ctItem,
		TableName: aws.String(cs.TableName),
	}

	_, err = cs.ddb.PutItem(input)
	if err != nil {
		log.Println("Error creating contact in the table", err)
		return models.Response{}, err
	}

	response := models.Response{
		ContactNumber: contact.Id,
	}
	return response, nil
}
