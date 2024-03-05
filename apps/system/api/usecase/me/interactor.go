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
	Find(ctx context.Context, aID account.ID) (openapi.APIV1MeGetRes, error)
	AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.AcceptInvitationRes, error)
	UpdateName(ctx context.Context, i UpdateNameInput) (openapi.APIV1MeProfilePutRes, error)
	UpdateMemberProfile(ctx context.Context, i UpdateMemberProfileInput) (openapi.APIV1MeMemberProfilePutRes, error)
	UpdateProfilePhoto(ctx context.Context, i UpdateProfilePhotoInput) (openapi.APIV1MeProfilePhotoPutRes, error)
	RemoveProfilePhoto(ctx context.Context, i RemoveProfilePhotoInput) (openapi.APIV1MeProfilePhotoDeleteRes, error)
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
	exec := u.dbp.GetExecutor(ctx, false)
	lastLoginRes, err := u.repo.FindLastLogin(ctx, exec, aID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return u.op.Find(lastLoginRes)
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
		mem, err := member.GenerateMember(m.Self())
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
		if err = u.repo.RecordLogin(pr, p, m); err != nil {
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

func (u *useCase) UpdateName(ctx context.Context, i UpdateNameInput) (openapi.APIV1MeProfilePutRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		current, err := u.repo.FindLastLogin(pr, p, i.AccountID)
		if err != nil {
			return nil, err
		}
		current = current.UpdateName(i.Name)
		return current, u.repo.UpdateName(pr, p, current.Self())
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
		m, err := u.repo.FindLastLogin(pr, p, i.AccountID)
		if err != nil {
			return nil, err
		}
		if m.NotJoined() {
			return nil, domainErr.NewUnauthenticated("Not joined")
		}
		m = m.UpdateMember(m.Member().UpdateProfile(i.Profile))
		if _, err = u.repo.UpdateMemberProfile(pr, p, m.Member()); err != nil {
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

func (u *useCase) UpdateProfilePhoto(ctx context.Context, i UpdateProfilePhotoInput) (openapi.APIV1MeProfilePhotoPutRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		m, err := u.repo.FindLastLogin(pr, p, i.AccountID)
		if err != nil {
			return nil, err
		}
		photo := media.NewUploadPhotoToR2(i.Photo)
		m = m.UpdateProfilePhoto(user.NewPhoto(photo.ID(), photo.HostingTo(), nil))
		if err = u.repo.UpdateProfilePhoto(pr, p, m, photo); err != nil {
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

func (u *useCase) RemoveProfilePhoto(ctx context.Context, i RemoveProfilePhotoInput) (openapi.APIV1MeProfilePhotoDeleteRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		m, err := u.repo.FindLastLogin(pr, p, i.AccountID)
		if err != nil {
			return nil, err
		}
		m = m.RemoveProfilePhoto()
		if err = u.repo.RemoveProfilePhoto(pr, p, m); err != nil {
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
