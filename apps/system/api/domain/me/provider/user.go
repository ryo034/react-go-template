package provider

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

type UserInfo struct {
	DisplayName string
	Email       string
	PhoneNumber string
	PhotoURL    string
	ProviderID  string
	UID         account.ID
}

type UserMetadata struct {
	CreationTimestamp    int64
	LastLogInTimestamp   int64
	LastRefreshTimestamp int64
}

type PhoneMultiFactorInfo struct {
	PhoneNumber string
}

type MultiFactorInfo struct {
	UID                 account.ID
	DisplayName         string
	EnrollmentTimestamp int64
	FactorID            string
	Phone               *PhoneMultiFactorInfo
}

type MultiFactorSettings struct {
	EnrolledFactors []*MultiFactorInfo
}

type User struct {
	*UserInfo
	CustomClaims           map[string]interface{}
	Disabled               bool
	EmailVerified          bool
	ProviderUserInfo       []*UserInfo
	TokensValidAfterMillis int64 // milliseconds since epoch.
	UserMetadata           *UserMetadata
	TenantID               string
	MultiFactor            *MultiFactorSettings
}

func NewUser(u *UserInfo) *User {
	return &User{UserInfo: u}
}
