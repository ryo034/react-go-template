package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
)

type UseCase interface {
	Login(ctx context.Context, aID account.ID) (*me.Me, error)
	Find(ctx context.Context, aID account.ID, isRead bool) (*me.Me, error)
}

type useCase struct {
	isLocal bool
	repo    me.Repository
}

func NewUseCase(isLocal bool, acRepo me.Repository) UseCase {
	return &useCase{isLocal, acRepo}
}

func (u *useCase) Login(ctx context.Context, aID account.ID) (*me.Me, error) {
	return u.repo.Find(ctx, aID)
}

func (u *useCase) Find(ctx context.Context, aID account.ID, isRead bool) (*me.Me, error) {
	return u.repo.Find(ctx, aID)
}
