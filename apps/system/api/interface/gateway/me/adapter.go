package me

import (
	"firebase.google.com/go/v4/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/models"
)

type Adapter interface {
	Adapt(s *models.SystemAccount, fu *auth.UserRecord) (*me.Me, error)
	AdaptFirebaseUser(fu *auth.UserRecord, name account.Name, phoneNumber *phone.Number) (*me.Me, error)
}
