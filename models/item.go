package models

import "github.com/zion-erp/e-rede/enums"

type Item struct {
	Amount       uint64         `json:"amount,omitempty"`
	Description  string         `json:"description,omitempty"`
	Freight      uint64         `json:"freight,omitempty"`
	Id           string         `json:"id,omitempty"`
	Quantity     uint64         `json:"quantity,omitempty"`
	ShippingType string         `json:"shippingType,omitempty"`
	Type         enums.ItemType `json:"type,omitempty"`
}
