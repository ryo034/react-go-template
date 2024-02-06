package me

import (
	"context"
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"github.com/uptrace/bun"
)

type UseCase interface {
	Find(ctx context.Context, aID account.ID) (openapi.APIV1MeGetRes, error)
	UpdateProfile(ctx context.Context, i *UpdateProfileInput) (openapi.APIV1MeProfilePutRes, error)
}

type useCase struct {
	txp   core.TransactionProvider
	dbp   core.Provider
	repo  me.Repository
	wRepo workspace.Repository
	op    OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, acRepo me.Repository, wRepo workspace.Repository, op OutputPort) UseCase {
	return &useCase{txp, dbp, acRepo, wRepo, op}
}

func (u *useCase) Find(ctx context.Context, aID account.ID) (openapi.APIV1MeGetRes, error) {
	exec := u.dbp.GetExecutor(ctx, true)
	lastLoginRes, err := u.repo.FindLastLogin(ctx, exec, aID)
	if lastLoginRes != nil {
		return u.op.Find(lastLoginRes), nil
	}
	// If there is no last login information,
	//　it is considered that the user has not joined the workspace.
	m, err := u.repo.FindBeforeOnboard(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	return u.op.Find(m), nil
}

func (u *useCase) UpdateProfile(ctx context.Context, i *UpdateProfileInput) (openapi.APIV1MeProfilePutRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		current, err := u.repo.FindProfile(ctx, p, i.user.AccountID())
		if err != nil {
			return nil, err
		}
		updateEmail := i.user.Email().ToString() != current.Self().Email().ToString()
		if i.user.HasNotName() {
			return nil, fmt.Errorf("name is required")
		}
		updateName := i.user.Name().ToString() != current.Self().Name().ToString()
		if !updateEmail && !updateName {
			return nil, nil
		}
		if err = u.repo.UpdateProfile(pr, p, i.user); err != nil {
			return nil, err
		}
		return u.repo.FindLastLogin(pr, p, i.user.AccountID())
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.Find(result.Value(0).(*me.Me)), nil
}

// updateMember check if the member's information has changed and update it if necessary.
func (u *useCase) updateMember(ctx context.Context, exec bun.IDB, m *me.Me, current *me.Me) error {
	change := false
	if m.Member().DisplayName().ToString() != current.Member().DisplayName().ToString() {
		change = true
	}
	if m.Member().IDNumber().ToString() != current.Member().IDNumber().ToString() {
		change = true
	}
	if !change {
		return nil
	}
	return u.repo.UpdateMember(ctx, exec, m)
}
