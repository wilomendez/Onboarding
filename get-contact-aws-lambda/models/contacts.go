package models

type Contacts struct {
	Id        string `dynamodbav:"id"`
	FirstName string `dynamodbav:"firstname"`
	LastName  string `dynamodbav:"lastname"`
	Status    string `dynamodbav:"status"`
}

type IContactsRepo interface {
	Find(id string) (Contacts, error)
}
