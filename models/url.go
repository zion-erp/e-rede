package models

import "github.com/zion-erp/e-rede/enums"

type Url struct {
	Kind enums.UrlKind `json:"kind,omitempty"`
	Type string        `json:"type,omitempty"`
	Url  string        `json:"url,omitempty"`
}
