package models

type Iata struct {
	Code         uint64    `json:"code,omitempty"`
	DepartureTax uint64    `json:"departureTax,omitempty"`
	Flight       []*Flight `json:"flight,omitempty"`
}
