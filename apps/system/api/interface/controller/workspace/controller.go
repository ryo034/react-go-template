package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
	"time"
)

type Controller interface {
	Create(ctx context.Context, i CreateInput) (openapi.APIV1WorkspacesPostRes, error)
	FindAllMembers(ctx context.Context) (openapi.APIV1MembersGetRes, error)
	InviteMembers(ctx context.Context, i InvitedMembersInput) (openapi.InviteMultipleUsersToWorkspaceRes, error)
	VerifyInvitationToken(ctx context.Context, i VerifyInvitationTokenInput) (openapi.VerifyInvitationRes, error)
}

type controller struct {
	wuc  workspaceUc.UseCase
	resl shared.Resolver
	co   infraShared.ContextOperator
}

func NewController(wuc workspaceUc.UseCase, resl shared.Resolver, co infraShared.ContextOperator) Controller {
	return &controller{wuc, resl, co}
}

type CreateInput struct {
	WorkspaceSubdomain string
}

type InvitedMember struct {
	Email string
}

type InvitedMembersInput struct {
	InvitedMembers []InvitedMember
}

type VerifyInvitationTokenInput struct {
	Token uuid.UUID
}

func (c *controller) Create(ctx context.Context, i CreateInput) (openapi.APIV1WorkspacesPostRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1WorkspacesPostRes), nil
	}
	in, err := workspaceUc.NewCreateInput(i.WorkspaceSubdomain, aID)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1WorkspacesPostRes), nil
	}
	res, err := c.wuc.Create(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1WorkspacesPostRes), nil
	}
	return res, nil
}

func (c *controller) FindAllMembers(ctx context.Context) (openapi.APIV1MembersGetRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MembersGetRes), nil
	}
	in := workspaceUc.NewFindAllMembersInput(aID)
	res, err := c.wuc.FindAllMembers(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MembersGetRes), nil
	}
	return res, nil
}

func NewInviteMembersInput(aID account.ID, ims []InvitedMember) *workspaceUc.InviteMembersInput {
	expiredAt := time.Now()
	ivs := make([]*member.InvitedMember, len(ims))
	for _, im := range ims {
		em, err := account.NewEmail(im.Email)
		if err != nil {
			return nil
		}
		m, err := member.NewInvitedMemberFromEmail(em, expiredAt)
		if err != nil {
			return nil
		}
		ivs = append(ivs, m)
	}
	return &workspaceUc.InviteMembersInput{
		AccountID:      aID,
		InvitedMembers: member.NewInvitedMembers(ivs),
	}
}

func (c *controller) InviteMembers(ctx context.Context, i InvitedMembersInput) (openapi.InviteMultipleUsersToWorkspaceRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.InviteMultipleUsersToWorkspaceRes), nil
	}
	in := NewInviteMembersInput(aID, i.InvitedMembers)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.InviteMultipleUsersToWorkspaceRes), nil
	}
	res, err := c.wuc.InviteMembers(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.InviteMultipleUsersToWorkspaceRes), nil
	}
	return res, nil
}

func (c *controller) VerifyInvitationToken(ctx context.Context, i VerifyInvitationTokenInput) (openapi.VerifyInvitationRes, error) {
	in := &workspaceUc.VerifyInvitationTokenInput{Token: i.Token}
	return c.wuc.VerifyInvitationToken(ctx, in)
}
