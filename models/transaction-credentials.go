package models

import "github.com/zion-erp/e-rede/enums"

type TransactionCredentials struct {
	CredentialId enums.TransactionCredentialType `json:"credentialId,omitempty"`
}
