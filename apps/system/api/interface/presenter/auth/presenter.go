package auth

import (
	meDomain "github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	authUc "github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
)

func NewPresenter(meAdapter me.Adapter) authUc.OutputPort {
	return &presenter{meAdapter}
}

type presenter struct {
	meAdapter me.Adapter
}

func (p *presenter) JwtToken(token string) *openapi.JwtToken {
	return &openapi.JwtToken{Token: token}
}

func (p *presenter) InvitationByToken(ri meDomain.ReceivedInvitation) (openapi.GetInvitationByTokenRes, error) {
	res, err := p.meAdapter.AdaptReceivedInvitation(ri)
	if err != nil {
		return nil, err
	}
	return &openapi.GetInvitationByTokenResponse{ReceivedInvitation: res}, nil
}
