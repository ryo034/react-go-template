package workspace

import (
	"context"

	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

type Controller interface {
	Create(ctx context.Context, i CreateInput) (openapi.APIV1WorkspacesPostRes, error)
	FindAllMembers(ctx context.Context) (openapi.APIV1MembersGetRes, error)
	InviteMembers(ctx context.Context, i InviteesInput) (openapi.InviteMultipleUsersToWorkspaceRes, error)
	RevokeInvitation(ctx context.Context, i RevokeInvitationInput) (openapi.RevokeInvitationRes, error)
	FindAllInvitation(ctx context.Context, i FindAllInvitationInput) (openapi.APIV1InvitationsGetRes, error)
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

type Invitee struct {
	Email       string
	DisplayName string
}

type InviteesInput struct {
	InvitedMembers []Invitee
}

type RevokeInvitationInput struct {
	InvitationID uuid.UUID
}

type FindAllInvitationInput struct {
	Status string
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

func NewInviteMembersInput(aID account.ID, ims []Invitee) (workspaceUc.InviteMembersInput, error) {
	ivs := make([]*invitation.Invitation, 0, len(ims))
	for _, im := range ims {
		i, err := invitation.GenInvitation(im.Email, im.DisplayName)
		if err != nil {
			return workspaceUc.InviteMembersInput{}, err
		}
		ivs = append(ivs, i)
	}
	return workspaceUc.InviteMembersInput{
		AccountID:   aID,
		Invitations: invitation.NewInvitations(ivs),
	}, nil
}

func (c *controller) InviteMembers(ctx context.Context, i InviteesInput) (openapi.InviteMultipleUsersToWorkspaceRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.InviteMultipleUsersToWorkspaceRes), nil
	}
	in, err := NewInviteMembersInput(aID, i.InvitedMembers)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.InviteMultipleUsersToWorkspaceRes), nil
	}
	res, err := c.wuc.InviteMembers(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.InviteMultipleUsersToWorkspaceRes), nil
	}
	return res, nil
}

func (c *controller) RevokeInvitation(ctx context.Context, i RevokeInvitationInput) (openapi.RevokeInvitationRes, error) {
	in := workspaceUc.RevokeInvitationInput{InvitationID: invitation.NewID(i.InvitationID)}
	res, err := c.wuc.RevokeInvitation(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.RevokeInvitationRes), nil
	}
	return res, nil
}

func (c *controller) FindAllInvitation(ctx context.Context, i FindAllInvitationInput) (openapi.APIV1InvitationsGetRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1InvitationsGetRes), nil
	}

	accepted := false
	if i.Status == "accepted" {
		accepted = true
	}
	in := workspaceUc.FindAllInvitationInput{AccountID: aID, IsAccepted: accepted}
	res, err := c.wuc.FindAllInvitation(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1InvitationsGetRes), nil
	}
	return res, nil
}
