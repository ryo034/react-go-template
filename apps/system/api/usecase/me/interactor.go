package me

import (
	"context"
	"database/sql"
	"github.com/go-faster/errors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Find(ctx context.Context, aID account.ID) (openapi.APIV1MeGetRes, error)
	UpdateProfile(ctx context.Context, i *UpdateProfileInput) (openapi.APIV1MeProfilePutRes, error)
}

type useCase struct {
	txp      core.TransactionProvider
	dbp      core.Provider
	repo     me.Repository
	wRepo    workspace.Repository
	fbDriver fbDr.Driver
	op       OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, acRepo me.Repository, wRepo workspace.Repository, fbDriver fbDr.Driver, op OutputPort) UseCase {
	return &useCase{txp, dbp, acRepo, wRepo, fbDriver, op}
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
		currentWorkspaceID, err := u.fbDriver.GetCurrentWorkspaceFromCustomClaim(ctx, i.user.AccountID())
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		var current *me.Me
		if currentWorkspaceID == nil {
			current, err = u.repo.FindProfile(pr, p, i.user.AccountID())
			if err != nil {
				return nil, err
			}
		} else {
			mem, err := u.wRepo.FindMember(pr, p, i.user.AccountID(), *currentWorkspaceID)
			if err != nil {
				return nil, err
			}
			current, err = u.repo.Find(pr, p, mem.ID())
			if err != nil {
				return nil, err
			}
		}
		if err = u.repo.UpdateProfile(pr, p, i.user); err != nil {
			return nil, err
		}
		if current.NotJoined() {
			return u.repo.FindBeforeOnboard(pr, p, i.user.AccountID())
		}
		return u.repo.FindLastLogin(pr, p, i.user.AccountID())
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.Find(result.Value(0).(*me.Me)), nil
}
