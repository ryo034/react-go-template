package middleware

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"

	authDr "github.com/ryo034/react-go-template/apps/system/api/driver/auth"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"

	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
)

var UnauthenticatedRoutes = []string{
	"/api/v1/ping",
	"/api/v1/auth/otp",
	"/api/v1/auth/otp/verify",
	"/api/v1/auth/invitations/process/email",
	//"/api/v1/auth/invitations/process/oauth",
	"/api/v1/auth/invitations",
	//"/api/v1/auth/oauth",
}

type authentication struct {
	co     shared.ContextOperator
	dbp    core.Provider
	authDr authDr.Driver
	fbDr   fbDr.Driver
}

func NewAuthenticationMiddleware(co shared.ContextOperator, dbp core.Provider, authDr authDr.Driver, fbDr fbDr.Driver) Middleware {
	return &authentication{co, dbp, authDr, fbDr}
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (m *authentication) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(UnauthenticatedRoutes, r.URL.Path) {
			h.ServeHTTP(w, r)
			return
		}
		token, ok := findAuthorization(r.Header, "Bearer")
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx, err := m.checkAndSetToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *authentication) setRoleToContextIfExists(ctx context.Context, customClaims map[string]interface{}) (context.Context, error) {
	role, err := m.co.GetRole(ctx)
	if err == nil {
		return ctx, nil
	}

	var noSuchDataErr *domainErr.NoSuchData
	if !errors.As(err, &noSuchDataErr) {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
	}

	tmpRole := customClaims[fbDr.CustomClaimRoleKey]
	if tmpRole != nil && tmpRole != "" {
		if role, err = member.NewRole(tmpRole.(string)); err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Invalid Role"))
		}
		return m.co.SetRole(ctx, role), nil
	}

	// if role is not set in Context, get role from ProviderInfo
	pi, err := m.fbDr.GetProviderInfo(ctx, fbDr.GetProviderInfoRequiredOption{CurrentWorkspaceID: false})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
	}
	if pi.CustomClaim.Role != nil {
		return m.co.SetRole(ctx, *pi.CustomClaim.Role), nil
	}
	return ctx, nil
}

func (m *authentication) checkAndSetToken(ctx context.Context, t string) (context.Context, error) {
	token, err := m.fbDr.VerifyIDToken(ctx, t)
	if err != nil || token == nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
	}
	apUID, err := provider.NewUID(token.UID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
	}
	ctx = m.co.SetAuthProviderUID(ctx, apUID)
	if ctx, err = m.setRoleToContextIfExists(ctx, token.Claims); err != nil {
		return nil, err
	}

	if err = uuid.Validate(token.UID); err == nil {
		return m.co.SetClaim(m.co.SetUID(ctx, uuid.MustParse(token.UID).String()), token.Claims), nil
	}

	var aID uuid.UUID
	tkInClaim := token.Claims[fbDr.CustomClaimAccountIDKey]
	if tkInClaim != nil && tkInClaim != "" {
		if err = uuid.Validate(tkInClaim.(string)); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
		}
		aID = uuid.MustParse(tkInClaim.(string))
	} else {
		aID, err = m.authDr.FindAccountIDByAuthProviderUID(ctx, m.dbp.GetExecutor(ctx, true), apUID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
			}
			// only OAuth provider
			providerInfo, ok := token.Claims["firebase"].(map[string]interface{})
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
			}
			signInProvider, ok := providerInfo["sign_in_provider"].(string)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
			}
			switch signInProvider {
			case "google.com":
				tmpAID, _ := account.GenerateID()
				aID = tmpAID.Value()
			}
		}
		if err = m.fbDr.SetAccountIDToCustomClaim(ctx, account.NewIDFromUUID(aID)); err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
		}
	}
	return m.co.SetClaim(m.co.SetUID(ctx, aID.String()), token.Claims), nil
}
