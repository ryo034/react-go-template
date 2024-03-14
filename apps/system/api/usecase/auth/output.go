//go:generate gomockhandler -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	JwtToken(token string) *openapi.JwtToken
	InvitationByToken(ri me.ReceivedInvitation) (openapi.APIV1GetInvitationByTokenRes, error)
	APIV1ProcessInvitationOAuth(me *me.Me) (openapi.APIV1ProcessInvitationOAuthRes, error)
	AuthByOAuth(me *me.Me) (openapi.APIV1AuthByOAuthRes, error)
}
