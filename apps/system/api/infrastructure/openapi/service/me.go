package service

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1MeGet(ctx context.Context) (openapi.APIV1MeGetRes, error) {
	return s.ctrl.Me.Find(ctx)
}

func (s *service) APIV1MeProfilePut(ctx context.Context, req *openapi.APIV1MeProfilePutReq) (openapi.APIV1MeProfilePutRes, error) {
	return s.ctrl.Me.UpdateProfile(ctx, me.UpdateProfileInput{Name: req.Profile.Name.Value})
}

func (s *service) AcceptInvitation(ctx context.Context, params openapi.AcceptInvitationParams) (openapi.AcceptInvitationRes, error) {
	return s.ctrl.Me.AcceptInvitation(ctx, me.AcceptInvitationInput{InvitationID: params.InvitationId})
}

func (s *service) APIV1MeMemberProfilePut(ctx context.Context, req *openapi.APIV1MeMemberProfilePutReq) (openapi.APIV1MeMemberProfilePutRes, error) {
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
