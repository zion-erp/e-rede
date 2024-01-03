package models

import "github.com/zion-erp/e-rede/enums"

// type AddressUsage uint16

// const (
// 	AddressUsage_Billing  AddressUsage = 1
// 	AddressUsage_Shipping AddressUsage = 2
// 	AddressUsage_Both     AddressUsage = 3
// )

type Address struct {
	Address       string            `json:"address,omitempty"`
	AddresseeName string            `json:"addresseeName,omitempty"`
	City          string            `json:"city,omitempty"`
	Country       string            `json:"country,omitempty"`
	EmailAdress   string            `json:"emailAdress,omitempty"`
	Number        string            `json:"number,omitempty"`
	PhoneNumber   string            `json:"phoneNumber,omitempty"`
	State         string            `json:"state,omitempty"`
	Type          enums.AddressType `json:"type,omitempty"`
	ZipCode       string            `json:"zipCode,omitempty"`
	PostalCode    string            `json:"postalcode,omitempty"`
}
