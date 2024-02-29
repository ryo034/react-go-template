package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
)

type ByOTPInput struct {
	Email account.Email
}

type ByOAuthInput struct {
	AccountID       account.ID
	AuthProviderUID provider.UID
}

type VerifyOTPInput struct {
	Email account.Email
	Otp   string
}

type ProcessInvitationEmailInput struct {
	Token invitation.Token
	Email account.Email
}

type ProcessInvitationOAuthInput struct {
	Token invitation.Token
}

type AcceptInvitationInput struct {
	AccountID    account.ID
	InvitationID invitation.ID
}

type InvitationByTokenInput struct {
	Token invitation.Token
}
