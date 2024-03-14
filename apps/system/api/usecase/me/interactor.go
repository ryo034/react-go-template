package me

import (
	"context"
	"database/sql"

	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"

	"github.com/go-faster/errors"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Find(ctx context.Context, aID account.ID) (openapi.APIV1GetMeRes, error)
	AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.APIV1AcceptInvitationRes, error)
	UpdateName(ctx context.Context, i UpdateNameInput) (openapi.APIV1UpdateProfileRes, error)
	UpdateMemberProfile(ctx context.Context, i UpdateMemberProfileInput) (openapi.APIV1UpdateMeMemberProfileRes, error)
	UpdateProfilePhoto(ctx context.Context, i UpdateProfilePhotoInput) (openapi.APIV1UpdateProfilePhotoRes, error)
	RemoveProfilePhoto(ctx context.Context, i RemoveProfilePhotoInput) (openapi.APIV1RemoveProfilePhotoRes, error)
	LeaveWorkspace(ctx context.Context, i LeaveWorkspaceInput) (openapi.APIV1LeaveWorkspaceRes, error)
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

func (u *useCase) Find(ctx context.Context, aID account.ID) (openapi.APIV1GetMeRes, error) {
	exec := u.dbp.GetExecutor(ctx, false)
	lastLoginRes, err := u.repo.FindLastLogin(ctx, exec, aID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewUnauthenticated("Not logged in")
		}
		return nil, err
	}
	return u.op.Find(lastLoginRes)
}

func (u *useCase) AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.APIV1AcceptInvitationRes, error) {
	exec := u.dbp.GetExecutor(ctx, true)

	invRes, wRes, err := u.wRepo.FindActiveInvitation(ctx, exec, i.InvitationID)
	if err != nil {
		return nil, err
	}

	if err = invRes.ValidateCanAccept(); err != nil {
		return nil, err
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	exec = u.dbp.GetExecutor(pr, false)
	fn := func() (*me.Me, error) {
		m, err := u.repo.FindByEmail(pr, exec, invRes.InviteeEmail())
		if err != nil {
			return nil, err
		}
		mem, err := member.GenerateMember(m.Self())
		if err != nil {
			return nil, err
		}
		mem, err = u.wRepo.AddMember(pr, exec, wRes, mem)
		if err != nil {
			return nil, err
		}
		if err = u.repo.AcceptInvitation(pr, exec, invRes.ID()); err != nil {
			return nil, err
		}
		m, err = u.repo.Find(pr, exec, mem.ID())
		if err != nil {
			return nil, err
		}
		if err = u.repo.RecordLogin(pr, exec, m); err != nil {
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

func (u *useCase) UpdateName(ctx context.Context, i UpdateNameInput) (openapi.APIV1UpdateProfileRes, error) {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	exec := u.dbp.GetExecutor(pr, false)
	fn := func() (*me.Me, error) {
		current, err := u.repo.FindLastLogin(pr, exec, i.AccountID)
		if err != nil {
			return nil, err
		}
		current = current.UpdateName(i.Name)
		return current, u.repo.UpdateName(pr, exec, current.Self())
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.UpdateProfile(result.Value(0).(*me.Me))
}

func (u *useCase) UpdateMemberProfile(ctx context.Context, i UpdateMemberProfileInput) (openapi.APIV1UpdateMeMemberProfileRes, error) {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	exec := u.dbp.GetExecutor(pr, false)
	fn := func() (*me.Me, error) {
		m, err := u.repo.FindLastLogin(pr, exec, i.AccountID)
		if err != nil {
			return nil, err
		}
		if m.NotJoined() {
			return nil, domainErr.NewUnauthenticated("Not joined")
		}
		m = m.UpdateMember(m.Member().UpdateProfile(i.Profile))
		if _, err = u.repo.UpdateMemberProfile(pr, exec, m.Member()); err != nil {
			return nil, err
		}
		return m, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}

	return u.op.UpdateMemberProfile(result.Value(0).(*me.Me))
}

func (u *useCase) UpdateProfilePhoto(ctx context.Context, i UpdateProfilePhotoInput) (openapi.APIV1UpdateProfilePhotoRes, error) {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	exec := u.dbp.GetExecutor(ctx, false)
	fn := func() (*me.Me, error) {
		m, err := u.repo.FindLastLogin(pr, exec, i.AccountID)
		if err != nil {
			return nil, err
		}
		photo := media.NewUploadPhotoToR2(i.File, i.Size, i.Ext)
		m = m.UpdateProfilePhoto(user.NewPhoto(photo.ID(), photo.HostingTo(), nil))
		if err = u.repo.UpdateProfilePhoto(pr, exec, m, photo); err != nil {
			return nil, err
		}
		return m, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.UpdateProfilePhoto(result.Value(0).(*me.Me))
}

func (u *useCase) RemoveProfilePhoto(ctx context.Context, i RemoveProfilePhotoInput) (openapi.APIV1RemoveProfilePhotoRes, error) {
	exec := u.dbp.GetExecutor(ctx, true)
	m, err := u.repo.FindLastLogin(ctx, exec, i.AccountID)
	if err != nil {
		return nil, err
	}
	if m.Self().HasNotPhoto() {
		return nil, domainErr.NewBadRequest("Account has no profile photo")
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	exec = u.dbp.GetExecutor(pr, false)
	fn := func() (*me.Me, error) {
		m = m.RemoveProfilePhoto()
		if err = u.repo.RemoveProfilePhoto(pr, exec, m); err != nil {
			return nil, err
		}
		return m, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.RemoveProfilePhoto(result.Value(0).(*me.Me))
}

func (u *useCase) LeaveWorkspace(ctx context.Context, i LeaveWorkspaceInput) (openapi.APIV1LeaveWorkspaceRes, error) {
	exec := u.dbp.GetExecutor(ctx, true)
	m, err := u.repo.FindLastLogin(ctx, exec, i.AccountID)
	if err != nil {
		return nil, err
	}
	if err = m.ValidateCanLeave(); err != nil {
		return nil, err
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	exec = u.dbp.GetExecutor(pr, false)
	fn := func() error {
		if err = u.wRepo.Leave(pr, exec, m.Member().ID(), m.Member().ID()); err != nil {
			return err
		}
		return u.repo.ClearMe(pr)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.LeaveWorkspace()
}
