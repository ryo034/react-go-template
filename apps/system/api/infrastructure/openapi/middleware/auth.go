package middleware

import (
	"context"
	"fmt"
	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Middleware struct {
	firebase *fb.Firebase
	co       shared.ContextOperator
}

// HandleBearer run only when the authentication method is Bearer
func (m *Middleware) HandleBearer(ctx context.Context, operationName string, t openapi.Bearer) (context.Context, error) {
	return m.checkAndSetToken(ctx, t.GetToken())
}

// checkAndSetToken return JWT information in the key token in the argument context
func (m *Middleware) checkAndSetToken(ctx context.Context, t string) (context.Context, error) {
	token, err := m.firebase.Auth.VerifyIDToken(ctx, t)
	if err != nil || token == nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
	}
	// Return JWT information in the key token in the argument context
	return m.co.SetClaim(m.co.SetUID(ctx, token.UID), token.Claims), nil
}

func NewSecMiddleware(firebase *fb.Firebase, co shared.ContextOperator) *Middleware {
	return &Middleware{firebase, co}
}
