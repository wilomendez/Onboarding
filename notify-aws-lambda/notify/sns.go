package notify

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/wilomendez/Onboarding/notify-aws-lambda/models"
)

const (
	REGION    string = "us-east-1"
	TOPIC_ARN string = "arn:aws:sns:us-east-1:161142984839:ContactsTopic_WM"
)

type Notify struct {
	topic  string
	notify *sns.SNS
}

func New() models.IMessageRepo {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(REGION)}))
	return &Notify{
		topic:  TOPIC_ARN,
		notify: sns.New(sess),
	}
}

func (n *Notify) Push(message, attrs string) error {
	result, err := n.notify.Publish(&sns.PublishInput{
		Message:  &message,
		TopicArn: &n.topic,
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(attrs),
			},
		},
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println(result.MessageId)
	return nil
}
