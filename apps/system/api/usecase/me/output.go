package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	Find(m *me.Me) (openapi.APIV1GetMeRes, error)
	AcceptInvitation(m *me.Me) (openapi.APIV1AcceptInvitationRes, error)
	UpdateProfile(m *me.Me) (openapi.APIV1UpdateProfileRes, error)
	UpdateMemberProfile(m *me.Me) (openapi.APIV1UpdateMeMemberProfileRes, error)
	UpdateProfilePhoto(m *me.Me) (openapi.APIV1UpdateProfilePhotoRes, error)
	RemoveProfilePhoto(m *me.Me) (openapi.APIV1RemoveProfilePhotoRes, error)
	LeaveWorkspace() (openapi.APIV1LeaveWorkspaceRes, error)
}
