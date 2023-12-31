package middleware

import (
	"context"
	"fmt"
	"strings"

	ga "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Authentication struct {
	firebase *fb.Firebase
	co       shared.ContextOperator
}

func NewAuthentication(firebase *fb.Firebase, co shared.ContextOperator) *Authentication {
	authInst := new(Authentication)
	authInst.firebase = firebase
	authInst.co = co
	return authInst
}

func (au *Authentication) Authenticate(ctx context.Context, req connect.AnyRequest) (context.Context, error) {
	authorization := req.Header().Get("Authorization")
	if authorization == "" {
		return nil, status.Errorf(codes.Unauthenticated, "token is required")
	}
	idToken := strings.Split(authorization, " ")[1]
	if idToken == "" {
		return nil, status.Errorf(codes.Unauthenticated, "token is required")
	}
	return au.setToken(ctx, idToken)
}

// LooseAuthenticate access from the web is authenticated
func (au *Authentication) LooseAuthenticate(ctx context.Context) (newCtx context.Context, err error) {
	idToken, _ := ga.AuthFromMD(ctx, "Bearer")
	if idToken == "" {
		return ctx, nil
	}
	return au.setToken(ctx, idToken)
}

// setToken return JWT information in the key token in the argument context
func (au *Authentication) setToken(ctx context.Context, idToken string) (newCtx context.Context, err error) {
	token, err := au.firebase.Auth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
	}

	// Return JWT information in the key token in the argument context
	return au.co.SetClaim(au.co.SetUID(ctx, token.UID), token.Claims), nil
}
