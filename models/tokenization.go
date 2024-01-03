package models

import "github.com/zion-erp/e-rede/enums"

type CreateTokenization struct {
	CardholderName  string                       `json:"cardholderName,omitempty"`
	CardNumber      string                       `json:"cardNumber"`
	Email           string                       `json:"email"`
	ExpirationMonth string                       `json:"expirationMonth"`
	ExpirationYear  string                       `json:"expirationYear"`
	SecurityCode    string                       `json:"securityCode,omitempty"`
	StorageCard     enums.TransactionCardStorage `json:"storageCard"`
}

type CreateTokenizationResponse struct {
	ReturnCode     string  `json:"returnCode,omitempty"`
	ReturnMessage  string  `json:"returnMessage,omitempty"`
	TokenizationId string  `json:"tokenizationId,omitempty"`
	Links          []*Link `json:"links,omitempty"`
}

type Token struct {
	Code           uint64 `json:"code,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
}

type QueryTokenizationResponse struct {
	ReturnCode         string `json:"returnCode,omitempty"`
	ReturnMessage      string `json:"returnMessage,omitempty"`
	Affiliation        uint64 `json:"affiliation,omitempty"`
	TokenizationId     string `json:"tokenizationId,omitempty"`
	TokenizationStatus string `json:"tokenizationStatus,omitempty"`
	Brand              *Brand `json:"brand,omitempty"`
	LastModifiedDate   string `json:"lastModifiedDate,omitempty"`
	Last4              string `json:"last4,omitempty"`
	Token              *Token `json:"token,omitempty"`
}

type CryptogramInfo struct {
	TokenCryptogram string `json:"tokenCryptogram,omitempty"`
	Eci             string `json:"eci,omitempty"`
	ExpirationDate  string `json:"expirationDate,omitempty"`
}

type QueryTokenizationCryptogramResponse struct {
	ReturnCode     string          `json:"returnCode,omitempty"`
	ReturnMessage  string          `json:"returnMessage,omitempty"`
	TokenizationId string          `json:"tokenizationId,omitempty"`
	CryptogramInfo *CryptogramInfo `json:"cryptogramInfo,omitempty"`
}
