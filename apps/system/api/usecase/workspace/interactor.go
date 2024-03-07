package workspace

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/notification"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Create(ctx context.Context, i CreateInput) (openapi.APIV1WorkspacesPostRes, error)
	FindAllMembers(ctx context.Context, i FindAllMembersInput) (openapi.APIV1MembersGetRes, error)
	InviteMembers(ctx context.Context, i InviteMembersInput) (openapi.InviteMultipleUsersToWorkspaceRes, error)
	RevokeInvitation(ctx context.Context, i RevokeInvitationInput) (openapi.RevokeInvitationRes, error)
	ResendInvitation(ctx context.Context, i ResendInvitationInput) (openapi.ResendInvitationRes, error)
	FindAllInvitation(ctx context.Context, i FindAllInvitationInput) (openapi.APIV1InvitationsGetRes, error)
	UpdateMemberRole(ctx context.Context, i UpdateMemberRoleInput) (openapi.APIV1MembersMemberIdRolePutRes, error)
}

type useCase struct {
	txp              core.TransactionProvider
	dbp              core.Provider
	repo             workspace.Repository
	meRepo           me.Repository
	invRepo          invitation.Repository
	notificationRepo notification.Repository
	op               OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, repo workspace.Repository, meRepo me.Repository, invRepo invitation.Repository, notificationRepo notification.Repository, op OutputPort) UseCase {
	return &useCase{txp, dbp, repo, meRepo, invRepo, notificationRepo, op}
}

func (u *useCase) Create(ctx context.Context, i CreateInput) (openapi.APIV1WorkspacesPostRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*workspace.Workspace, error) {
		meRes, err := u.meRepo.FindBeforeOnboard(pr, p, i.AccountID)
		if err != nil {
			return nil, err
		}
		w := i.Workspace
		wres, err := u.repo.Create(pr, p, w)
		if err != nil {
			return nil, err
		}
		m, err := member.GenerateAsWorkspaceOwner(meRes.Self(), member.NewDisplayName(meRes.Self().Name().ToString()))
		if err != nil {
			return nil, err
		}
		memRes, err := u.repo.AddMember(pr, p, wres, m)
		if err != nil {
			return nil, err
		}
		meRes, err = u.meRepo.Find(pr, p, memRes.ID())
		if err != nil {
			return nil, err
		}
		if err = u.meRepo.RecordLogin(pr, p, meRes); err != nil {
			return nil, err
		}
		return wres, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	res := result.Value(0).(*workspace.Workspace)
	return u.op.Create(res)
}

func (u *useCase) FindAllMembers(ctx context.Context, i FindAllMembersInput) (openapi.APIV1MembersGetRes, error) {
	exec := u.dbp.GetExecutor(ctx, true)
	ms, err := u.repo.FindAllMembers(ctx, exec, i.CurrentWorkspaceID)
	if err != nil {
		return nil, err
	}
	return u.op.FindAllMembers(ms)
}

func (u *useCase) InviteMembers(ctx context.Context, i InviteMembersInput) (openapi.InviteMultipleUsersToWorkspaceRes, error) {
	if len(i.Invitations) == 0 {
		return nil, nil
	}
	// Exclude already registered members
	exec := u.dbp.GetExecutor(ctx, false)
	meRes, err := u.meRepo.FindLastLogin(ctx, exec, i.AccountID)
	if err != nil {
		return nil, err
	}
	inviter := workspace.NewInviter(meRes.Member(), meRes.Workspace())

	invs := make([]*invitation.Invitation, 0, len(i.Invitations))
	for _, im := range i.Invitations {
		inv, err := invitation.GenInvitation(im.InviteeEmail, im.InviteeDisplayName, meRes.Member())
		if err != nil {
			return nil, err
		}
		invs = append(invs, inv)
	}
	invitations := invitation.NewInvitations(invs)

	members, err := u.repo.FindAllMembers(ctx, exec, i.CurrentWorkspaceID)
	if err != nil {
		return nil, err
	}

	targets := make([]*invitation.Invitation, 0)
	alreadyRegisteredList := make([]*invitation.Invitation, 0)
	for _, im := range invitations.AsSlice() {
		if members.Exist(im.InviteeEmail()) {
			alreadyRegisteredList = append(alreadyRegisteredList, im)
		} else {
			targets = append(targets, im)
		}
	}

	is := invitation.NewInvitations(targets)
	successSendMailList, failedSendMailList := u.notificationRepo.NotifyMembersInvited(ctx, inviter, is)
	if err = u.repo.InviteMembers(ctx, exec, inviter, successSendMailList); err != nil {
		return nil, err
	}

	return u.op.InviteMembers(
		invitations,
		invitation.NewInvitations(alreadyRegisteredList),
		successSendMailList,
		failedSendMailList,
	)
}

func (u *useCase) RevokeInvitation(ctx context.Context, i RevokeInvitationInput) (openapi.RevokeInvitationRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	inv, err := u.invRepo.Find(ctx, p, i.InvitationID)
	if err != nil {
		return nil, err
	}
	if err = inv.ValidateCanRevoke(i.AccountID); err != nil {
		return nil, err
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (invitation.Invitations, error) {
		if err = u.invRepo.Revoke(pr, p, i.InvitationID); err != nil {
			return nil, err
		}
		return u.repo.FindAllInvitations(pr, p, i.CurrentWorkspaceID)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	res := result.Value(0).(invitation.Invitations)
	return u.op.RevokeInvitation(res.ExcludeRevoked().ExcludeVerified().ExcludeAccepted().Sort())
}

func (u *useCase) ResendInvitation(ctx context.Context, i ResendInvitationInput) (openapi.ResendInvitationRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	inv, err := u.invRepo.Find(ctx, p, i.InvitationID)
	if err != nil {
		return nil, err
	}
	if err = inv.ValidateCanResend(i.AccountID); err != nil {
		return nil, err
	}

	meRes, err := u.meRepo.FindLastLogin(ctx, p, i.AccountID)
	if err != nil {
		return nil, err
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*invitation.Invitation, error) {
		res, err := u.invRepo.Resend(pr, p, i.InvitationID)
		if err != nil {
			return nil, err
		}
		inviter := workspace.NewInviter(meRes.Member(), meRes.Workspace())
		if err = u.notificationRepo.NotifyInvite(ctx, inviter, inv); err != nil {
			return nil, err
		}
		return res, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	res := result.Value(0).(*invitation.Invitation)
	return u.op.ResendInvitation(res)
}

func (u *useCase) FindAllInvitation(ctx context.Context, i FindAllInvitationInput) (openapi.APIV1InvitationsGetRes, error) {
	p := u.dbp.GetExecutor(ctx, true)
	res, err := u.repo.FindAllInvitations(ctx, p, i.CurrentWorkspaceID)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	if i.IsAccepted {
		return u.op.FindAllInvitation(res.OnlyAccepted().Sort())
	}
	return u.op.FindAllInvitation(res.ExcludeRevoked().ExcludeVerified().ExcludeAccepted().Sort())
}

func (u *useCase) UpdateMemberRole(ctx context.Context, i UpdateMemberRoleInput) (openapi.APIV1MembersMemberIdRolePutRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	m, err := u.meRepo.Find(ctx, p, i.MemberID)
	if err != nil {
		return nil, err
	}
	mem, err := m.Member().UpdateRole(i.Role)
	if err != nil {
		return nil, err
	}
	m = m.UpdateMember(mem)

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*member.Member, error) {
		res, err := u.repo.UpdateMemberRole(pr, p, m.Member())
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	res := result.Value(0).(*member.Member)
	return u.op.UpdateMemberRole(res)
}
