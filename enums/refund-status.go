package enums

type RefundStatus string

const (
	RefundStatus_Done       RefundStatus = "Done"
	RefundStatus_Denied     RefundStatus = "Denied"
	RefundStatus_Processing RefundStatus = "Processing"
)
