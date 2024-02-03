package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/uptrace/bun"
)

type Repository interface {
	Find(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*Me, error)
	FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error)
	LastLogin(ctx context.Context, exec bun.IDB, m *Me) error
	FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error)
	Update(ctx context.Context, exec bun.IDB, me *Me) (*Me, error)
	UpdateName(ctx context.Context, exec bun.IDB, aID account.ID, name account.Name) (*Me, error)
}
