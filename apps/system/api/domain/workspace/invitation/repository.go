//go:generate gomockhandler -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package invitation

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/uptrace/bun"
)

type Repository interface {
	Find(ctx context.Context, exec bun.IDB, iID ID) (*Invitation, error)
	FindByToken(ctx context.Context, exec bun.IDB, token Token) (*Invitation, error)
	VerifyByToken(ctx context.Context, exec bun.IDB, token Token) error
	FindActiveByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*Invitation, error)
	FindActiveAllByEmail(ctx context.Context, exec bun.IDB, email account.Email) (Invitations, error)
	Accept(ctx context.Context, exec bun.IDB, id ID) error
	Revoke(ctx context.Context, exec bun.IDB, id ID) error
	Resend(ctx context.Context, exec bun.IDB, id ID) (*Invitation, error)
}
