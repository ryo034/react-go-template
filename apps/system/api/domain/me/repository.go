package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/uptrace/bun"
)

type Repository interface {
	Find(ctx context.Context, exec bun.IDB, mID member.ID) (*Me, error)
	FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error)
	LastLogin(ctx context.Context, exec bun.IDB, m *Me) error
	FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error)
	FindProfile(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error)
	FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*Me, error)
	UpdateMember(ctx context.Context, exec bun.IDB, me *Me) error
	UpdateProfile(ctx context.Context, exec bun.IDB, usr *user.User) error
}
