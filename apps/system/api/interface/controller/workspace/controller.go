package workspace

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"

	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"

	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

type Controller interface {
	Create(ctx context.Context, i CreateInput) (openapi.APIV1CreateWorkspaceRes, error)
	FindAllMembers(ctx context.Context) (openapi.APIV1GetMembersRes, error)
	InviteMembers(ctx context.Context, i InviteesInput) (openapi.APIV1InviteMultipleUsersRes, error)
	RevokeInvitation(ctx context.Context, i RevokeInvitationInput) (openapi.APIV1RevokeInvitationRes, error)
	ResendInvitation(ctx context.Context, i ResendInvitationInput) (openapi.APIV1ResendInvitationRes, error)
	FindAllInvitation(ctx context.Context, i FindAllInvitationInput) (openapi.APIV1GetInvitationsRes, error)
	UpdateMemberRole(ctx context.Context, i UpdateMemberRoleInput) (openapi.APIV1UpdateMemberRoleRes, error)
	UpdateWorkspace(ctx context.Context, i UpdateWorkspaceInput) (openapi.APIV1UpdateWorkspaceRes, error)
	Leave(ctx context.Context, i LeaveInput) (openapi.APIV1RemoveMemberRes, error)
}

type controller struct {
	wuc  workspaceUc.UseCase
	fbDr fbDr.Driver
	resl shared.Resolver
	co   infraShared.ContextOperator
}

func NewController(wuc workspaceUc.UseCase, fbDr fbDr.Driver, resl shared.Resolver, co infraShared.ContextOperator) Controller {
	return &controller{wuc, fbDr, resl, co}
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

type ResendInvitationInput struct {
	InvitationID uuid.UUID
}

type FindAllInvitationInput struct {
	Status string
}

type UpdateMemberRoleInput struct {
	MemberID uuid.UUID
	Role     string
}

type UpdateWorkspaceInput struct {
	WorkspaceID uuid.UUID
	Name        string
	Subdomain   string
}

type LeaveInput struct {
	MemberID uuid.UUID
}

func (c *controller) Create(ctx context.Context, i CreateInput) (openapi.APIV1CreateWorkspaceRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1CreateWorkspaceRes), nil
	}
	in, err := workspaceUc.NewCreateInput(i.WorkspaceSubdomain, aID)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1CreateWorkspaceRes), nil
	}
	res, err := c.wuc.Create(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1CreateWorkspaceRes), nil
	}
	return res, nil
}

func (c *controller) FindAllMembers(ctx context.Context) (openapi.APIV1GetMembersRes, error) {
	clm, err := c.fbDr.GetProviderInfo(ctx, fbDr.GetProviderInfoRequiredOption{CurrentWorkspaceID: true})
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1GetMembersRes), nil
	}
	in := workspaceUc.FindAllMembersInput{CurrentWorkspaceID: *clm.CustomClaim.CurrentWorkspaceID}
	res, err := c.wuc.FindAllMembers(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1GetMembersRes), nil
	}
	return res, nil
}

func NewInviteMembersInput(cwID workspace.ID, aID account.ID, ims []Invitee) (workspaceUc.InviteMembersInput, error) {
	ivs := make([]workspaceUc.CreateInvitation, 0, len(ims))
	for _, im := range ims {
		ivs = append(ivs, workspaceUc.CreateInvitation{
			InviteeEmail:       im.Email,
			InviteeDisplayName: im.DisplayName,
		})
	}
	return workspaceUc.InviteMembersInput{
		CurrentWorkspaceID: cwID,
		AccountID:          aID,
		Invitations:        ivs,
	}, nil
}

func (c *controller) InviteMembers(ctx context.Context, i InviteesInput) (openapi.APIV1InviteMultipleUsersRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1InviteMultipleUsersRes), nil
	}
	clm, err := c.fbDr.GetProviderInfo(ctx, fbDr.GetProviderInfoRequiredOption{CurrentWorkspaceID: true})
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1InviteMultipleUsersRes), nil
	}
	in, err := NewInviteMembersInput(*clm.CustomClaim.CurrentWorkspaceID, aID, i.InvitedMembers)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1InviteMultipleUsersRes), nil
	}
	res, err := c.wuc.InviteMembers(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1InviteMultipleUsersRes), nil
	}
	return res, nil
}

func (c *controller) RevokeInvitation(ctx context.Context, i RevokeInvitationInput) (openapi.APIV1RevokeInvitationRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1RevokeInvitationRes), nil
	}
	clm, err := c.fbDr.GetProviderInfo(ctx, fbDr.GetProviderInfoRequiredOption{CurrentWorkspaceID: true})
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1RevokeInvitationRes), nil
	}
	in := workspaceUc.APIV1RevokeInvitationInput{
		AccountID:          aID,
		CurrentWorkspaceID: *clm.CustomClaim.CurrentWorkspaceID,
		InvitationID:       invitation.NewID(i.InvitationID),
	}
	res, err := c.wuc.APIV1RevokeInvitation(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1RevokeInvitationRes), nil
	}
	return res, nil
}

func (c *controller) ResendInvitation(ctx context.Context, i ResendInvitationInput) (openapi.APIV1ResendInvitationRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1ResendInvitationRes), nil
	}
	clm, err := c.fbDr.GetProviderInfo(ctx, fbDr.GetProviderInfoRequiredOption{CurrentWorkspaceID: true})
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1ResendInvitationRes), nil
	}
	in := workspaceUc.APIV1ResendInvitationInput{
		AccountID:          aID,
		CurrentWorkspaceID: *clm.CustomClaim.CurrentWorkspaceID,
		InvitationID:       invitation.NewID(i.InvitationID),
	}
	res, err := c.wuc.APIV1ResendInvitation(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1ResendInvitationRes), nil
	}
	return res, nil
}

func (c *controller) FindAllInvitation(ctx context.Context, i FindAllInvitationInput) (openapi.APIV1GetInvitationsRes, error) {
	accepted := false
	if i.Status == "accepted" {
		accepted = true
	}
	clm, err := c.fbDr.GetProviderInfo(ctx, fbDr.GetProviderInfoRequiredOption{CurrentWorkspaceID: true})
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1GetInvitationsRes), nil
	}
	in := workspaceUc.FindAllInvitationInput{
		CurrentWorkspaceID: *clm.CustomClaim.CurrentWorkspaceID,
		IsAccepted:         accepted,
	}
	res, err := c.wuc.FindAllInvitation(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1GetInvitationsRes), nil
	}
	if res == nil {
		return &openapi.InvitationsResponse{Invitations: nil}, nil
	}
	return res, nil
}

func (c *controller) UpdateMemberRole(ctx context.Context, i UpdateMemberRoleInput) (openapi.APIV1UpdateMemberRoleRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1UpdateMemberRoleRes), nil
	}
	role, err := member.NewRole(i.Role)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1UpdateMemberRoleRes), nil
	}
	in := workspaceUc.UpdateMemberRoleInput{
		AccountID: aID,
		MemberID:  member.NewIDFromUUID(i.MemberID),
		Role:      role,
	}
	res, err := c.wuc.UpdateMemberRole(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1UpdateMemberRoleRes), nil
	}
	return res, nil
}

func (c *controller) UpdateWorkspace(ctx context.Context, i UpdateWorkspaceInput) (openapi.APIV1UpdateWorkspaceRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1UpdateWorkspaceRes), nil
	}
	n, err := workspace.NewName(i.Name)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1UpdateWorkspaceRes), nil
	}
	su, err := workspace.NewSubdomain(i.Subdomain)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1UpdateWorkspaceRes), nil
	}
	in := workspaceUc.UpdateWorkspaceInput{
		AccountID:   aID,
		WorkspaceID: workspace.NewIDFromUUID(i.WorkspaceID),
		Name:        n,
		Subdomain:   su,
	}
	res, err := c.wuc.UpdateWorkspace(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1UpdateWorkspaceRes), nil
	}
	return res, nil
}

func (c *controller) Leave(ctx context.Context, i LeaveInput) (openapi.APIV1RemoveMemberRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1RemoveMemberRes), nil
	}
	in := workspaceUc.LeaveInput{
		ExecutorID: aID,
		MemberID:   member.NewIDFromUUID(i.MemberID),
	}
	res, err := c.wuc.Leave(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1RemoveMemberRes), nil
	}
	return res, nil
}
