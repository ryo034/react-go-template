package service

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1GetWorkspaces(ctx context.Context) (openapi.APIV1GetWorkspacesRes, error) {
	return nil, nil
}

func (s *service) APIV1CreateWorkspace(ctx context.Context, req *openapi.APIV1CreateWorkspaceReq) (openapi.APIV1CreateWorkspaceRes, error) {
	return s.ctrl.Workspace.Create(ctx, workspace.CreateInput{
		WorkspaceSubdomain: req.Subdomain,
	})
}

func (s *service) APIV1GetMembers(ctx context.Context) (openapi.APIV1GetMembersRes, error) {
	return s.ctrl.Workspace.FindAllMembers(ctx)
}

func (s *service) APIV1InviteMultipleUsers(ctx context.Context, req *openapi.APIV1InviteMultipleUsersReq) (openapi.APIV1InviteMultipleUsersRes, error) {
	ims := make([]workspace.Invitee, 0, len(req.Invitees))
	for _, m := range req.Invitees {
		ims = append(ims, workspace.Invitee{Email: m.Email, DisplayName: m.Name})
	}
	return s.ctrl.Workspace.InviteMembers(ctx, workspace.InviteesInput{InvitedMembers: ims})
}

func (s *service) APIV1RevokeInvitation(ctx context.Context, params openapi.APIV1RevokeInvitationParams) (openapi.APIV1RevokeInvitationRes, error) {
	return s.ctrl.Workspace.RevokeInvitation(ctx, workspace.RevokeInvitationInput{
		InvitationID: params.InvitationId,
	})
}

func (s *service) APIV1ResendInvitation(ctx context.Context, params openapi.APIV1ResendInvitationParams) (openapi.APIV1ResendInvitationRes, error) {
	return s.ctrl.Workspace.ResendInvitation(ctx, workspace.ResendInvitationInput{
		InvitationID: params.InvitationId,
	})
}

func (s *service) APIV1GetInvitations(ctx context.Context, params openapi.APIV1GetInvitationsParams) (openapi.APIV1GetInvitationsRes, error) {
	status := ""
	if params.Status.IsSet() {
		//v, _ := params.Status.Get()
		status = string(params.Status.Value)
	}
	return s.ctrl.Workspace.FindAllInvitation(ctx, workspace.FindAllInvitationInput{Status: status})
}

func (s *service) APIV1UpdateMemberRole(ctx context.Context, req *openapi.APIV1UpdateMemberRoleReq, params openapi.APIV1UpdateMemberRoleParams) (openapi.APIV1UpdateMemberRoleRes, error) {
	return s.ctrl.Workspace.UpdateMemberRole(ctx, workspace.UpdateMemberRoleInput{
		MemberID: params.MemberId,
		Role:     string(req.GetRole()),
	})
}

func (s *service) APIV1UpdateWorkspace(ctx context.Context, req *openapi.APIV1UpdateWorkspaceReq, params openapi.APIV1UpdateWorkspaceParams) (openapi.APIV1UpdateWorkspaceRes, error) {
	n, ok := req.Name.Get()
	if !ok {
		n = ""
	}
	su, ok := req.Subdomain.Get()
	if !ok {
		su = ""
	}
	return s.ctrl.Workspace.UpdateWorkspace(ctx, workspace.UpdateWorkspaceInput{
		WorkspaceID: params.WorkspaceId,
		Name:        n,
		Subdomain:   su,
	})
}

func (s *service) APIV1RemoveMember(ctx context.Context, params openapi.APIV1RemoveMemberParams) (openapi.APIV1RemoveMemberRes, error) {
	return s.ctrl.Workspace.Leave(ctx, workspace.LeaveInput{
		MemberID: params.MemberId,
	})
}
