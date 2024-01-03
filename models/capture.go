package models

type Capture struct {
	Amount   uint64 `json:"amount,omitempty"`
	DateTime string `json:"dateTime,omitempty"`
	Nsu      string `json:"nsu,omitempty"`
}
