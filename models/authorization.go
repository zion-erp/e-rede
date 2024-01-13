package models

import (
	"time"

	"github.com/zion-erp/e-rede/enums"
)

type Authorization struct {
	Affiliation            uint64                    `json:"affiliation,omitempty"`
	Brand                  *Brand                    `json:"brand,omitempty"`
	Amount                 uint64                    `json:"amount,omitempty"`
	AuthorizationCode      string                    `json:"authorizationCode,omitempty"`
	CardBin                string                    `json:"cardBin,omitempty"`
	CardHolderName         string                    `json:"cardHolderName,omitempty"`
	DateTime               *time.Time                `json:"dateTime,omitempty"`
	DistributorAffiliation uint64                    `json:"distributorAffiliation,omitempty"` // Número de filiação do distribuidor (PV).
	Installments           uint8                     `json:"installments,omitempty"`           // max 12
	Kind                   string                    `json:"kind,omitempty"`
	Last4                  string                    `json:"last4,omitempty"`
	Nsu                    string                    `json:"nsu,omitempty"`
	Origin                 enums.TransactionOrigin   `json:"origin,omitempty"`
	Reference              string                    `json:"reference,omitempty"`
	ReturnCode             string                    `json:"returnCode,omitempty"`
	ReturnMessage          string                    `json:"returnMessage,omitempty"`
	Status                 enums.AuthorizationStatus `json:"status,omitempty"`
	SoftDescriptor         string                    `json:"softDescriptor,omitempty"`
	Subscription           bool                      `json:"subscription,omitempty"`
	Tid                    string                    `json:"tid,omitempty"`
}
