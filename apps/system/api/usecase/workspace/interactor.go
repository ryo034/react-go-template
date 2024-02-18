package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Create(ctx context.Context, i CreateInput) (openapi.APIV1WorkspacesPostRes, error)
	FindAllMembers(ctx context.Context, i FindAllMembersInput) (openapi.APIV1MembersGetRes, error)
	InviteMembers(ctx context.Context, i InviteMembersInput) (openapi.InviteMultipleUsersToWorkspaceRes, error)
	VerifyInvitationToken(ctx context.Context, i VerifyInvitationTokenInput) (openapi.VerifyInvitationRes, error)
}

type useCase struct {
	txp         core.TransactionProvider
	dbp         core.Provider
	repo        workspace.Repository
	meRepo      me.Repository
	invRepo     invitation.Repository
	fbDriver    fbDr.Driver
	emailDriver email.Driver
	op          OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, repo workspace.Repository, meRepo me.Repository, invRepo invitation.Repository, fbDriver fbDr.Driver, emailDriver email.Driver, op OutputPort) UseCase {
	return &useCase{txp, dbp, repo, meRepo, invRepo, fbDriver, emailDriver, op}
}

func (u *useCase) Create(ctx context.Context, i CreateInput) (openapi.APIV1WorkspacesPostRes, error) {
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
		dn := member.NewDisplayName(meRes.Self().Name().ToString())
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

func (u *useCase) FindAllMembers(ctx context.Context, i FindAllMembersInput) (openapi.APIV1MembersGetRes, error) {
	exec := u.dbp.GetExecutor(ctx, true)
	currentWorkspaceID, err := u.fbDriver.MustGetCurrentWorkspaceFromCustomClaim(ctx, i.accountID)
	if err != nil {
		return nil, err
	}
	ms, err := u.repo.FindAllMembers(ctx, exec, currentWorkspaceID)
	if err != nil {
		return nil, err
	}
	return u.op.FindAllMembers(ms), nil
}

func (u *useCase) InviteMembers(ctx context.Context, i InviteMembersInput) (openapi.InviteMultipleUsersToWorkspaceRes, error) {
	// Exclude already registered members
	currentWorkspaceID, err := u.fbDriver.MustGetCurrentWorkspaceFromCustomClaim(ctx, i.AccountID)
	if err != nil {
		return nil, err
	}
	exec := u.dbp.GetExecutor(ctx, false)
	members, err := u.repo.FindAllMembers(ctx, exec, currentWorkspaceID)
	if err != nil {
		return nil, err
	}

	targets := make([]*invitation.Invitation, 0)
	alreadyRegisteredList := make([]*invitation.Invitation, 0)
	for _, im := range i.Invitations.AsSlice() {
		if members.Exist(im.InviteeEmail()) {
			alreadyRegisteredList = append(alreadyRegisteredList, im)
		} else {
			targets = append(targets, im)
		}
	}

	meRes, err := u.meRepo.FindLastLogin(ctx, exec, i.AccountID)
	if err != nil {
		return nil, err
	}

	inviter := workspace.NewInviter(meRes.Member(), meRes.Workspace())

	is := invitation.NewInvitations(targets)
	successSendMailList, failedSendMailList := u.emailDriver.SendInvitations(ctx, inviter, is)
	if err = u.repo.InviteMembers(ctx, exec, inviter, successSendMailList); err != nil {
		return nil, err
	}

	return u.op.InviteMembers(
		i.Invitations,
		invitation.NewInvitations(alreadyRegisteredList),
		successSendMailList,
		failedSendMailList,
	)
}

func (u *useCase) VerifyInvitationToken(ctx context.Context, i VerifyInvitationTokenInput) (openapi.VerifyInvitationRes, error) {
	p := u.dbp.GetExecutor(ctx, true)
	res, err := u.invRepo.FindByToken(ctx, p, i.Token)
	if err != nil {
		return nil, err
	}
	if res.IsExpired() {
		return nil, domainErr.NewExpiredInviteToken(res.Token().Value())
	}
	w, err := u.repo.FindInviterWorkspaceFromToken(ctx, p, i.Token)
	if err != nil {
		return nil, err
	}
	return u.op.VerifyInvitationToken(w, res), nil
}
