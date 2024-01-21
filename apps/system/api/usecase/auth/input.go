package auth

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

type ByTOTPInput struct {
	Email account.Email
}

type VerifyTOTPInput struct {
	Email account.Email
	Otp   string
}
