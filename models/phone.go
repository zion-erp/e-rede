package models

import "github.com/zion-erp/e-rede/enums"

type Phone struct {
	Ddd    string          `json:"ddd,omitempty"`
	Number string          `json:"number,omitempty"`
	Type   enums.PhoneType `json:"type,omitempty"`
}
