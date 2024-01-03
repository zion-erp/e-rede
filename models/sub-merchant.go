package models

type SubMerchant struct {
	Address       string `json:"address,omitempty"`
	Cep           string `json:"cep,omitempty"`
	City          string `json:"city,omitempty"`
	Cnpj          string `json:"cnpj,omitempty"`
	Country       string `json:"country,omitempty"`
	Mcc           string `json:"mcc,omitempty"`
	State         string `json:"state,omitempty"`
	SubMerchantID string `json:"subMerchantID"`
	TaxIdNumber   string `json:"taxIdNumber"`
}
