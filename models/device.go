package models

type Device struct {
	ColorDepth     uint16 `json:"colorDepth,omitempty"`
	DeviceType3ds  string `json:"deviceType3ds,omitempty"`
	JavaEnabled    bool   `json:"javaEnabled,omitempty"`
	Language       string `json:"language,omitempty"` //"BR"
	ScreenHeight   uint16 `json:"screenHeight,omitempty"`
	ScreenWidth    uint16 `json:"screenWidth,omitempty"`
	TimeZoneOffset uint16 `json:"timeZoneOffset,omitempty"` //3
}
