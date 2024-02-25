package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	Find(m *me.Me) (openapi.APIV1MeGetRes, error)
	AcceptInvitation(m *me.Me) (openapi.AcceptInvitationRes, error)
	UpdateProfile(m *me.Me) (openapi.APIV1MeProfilePutRes, error)
	UpdateMemberProfile(m *me.Me) (openapi.APIV1MeMemberProfilePutRes, error)
}
