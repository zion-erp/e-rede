package models

import (
	"time"

	"github.com/zion-erp/e-rede/enums"
)

type RefundStatusHistory struct {
	Status   enums.RefundStatus `json:"status,omitempty"`
	DateTime *time.Time         `json:"dateTime,omitempty"`
}
