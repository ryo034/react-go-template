package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	authUc "github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
)

func NewPresenter() authUc.OutputPort {
	return &presenter{}
}

type presenter struct {
}

func (p *presenter) JwtToken(token string) *openapi.JwtToken {
	return &openapi.JwtToken{Token: token}
}
