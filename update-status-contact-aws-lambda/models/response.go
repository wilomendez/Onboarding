package models

type Data struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
}

type Message struct {
	EventID     string `json:"event_id"`
	EventName   string `json:"event_name"`
	EventSource string `json:"event_source"`
	Item        Data   `json:"item"`
}
