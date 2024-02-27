package me

import (
	"context"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Find(ctx context.Context, aID account.ID) (openapi.APIV1MeGetRes, error)
	AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.AcceptInvitationRes, error)
	UpdateProfile(ctx context.Context, i UpdateProfileInput) (openapi.APIV1MeProfilePutRes, error)
	UpdateMemberProfile(ctx context.Context, i UpdateMemberProfileInput) (openapi.APIV1MeMemberProfilePutRes, error)
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
	exec := u.dbp.GetExecutor(ctx, false)
	lastLoginRes, err := u.repo.FindLastLogin(ctx, exec, aID)
	if lastLoginRes != nil {
		return u.op.Find(lastLoginRes)
	}
	// If there is no last login information,
	// it is considered that the user has not joined the workspace.
	m, err := u.repo.FindBeforeOnboard(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	return u.op.Find(m)
}

func (u *useCase) AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.AcceptInvitationRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}

	invRes, wRes, err := u.wRepo.FindActiveInvitation(ctx, p, i.InvitationID)
	if err != nil {
		return nil, err
	}

	if err = invRes.ValidateCanAccept(); err != nil {
		return nil, err
	}

	fn := func() (*me.Me, error) {
		m, err := u.repo.FindByEmail(pr, p, invRes.InviteeEmail())
		if err != nil {
			return nil, err
		}
		id, err := member.GenerateID()
		if err != nil {
			return nil, err
		}
		mem := member.NewMember(id, m.Self(), member.NewEmptyProfile())
		if err != nil {
			return nil, err
		}
		mem, err = u.wRepo.AddMember(pr, p, wRes, mem)
		if err != nil {
			return nil, err
		}
		if err = u.repo.AcceptInvitation(pr, p, invRes.ID()); err != nil {
			return nil, err
		}
		m, err = u.repo.Find(pr, p, mem.ID())
		if err != nil {
			return nil, err
		}
		if err = u.repo.LastLogin(pr, p, m); err != nil {
			return nil, err
		}
		return m, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.AcceptInvitation(result.Value(0).(*me.Me))
}

func (u *useCase) UpdateProfile(ctx context.Context, i UpdateProfileInput) (openapi.APIV1MeProfilePutRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		var currentWorkspaceID *workspace.ID
		currentWorkspaceID, err = u.fbDriver.GetCurrentWorkspaceFromCustomClaim(ctx, i.AccountID)
		if err != nil {
			return nil, err
		}
		var current *me.Me
		if currentWorkspaceID == nil {
			current, err = u.repo.FindProfile(pr, p, i.AccountID)
			if err != nil {
				return nil, err
			}
		} else {
			mem, err := u.wRepo.FindMember(pr, p, i.AccountID, *currentWorkspaceID)
			if err != nil {
				return nil, err
			}
			current, err = u.repo.Find(pr, p, mem.ID())
			if err != nil {
				return nil, err
			}
		}

		usr := current.Self().UpdateName(i.Name)

		if err = u.repo.UpdateProfile(pr, p, usr); err != nil {
			return nil, err
		}
		if current.NotJoined() {
			return u.repo.FindBeforeOnboard(pr, p, i.AccountID)
		}
		return u.repo.FindLastLogin(pr, p, i.AccountID)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.UpdateProfile(result.Value(0).(*me.Me))
}

func (u *useCase) UpdateMemberProfile(ctx context.Context, i UpdateMemberProfileInput) (openapi.APIV1MeMemberProfilePutRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		var currentWorkspaceID *workspace.ID
		currentWorkspaceID, err = u.fbDriver.GetCurrentWorkspaceFromCustomClaim(ctx, i.AccountID)
		if err != nil {
			return nil, err
		}
		if currentWorkspaceID == nil {
			return nil, domainErr.NewUnauthenticated("No current workspace")
		}
		mem, err := u.wRepo.FindMember(pr, p, i.AccountID, *currentWorkspaceID)
		if err != nil {
			return nil, err
		}
		if mem, err = u.repo.UpdateMemberProfile(pr, p, mem.UpdateProfile(i.Profile)); err != nil {
			return nil, err
		}
		return u.repo.FindLastLogin(pr, p, i.AccountID)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.UpdateMemberProfile(result.Value(0).(*me.Me))
}
