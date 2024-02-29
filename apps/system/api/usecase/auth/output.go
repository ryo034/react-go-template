package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	JwtToken(token string) *openapi.JwtToken
	InvitationByToken(ri me.ReceivedInvitation) (openapi.GetInvitationByTokenRes, error)
	AuthByAuth(me *me.Me) (openapi.APIV1AuthOAuthPostRes, error)
}
