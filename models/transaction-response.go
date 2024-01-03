package models

import (
	"time"

	"github.com/zion-erp/e-rede/enums"
)

type TransactionResponse struct {
	Amount                 uint64                       `json:"amount,omitempty"`
	Authorization          *Authorization               `json:"authorization,omitempty"`
	AuthorizationCode      string                       `json:"authorizationCode,omitempty"`
	Brand                  *Brand                       `json:"brand,omitempty"`
	BrandTid               string                       `json:"brandTid,omitempty"`
	CancelId               string                       `json:"cancelId,omitempty"`
	Capture                *Capture                     `json:"capture,omitempty"`
	CardBin                string                       `json:"cardBin,omitempty"`
	CardHolderName         string                       `json:"cardHolderName,omitempty"`
	CardNumber             string                       `json:"cardNumber,omitempty"`
	Cart                   Cart                         `json:"cart,omitempty"`
	DateTime               *time.Time                   `json:"dateTime,omitempty"`
	DistributorAffiliation uint64                       `json:"distributorAffiliation,omitempty"`
	ExpirationMonth        string                       `json:"expirationMonth,omitempty"`
	ExpirationYear         string                       `json:"expirationYear,omitempty"`
	Iata                   *Iata                        `json:"iata,omitempty"`
	Installments           uint8                        `json:"installments,omitempty"` // max 12
	Kind                   string                       `json:"kind,omitempty"`
	Last4                  string                       `json:"last4,omitempty"`
	Links                  []*Link                      `json:"links,omitempty"`
	Nsu                    string                       `json:"nsu,omitempty"`
	Origin                 string                       `json:"origin,omitempty"`
	Reference              string                       `json:"reference,omitempty"`
	RefundDateTime         *time.Time                   `json:"refundDateTime,omitempty"`
	RefundId               string                       `json:"refundId,omitempty"`
	Refunds                []*Refund                    `json:"refunds,omitempty"`
	RequestDateTime        *time.Time                   `json:"requestDateTime,omitempty"`
	ReturnCode             string                       `json:"returnCode,omitempty"`
	ReturnMessage          string                       `json:"returnMessage,omitempty"`
	SecurityCode           string                       `json:"securityCode,omitempty"`
	SoftDescriptor         string                       `json:"softDescriptor,omitempty"`
	StorageCard            enums.TransactionCardStorage `json:"storageCard,omitempty"`
	Subscription           bool                         `json:"subscription,omitempty"`
	ThreeDSecure           *ThreeDSecure                `json:"threeDSecure,omitempty"`
	Tid                    string                       `json:"tid,omitempty"`
	Urls                   []*Url                       `json:"urls,omitempty"`
}
