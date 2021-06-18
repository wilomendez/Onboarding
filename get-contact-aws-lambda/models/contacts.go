package models

type Contacts struct {
	Id        string `dynamodbav:"id" json:"id"`
	FirstName string `dynamodbav:"firstname" json:"first_name"`
	LastName  string `dynamodbav:"lastname" json:"last_name"`
	Status    string `dynamodbav:"status" json:"status"`
}

type IContactsRepo interface {
	Find(id string) (Contacts, error)
}
