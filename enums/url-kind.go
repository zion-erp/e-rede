package enums

type UrlKind string

const (
	UrlKind_Callback             UrlKind = "callback"
	UrlKind_ThreeDSecureFailure  UrlKind = "threeDSecureFailure"
	UrlKind_ThreeDSecureSuccess  UrlKind = "threeDSecureSuccess"
	UrlKind_ThreeDSecureCallback UrlKind = "threeDSecureCallback"
)
