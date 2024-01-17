package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/uptrace/bun"
)

type Repository interface {
	Find(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error)
	Update(ctx context.Context, exec bun.IDB, me *Me) (*Me, error)
}
