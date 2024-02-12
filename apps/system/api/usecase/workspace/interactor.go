package workspace

import (
	"context"
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Create(ctx context.Context, i *CreateInput) (openapi.APIV1WorkspacesPostRes, error)
	FindAllMembers(ctx context.Context, i *FindAllMembersInput) (openapi.APIV1MembersGetRes, error)
	InviteMembers(ctx context.Context, i *InviteMembersInput) (openapi.InviteMultipleUsersToWorkspaceRes, error)
	VerifyInvitationToken(ctx context.Context, i *VerifyInvitationTokenInput) (openapi.VerifyInvitationRes, error)
}

type useCase struct {
	txp         core.TransactionProvider
	dbp         core.Provider
	repo        workspace.Repository
	meRepo      me.Repository
	fbDriver    fbDr.Driver
	emailDriver email.Driver
	op          OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, repo workspace.Repository, meRepo me.Repository, fbDriver fbDr.Driver, emailDriver email.Driver, op OutputPort) UseCase {
	return &useCase{txp, dbp, repo, meRepo, fbDriver, emailDriver, op}
}

func (u *useCase) Create(ctx context.Context, i *CreateInput) (openapi.APIV1WorkspacesPostRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*workspace.Workspace, error) {
		w := i.Workspace
		meRes, err := u.meRepo.FindBeforeOnboard(pr, p, i.AccountID)
		if err != nil {
			return nil, err
		}
		wres, err := u.repo.Create(pr, p, w)
		if err != nil {
			return nil, err
		}
		meID, err := member.GenerateID()
		if err != nil {
			return nil, err
		}
		dn, err := member.NewDisplayName(meRes.Self().Name().ToString())
		if err != nil {
			return nil, err
		}

		m := member.NewMember(meID, meRes.Self(), dn, nil)
		memRes, err := u.repo.AddMember(pr, p, wres, m)
		if err != nil {
			return nil, err
		}

		meRes, err = u.meRepo.Find(pr, p, memRes.ID())
		if err != nil {
			return nil, err
		}
		if err = u.meRepo.LastLogin(pr, p, meRes); err != nil {
			return nil, err
		}
		return wres, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	res := result.Value(0).(*workspace.Workspace)
	return u.op.Create(res), nil
}

func (u *useCase) FindAllMembers(ctx context.Context, i *FindAllMembersInput) (openapi.APIV1MembersGetRes, error) {
	exec := u.dbp.GetExecutor(ctx, true)
	currentWorkspaceID, err := u.fbDriver.GetCurrentWorkspaceFromCustomClaim(ctx, i.accountID)
	if err != nil {
		return nil, err
	}
	ms, err := u.repo.FindAllMembers(ctx, exec, *currentWorkspaceID)
	if err != nil {
		return nil, err
	}
	return u.op.FindAllMembers(ms), nil
}

func (u *useCase) inviteMember(ctx context.Context, wID workspace.ID, wName workspace.Name, invitedBy member.InvitedBy, ivm *member.InvitedMember) error {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return err
	}
	fn := func() error {
		if err = u.repo.InviteMember(pr, p, wID, invitedBy, ivm); err != nil {
			return err
		}
		return u.emailDriver.SendInvite(pr, invitedBy, ivm)
	}
	result := pr.Transactional(fn)()
	return result.Error()
}

func (u *useCase) InviteMembers(ctx context.Context, i *InviteMembersInput) (openapi.InviteMultipleUsersToWorkspaceRes, error) {
	// Exclude already registered members
	exec := u.dbp.GetExecutor(ctx, true)
	currentWorkspaceID, err := u.fbDriver.GetCurrentWorkspaceFromCustomClaim(ctx, i.AccountID)
	if err != nil {
		return nil, err
	}
	if currentWorkspaceID == nil {
		//TODO: error handling
		return nil, fmt.Errorf("current workspace is not found")
	}
	members, err := u.repo.FindAllMembers(ctx, exec, *currentWorkspaceID)
	if err != nil {
		return nil, err
	}

	targetMembers := make([]*member.InvitedMember, 0)
	alreadyRegisteredMembers := make([]*member.InvitedMember, 0)
	for _, im := range i.InvitedMembers.AsSlice() {
		if members.Exist(im.Email()) {
			alreadyRegisteredMembers = append(alreadyRegisteredMembers, im)
		} else {
			targetMembers = append(targetMembers, im)
		}
	}

	meRes, err := u.meRepo.FindLastLogin(ctx, exec, i.AccountID)
	if err != nil {
		return nil, err
	}
	wd := meRes.Workspace().Detail()

	failedMembers := make([]*member.InvitedMember, 0)
	invitedBy := member.NewInvitedBy(meRes.Member().ID(), meRes.Member().DisplayName())
	for _, im := range targetMembers {
		if err = u.inviteMember(
			ctx,
			meRes.Workspace().ID(),
			wd.Name(),
			invitedBy,
			im,
		); err != nil {
			failedMembers = append(failedMembers, im)
		}
	}

	return u.op.InviteMembers(
		i.InvitedMembers,
		member.NewInvitedMembers(alreadyRegisteredMembers),
		member.NewInvitedMembers(failedMembers),
	), nil
}

func (u *useCase) VerifyInvitationToken(ctx context.Context, i *VerifyInvitationTokenInput) (openapi.VerifyInvitationRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*member.InvitedMember, error) {
		return u.repo.VerifyInvitedMember(pr, p, i.Token)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.VerifyInvitationToken(result.Value(0).(*member.InvitedMember)), nil
}
