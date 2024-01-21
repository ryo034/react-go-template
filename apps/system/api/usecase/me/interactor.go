package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Login(ctx context.Context, aID account.ID, wID workspace.ID) (openapi.LoginRes, error)
	Find(ctx context.Context, aID account.ID, wID workspace.ID) (openapi.MeGetRes, error)
	Update(ctx context.Context, me *me.Me) (*openapi.Me, error)
}

type useCase struct {
	txp  core.TransactionProvider
	dbp  core.Provider
	repo me.Repository
	op   OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, acRepo me.Repository, op OutputPort) UseCase {
	return &useCase{txp, dbp, acRepo, op}
}

func (u *useCase) Login(ctx context.Context, aID account.ID, wID workspace.ID) (openapi.LoginRes, error) {
	res, err := u.repo.Find(ctx, u.dbp.GetExecutor(ctx, true), aID, wID)
	if err != nil {
		return nil, err
	}
	return u.op.Find(res), nil
}

func (u *useCase) Find(ctx context.Context, aID account.ID, wID workspace.ID) (openapi.MeGetRes, error) {
	res, err := u.repo.Find(ctx, u.dbp.GetExecutor(ctx, true), aID, wID)
	if err != nil {
		return nil, err
	}
	return u.op.Find(res), nil
}

func (u *useCase) Update(ctx context.Context, m *me.Me) (*openapi.Me, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		return u.repo.Update(ctx, p, m)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.Find(result.Value(0).(*me.Me)), nil
}
