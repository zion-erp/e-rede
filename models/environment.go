package models

type Environment struct {
	Endpoint  string `json:"endpoint,omitempty"`
	Ip        string `json:"ip,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
}
