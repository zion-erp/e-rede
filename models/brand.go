package models

import "github.com/zion-erp/e-rede/enums"

type Brand struct {
	AuthorizationCode  string          `json:"authorizationCode,omitempty"`
	BrandTid           string          `json:"brandTid,omitempty"`
	MerchantAdviceCode string          `json:"merchantAdviceCode,omitempty"`
	Message            string          `json:"message,omitempty"`
	Name               enums.BrandName `json:"name,omitempty"`
	ReturnCode         string          `json:"returnCode,omitempty"`
	ReturnMessage      string          `json:"returnMessage,omitempty"`
}
