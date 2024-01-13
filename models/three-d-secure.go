package models

import "github.com/zion-erp/e-rede/enums"

type ThreeDSecure struct {
	Billing                      *Address                              `json:"billing,omitempty"`
	Cavv                         string                                `json:"cavv,omitempty"`
	ChallengePreference          enums.ThreeDSecureChallengePreference `json:"challengePreference,omitempty"`
	Device                       *Device                               `json:"device,omitempty"`
	DirectoryServerTransactionId string                                `json:"directoryServerTransactionId,omitempty"`
	Eci                          string                                `json:"eci,omitempty"`
	Embedded                     bool                                  `json:"embedded,omitempty"` // mpi(ChallengePreference) == MpiRede;
	IpAddress                    *string                               `json:"ipAddress,omitempty"`
	OnFailure                    enums.ThreeDSecureOnFailureAction     `json:"onFailure,omitempty"` // DeclineOnFailure;
	ReturnCode                   string                                `json:"returnCode,omitempty"`
	ReturnMessage                string                                `json:"returnMessage,omitempty"`
	ThreeDIndicator              uint64                                `json:"threeDIndicator,omitempty"`
	Url                          string                                `json:"url,omitempty"`
	UserAgent                    string                                `json:"userAgent,omitempty"`
	Xid                          string                                `json:"xid,omitempty"`
}
