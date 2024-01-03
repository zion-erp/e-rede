package enums

type AuthorizationStatus string

const (
	AuthorizationStatus_Approved AuthorizationStatus = "Approved"
	AuthorizationStatus_Denied   AuthorizationStatus = "Denied"
	AuthorizationStatus_Canceled AuthorizationStatus = "Canceled"
	AuthorizationStatus_Pending  AuthorizationStatus = "Pending"
)
