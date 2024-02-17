package invitation

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Repository interface {
	Find(ctx context.Context, exec bun.IDB, iID ID) (*Invitation, error)
	FindByToken(ctx context.Context, exec bun.IDB, token uuid.UUID) (*Invitation, error)
	VerifyByToken(ctx context.Context, exec bun.IDB, token uuid.UUID) error
}
