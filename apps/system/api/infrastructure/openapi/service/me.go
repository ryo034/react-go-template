package service

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1GetMe(ctx context.Context) (openapi.APIV1GetMeRes, error) {
	return s.ctrl.Me.Find(ctx)
}

func (s *service) APIV1UpdateProfile(ctx context.Context, req *openapi.APIV1UpdateProfileReq) (openapi.APIV1UpdateProfileRes, error) {
	return s.ctrl.Me.UpdateName(ctx, me.UpdateProfileInput{Name: req.Profile.Name.Value})
}

func (s *service) APIV1AcceptInvitation(ctx context.Context, params openapi.APIV1AcceptInvitationParams) (openapi.APIV1AcceptInvitationRes, error) {
	return s.ctrl.Me.AcceptInvitation(ctx, me.AcceptInvitationInput{InvitationID: params.InvitationId})
}

func (s *service) APIV1UpdateMeMemberProfile(ctx context.Context, req *openapi.APIV1UpdateMeMemberProfileReq) (openapi.APIV1UpdateMeMemberProfileRes, error) {
	idNum := ""
	if req.MemberProfile.IdNumber.Set {
		idNum = req.MemberProfile.IdNumber.Value
	}
	bio := ""
	if req.MemberProfile.Bio.Set {
		bio = req.MemberProfile.Bio.Value
	}
	return s.ctrl.Me.UpdateMemberProfile(ctx, me.UpdateMemberProfileInput{
		DisplayName: req.MemberProfile.DisplayName,
		IdNumber:    idNum,
		Bio:         bio,
	})
}

func (s *service) APIV1UpdateProfilePhoto(ctx context.Context, req *openapi.APIV1UpdateProfilePhotoReq) (openapi.APIV1UpdateProfilePhotoRes, error) {
	return s.ctrl.Me.UpdateProfilePhoto(ctx, me.UpdateProfilePhotoInput{
		File:   req.GetPhoto().File,
		Name:   req.GetPhoto().Name,
		Header: req.GetPhoto().Header,
	})
}

func (s *service) APIV1RemoveProfilePhoto(ctx context.Context) (openapi.APIV1RemoveProfilePhotoRes, error) {
	return s.ctrl.Me.RemoveProfilePhoto(ctx)
}

func (s *service) APIV1LeaveWorkspace(ctx context.Context) (openapi.APIV1LeaveWorkspaceRes, error) {
	return s.ctrl.Me.LeaveWorkspace(ctx)
}
