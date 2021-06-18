package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/controllers"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/models"
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

func (cs *ContactService) Find(id string) (models.Contacts, error) {
	var errMsg string
	result, err := cs.ddb.GetItem(
		&dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(id),
				},
			},
			TableName: aws.String(TABLENAME),
		})
	if err != nil {
		log.Println("Got error calling GetItem ", err.Error())
		return models.Contacts{}, err
	}

	if result.Item == nil {
		errMsg = "'" + id + "' Not Found!"
		log.Println(errMsg)
		return models.Contacts{}, controllers.HandleNotFoundError(errMsg)
	}

	contact := models.Contacts{}
	if err = dynamodbattribute.UnmarshalMap(result.Item, &contact); err != nil {
		errMsg = fmt.Sprintf("Failed to Unmarshall Record! %v", err)
		log.Println(errMsg)
		return models.Contacts{}, controllers.HandleInternalError(errMsg)
	}
	return contact, nil
}
