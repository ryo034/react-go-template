package auth

import (
	meDomain "github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	JwtToken(token string) *openapi.JwtToken
	InvitationByToken(ri meDomain.ReceivedInvitation) (openapi.GetInvitationByTokenRes, error)
}
