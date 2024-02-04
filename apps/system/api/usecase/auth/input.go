package auth

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

type ByOTPInput struct {
	Email account.Email
}

type VerifyOTPInput struct {
	Email account.Email
	Otp   string
}
