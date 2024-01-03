package enums

type TransactionKind string

const (
	TransactionKind_Credit TransactionKind = "credit"
	TransactionKind_Debit  TransactionKind = "debit"
)
