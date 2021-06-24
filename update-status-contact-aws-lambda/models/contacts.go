package models

const PROCESSED_STATUS string = "PROCESSED"

type Contacts struct {
	Id        string `dynamodbav:"id" json:"id"`
	FirstName string `dynamodbav:"firstname" json:"first_name"`
	LastName  string `dynamodbav:"lastname" json:"last_name"`
	Status    string `dynamodbav:"status" json:"status"`
}

func (c *Contacts) Process() {
	c.Status = PROCESSED_STATUS
}

type IContactsRepo interface {
	Find(id string) (Contacts, error)
	ChangeStatus(id string) error
}
