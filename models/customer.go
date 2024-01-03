package models

import "github.com/zion-erp/e-rede/enums"

type Customer struct {
	Cpf       string               `json:"cpf,omitempty"`
	Documents []*Document          `json:"documents,omitempty"`
	Email     string               `json:"email,omitempty"`
	Gender    enums.CustomerGender `json:"gender,omitempty"`
	Name      string               `json:"name,omitempty"`
	Phone     *Phone               `json:"phone,omitempty"`
}
