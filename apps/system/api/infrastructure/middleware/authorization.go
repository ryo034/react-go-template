package middleware

import (
	"log"
	"net/http"
	"regexp"
	"slices"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/authorization/casbin"
)

type authorization struct {
	co shared.ContextOperator
}

var UnauthorizedRoutes = []string{
	"/api/v1/ping",
	"/api/v1/auth/otp",
	"/api/v1/auth/otp/verify",
	"/api/v1/auth/invitations/process/email",
	"/api/v1/auth/invitations/process/oauth",
	"/api/v1/auth/invitations",
	"/api/v1/workspaces",
	"/api/v1/auth/oauth",
	"/api/v1/me",
	"/api/v1/me/profile",
	"/api/v1/me/member/profile",
	"/api/v1/members/invitations/{invitationID}/accept",
}

var UnauthorizedRegexRoutes = []*regexp.Regexp{
	regexp.MustCompile(`^/api/v1/members/invitations/.+/accept$`),
}

func NewAuthorizationMiddleware(co shared.ContextOperator) Middleware {
	return &authorization{co}
}

func IsUnauthorizedRegexRoutes(path string) bool {
	for _, r := range UnauthorizedRegexRoutes {
		if r.MatchString(path) {
			return true
		}
	}
	return false
}

func (am *authorization) isUnauthorizedRoute(path string) bool {
	return slices.Contains(UnauthenticatedRoutes, path) || slices.Contains(UnauthorizedRoutes, path) || IsUnauthorizedRegexRoutes(path)
}

func (am *authorization) Handler(h http.Handler) http.Handler {
	enf, err := casbin.NewEnforcer()
	if err != nil {
		log.Fatal(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if am.isUnauthorizedRoute(r.URL.Path) {
			h.ServeHTTP(w, r)
			return
		}

		obj := r.URL.Path
		act := r.Method
		ctx := r.Context()
		role, err := am.co.GetRole(ctx)
		if err != nil {
			http.Error(w, "role not found", http.StatusInternalServerError)
			return
		}

		allowed, reason, err := enf.EnforceEx(role.ToString(), obj, act)
		if err != nil {
			http.Error(w, "Enforce error", http.StatusInternalServerError)
			return
		}
		if allowed {
			h.ServeHTTP(w, r)
			return
		}

		if len(reason) > 0 && reason[3] == "deny" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		http.Error(w, "forbidden", http.StatusForbidden)
	})
}
