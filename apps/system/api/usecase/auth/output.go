package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	JwtToken(token string) *openapi.JwtToken
}
