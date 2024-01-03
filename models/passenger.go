package models

type Passenger struct {
	Email  string `json:"email,omitempty"`
	Name   string `json:"name,omitempty"`
	Phone  *Phone `json:"phone,omitempty"`
	Ticket string `json:"ticket,omitempty"`
}
