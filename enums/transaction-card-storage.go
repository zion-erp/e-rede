package enums

type TransactionCardStorage string

const (
	TransactionCardStorage_CredentialNotStored            TransactionCardStorage = "0"
	TransactionCardStorage_CredentialStoredFirstTimeUsage TransactionCardStorage = "1"
	TransactionCardStorage_CredentialStored               TransactionCardStorage = "2"
)
