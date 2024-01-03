package models

type Brand struct {
	Name          string `json:"name,omitempty"`
	ReturnCode    string `json:"returnCode,omitempty"`
	ReturnMessage string `json:"returnMessage,omitempty"`
	Message       string `json:"message,omitempty"`
}
