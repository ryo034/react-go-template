//go:generate gomockhandler -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	Find(ctx context.Context, exec boil.ContextExecutor, aID account.ID) (*Me, error)
	VerifyEmail(ctx context.Context, aID account.ID) error
	SaveFromTemporary(ctx context.Context, aID account.ID, firstName account.FirstName, lastName account.LastName) (*Me, error)
	Update(ctx context.Context, me *Me) error
	EmailVerified(ctx context.Context, aID account.ID) (bool, string, error)
	UpdateEmail(ctx context.Context, aID account.ID, em account.Email) (*Me, error)
	UpdateName(ctx context.Context, aID account.ID, fin account.FirstName, ln account.LastName) (*Me, error)
	UpdatePhoneNumber(ctx context.Context, aID account.ID, ph phone.Number) (*Me, error)
}
