package models

import (
	"time"

	"github.com/zion-erp/e-rede/enums"
)

type Refund struct {
	Amount         uint64                 `json:"amount,omitempty"`
	CancelId       string                 `json:"cancelId,omitempty"`
	Links          []*Link                `json:"links,omitempty"`
	RefundDateTime *time.Time             `json:"refundDateTime,omitempty"`
	RefundId       string                 `json:"refundId,omitempty"`
	Status         enums.RefundStatus     `json:"status,omitempty"`
	StatusHistory  *[]RefundStatusHistory `json:"statusHistory,omitempty"`
	ReturnCode     string                 `json:"returnCode,omitempty"`
	ReturnMessage  string                 `json:"returnMessage,omitempty"`
	Tid            string                 `json:"tid,omitempty"`
}
