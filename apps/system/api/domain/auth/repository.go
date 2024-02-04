package auth

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/uptrace/bun"
)

type Repository interface {
	GenTOTP(ctx context.Context, email account.Email) (string, error)
	VerifyOTP(ctx context.Context, email account.Email, code string) (bool, error)
	Create(ctx context.Context, exec bun.IDB, aID account.ID, email account.Email) (*user.User, error)
	Find(ctx context.Context, exec bun.IDB, email account.Email) (*user.User, error)
}
