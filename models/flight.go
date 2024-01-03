package models

type Flight struct {
	Date      string       `json:"date,omitempty"`
	Ip        string       `json:"ip,omitempty"`
	Number    string       `json:"number,omitempty"`
	Passenger []*Passenger `json:"passenger,omitempty"`
	To        string       `json:"to,omitempty"`
}
