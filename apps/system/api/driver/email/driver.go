package email

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
)

type Driver interface {
	Send(ctx context.Context, email account.Email, code string) error
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (d *driver) Send(ctx context.Context, email account.Email, code string) error {
	return nil
}
