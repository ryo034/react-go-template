package middleware

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

//type Middleware struct {
//	firebase *fb.Firebase
//	co       shared.ContextOperator
//	dbp      core.Provider
//	authDr   authDr.Driver
//	fbDr     fbDr.Driver
//}

type Middleware struct{}

func NewSecMiddleware() *Middleware {
	return &Middleware{}
}

// HandleBearer run only when the authentication method is Bearer
// Debugできなくなったのでinfrastructureのmiddlewareを使う
func (m *Middleware) HandleBearer(ctx context.Context, operationName string, t openapi.Bearer) (context.Context, error) {
	//return m.checkAndSetToken(ctx, t.GetToken())
	return ctx, nil
}

//func (m *Middleware) setRoleToContextIfNotExists(ctx context.Context, customClaims map[string]interface{}) (context.Context, error) {
//	var noSuchDataErr *domainErr.NoSuchData
//	role, err := m.co.GetRole(ctx)
//	if err == nil {
//		return ctx, nil
//	}
//	if !errors.As(err, &noSuchDataErr) {
//		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
//	}
//	//role, err = m.fbDr.GetRoleFromCustomClaim(ctx)
//	//if err != nil {
//	//	return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
//	//}
//	tmpRole := customClaims[fbDr.CustomClaimRoleKey].(string)
//	if tmpRole == "" {
//		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Invalid Role"))
//	}
//	if role, err = member.NewRole(tmpRole); err != nil {
//		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Invalid Role"))
//	}
//	return m.co.SetRole(ctx, role), nil
//}
//
//func (m *Middleware) checkAndSetToken(ctx context.Context, t string) (context.Context, error) {
//	token, err := m.firebase.Auth.VerifyIDToken(ctx, t)
//	if err != nil || token == nil {
//		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
//	}
//	apUID, err := provider.NewUID(token.UID)
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
//	}
//	ctx = m.co.SetAuthProviderUID(ctx, apUID)
//	if ctx, err = m.setRoleToContextIfNotExists(ctx, token.Claims); err != nil {
//		return nil, err
//	}
//
//	err = uuid.Validate(token.UID)
//	if err == nil {
//		return m.co.SetClaim(m.co.SetUID(ctx, uuid.MustParse(token.UID).String()), token.Claims), nil
//	}
//
//	var aID uuid.UUID
//	tkInClaim := token.Claims[fbDr.CustomClaimAccountIDKey]
//	if tkInClaim != nil && tkInClaim != "" {
//		if err = uuid.Validate(tkInClaim.(string)); err != nil {
//			return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
//		}
//		aID = uuid.MustParse(tkInClaim.(string))
//	} else {
//		aID, err = m.authDr.FindAccountIDByAuthProviderUID(ctx, m.dbp.GetExecutor(ctx, true), apUID)
//		if err != nil {
//			if !errors.Is(err, sql.ErrNoRows) {
//				return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
//			}
//			// only OAuth provider
//			providerInfo, ok := token.Claims["firebase"].(map[string]interface{})
//			if !ok {
//				return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
//			}
//			signInProvider, ok := providerInfo["sign_in_provider"].(string)
//			if !ok {
//				return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
//			}
//			switch signInProvider {
//			case "google.com":
//				tmpAID, _ := account.GenerateID()
//				aID = tmpAID.Value()
//			}
//		}
//		if err = m.fbDr.SetAccountIDToCustomClaim(ctx, account.NewIDFromUUID(aID)); err != nil {
//			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
//		}
//	}
//	return m.co.SetClaim(m.co.SetUID(ctx, aID.String()), token.Claims), nil
//}
