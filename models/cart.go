package models

type Cart struct {
	Addresses   []*Address   `json:"addresses,omitempty"`
	Billing     *Address     `json:"billing,omitempty"`
	Customer    *Customer    `json:"customer,omitempty"`
	Environment *Environment `json:"environment,omitempty"`
	Iata        *Iata        `json:"iata,omitempty"`
	Items       []*Item      `json:"items,omitempty"`
	Shipping    *Address     `json:"shipping,omitempty"`
}
