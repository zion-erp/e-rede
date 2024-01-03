package enums

type WalletProcessingType string

const (
	WalletProcessingType_Purchase       WalletProcessingType = "01"
	WalletProcessingType_CashIn         WalletProcessingType = "02"
	WalletProcessingType_Elo            WalletProcessingType = "03"
	WalletProcessingType_VisaMastercard WalletProcessingType = "04"
)
