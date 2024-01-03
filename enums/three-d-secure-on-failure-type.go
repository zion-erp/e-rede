package enums

type ThreeDSecureOnFailureAction string

const (
	ThreeDSecureOnFailureAction_ContinueOnFailure ThreeDSecureOnFailureAction = "continue"
	ThreeDSecureOnFailureAction_DeclineOnFailure  ThreeDSecureOnFailureAction = "decline"
)
