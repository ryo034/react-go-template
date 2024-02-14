package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
)

type ByOTPInput struct {
	Email account.Email
}

type VerifyOTPInput struct {
	Email account.Email
	Otp   string
}

type ProcessInvitationInput struct {
	Token invitation.Token
	Email account.Email
}

type AcceptInvitationInput struct {
	AccountID    account.ID
	InvitationID invitation.ID
}
