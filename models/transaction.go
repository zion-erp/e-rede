package models

import "github.com/zion-erp/e-rede/enums"

type Transaction struct {
	Additional                     *Additional                  `json:"additional,omitempty"`
	Amount                         uint64                       `json:"amount"`
	Authorization                  *Authorization               `json:"authorization,omitempty"`
	AuthorizationCode              string                       `json:"authorizationCode,omitempty"`
	Brand                          *Brand                       `json:"brand,omitempty"`
	BrandTid                       string                       `json:"brandTid,omitempty"`
	BusinessApplicationIdentifier  string                       `json:"businessApplicationIdentifier,omitempty"`
	CancelId                       string                       `json:"cancelId,omitempty"`
	Capture                        bool                         `json:"capture,omitempty"`
	CardBin                        string                       `json:"cardBin,omitempty"`
	CardHolderName                 string                       `json:"cardHolderName,omitempty"`
	CardNumber                     string                       `json:"cardNumber,omitempty"`
	Cart                           *Cart                        `json:"cart,omitempty"`
	ConsumerBillPaymentService     *ConsumerBillPaymentService  `json:"consumerBillPaymentService,omitempty"`
	CredentialId                   string                       `json:"credentialId,omitempty"`
	DateTime                       string                       `json:"dateTime,omitempty"`
	DistributorAffiliation         uint64                       `json:"distributorAffiliation,omitempty"`
	ExpirationMonth                string                       `json:"expirationMonth,omitempty"`
	ExpirationYear                 string                       `json:"expirationYear,omitempty"`
	Iata                           *Iata                        `json:"iata,omitempty"`
	IndependentSalesOrganizationID int64                        `json:"independentSalesOrganizationID,omitempty"`
	Installments                   uint8                        `json:"installments,omitempty"` // max 12
	Kind                           enums.TransactionKind        `json:"kind,omitempty"`
	Last4                          string                       `json:"last4,omitempty"`
	Nsu                            string                       `json:"nsu,omitempty"`
	Origin                         enums.TransactionOrigin      `json:"origin,omitempty"`
	PaymentFacilitatorID           string                       `json:"paymentFacilitatorID,omitempty"`
	Reference                      string                       `json:"reference,omitempty"`
	RefundDateTime                 string                       `json:"refundDateTime,omitempty"`
	RefundId                       string                       `json:"refundId,omitempty"`
	Refunds                        []*Refund                    `json:"refunds,omitempty"`
	RequestDateTime                string                       `json:"requestDateTime,omitempty"`
	ReturnCode                     string                       `json:"returnCode,omitempty"`
	ReturnMessage                  string                       `json:"returnMessage,omitempty"`
	SecurityAuthentication         *SecurityAuthentication      `json:"securityAuthentication,omitempty"`
	SecurityCode                   string                       `json:"securityCode,omitempty"`
	SoftDescriptor                 string                       `json:"softDescriptor,omitempty"` // maxLength 18
	StorageCard                    enums.TransactionCardStorage `json:"storageCard,omitempty"`
	SubMerchant                    *SubMerchant                 `json:"subMerchant,omitempty"`
	Subscription                   bool                         `json:"subscription,omitempty"`
	ThreeDSecure                   *ThreeDSecure                `json:"threeDSecure,omitempty"`
	Tid                            string                       `json:"tid,omitempty"`
	TokenCryptogram                string                       `json:"tokenCryptogram,omitempty"`
	TransactionCredentials         *TransactionCredentials      `json:"transactionCredentials,omitempty"`
	Urls                           []*Url                       `json:"urls,omitempty"`
	Wallet                         *Wallet                      `json:"wallet,omitempty"`
}
