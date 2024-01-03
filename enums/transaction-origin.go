package enums

type TransactionOrigin uint8

const (
	TransactionOrigin_OriginERede        TransactionOrigin = 1
	TransactionOrigin_OriginVisaCheckout TransactionOrigin = 4
	TransactionOrigin_OriginMasterpass   TransactionOrigin = 6
)
