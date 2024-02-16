package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1WorkspacesGet(ctx context.Context) (openapi.APIV1WorkspacesGetRes, error) {
	return nil, nil
}

func (s *service) APIV1WorkspacesPost(ctx context.Context, req *openapi.APIV1WorkspacesPostReq) (openapi.APIV1WorkspacesPostRes, error) {
	return s.ctrl.Workspace.Create(ctx, workspace.CreateInput{
		WorkspaceSubdomain: req.Subdomain,
	})
}

func (s *service) APIV1MembersGet(ctx context.Context) (openapi.APIV1MembersGetRes, error) {
	return s.ctrl.Workspace.FindAllMembers(ctx)
}

func (s *service) InviteMultipleUsersToWorkspace(ctx context.Context, req *openapi.InviteMultipleUsersToWorkspaceReq) (openapi.InviteMultipleUsersToWorkspaceRes, error) {
	ims := make([]workspace.Invitee, 0, len(req.Invitees))
	for _, m := range req.Invitees {
		ims = append(ims, workspace.Invitee{Email: m.Email, DisplayName: m.Name})
	}
	return s.ctrl.Workspace.InviteMembers(ctx, workspace.InviteesInput{InvitedMembers: ims})
}

func (s *service) VerifyInvitation(ctx context.Context, params openapi.VerifyInvitationParams) (openapi.VerifyInvitationRes, error) {
	return s.ctrl.Workspace.VerifyInvitationToken(ctx, workspace.VerifyInvitationTokenInput{
		Token: params.Token,
	})
}
