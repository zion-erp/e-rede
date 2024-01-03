package models

import "github.com/zion-erp/e-rede/enums"

type Wallet struct {
	ProcessingType          enums.WalletProcessingType `json:"processingType"`
	SenderTaxIdentification string                     `json:"senderTaxIdentification"`
	WalletCode              enums.WalletCode           `json:"walletCode,omitempty"` // Uso exclusivo para  processing type 3 ou 4
	WalletId                uint64                     `json:"walletId,omitempty"`   // Identifica a Wallet originária da transação: * 52810030273 – Apple Pay * 52894351835 – Google Pay  Obrigatório caso processingType=03.
}
