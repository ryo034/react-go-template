//go:generate gomockhandler -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package core

import (
	"context"

	"github.com/uptrace/bun"
)

type provider struct {
	primary *bun.DB
	replica *bun.DB
}

func NewDatabaseProvider(p *bun.DB, r *bun.DB) Provider {
	return &provider{p, r}
}

func (i *provider) GetExecutor(ctx context.Context, isRead bool) bun.IDB {
	if isRead {
		return i.replica
	}
	if result, ok := ctx.(*transactionResult); ok {
		return result
	}
	return i.primary
}

type Provider interface {
	GetExecutor(ctx context.Context, isRead bool) bun.IDB
}
