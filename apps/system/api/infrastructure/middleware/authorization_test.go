package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"

	"gotest.tools/v3/assert"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
)

func TestAuthorizationMiddleware_UnauthorizedRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		role           member.Role
		path           string
		method         string
		expectedStatus int
	}{
		{"Me Owner", member.RoleOwner, "/api/v1/me", http.MethodGet, http.StatusOK},
		{"Me Admin", member.RoleAdmin, "/api/v1/me", http.MethodGet, http.StatusOK},
		{"Me Member", member.RoleMember, "/api/v1/me", http.MethodGet, http.StatusOK},
		{"Me Guest", member.RoleGuest, "/api/v1/me", http.MethodGet, http.StatusOK},

		{"Me Profile Owner", member.RoleOwner, "/api/v1/me/profile", http.MethodPut, http.StatusOK},
		{"Me Profile Admin", member.RoleAdmin, "/api/v1/me/profile", http.MethodPut, http.StatusOK},
		{"Me Profile Member", member.RoleMember, "/api/v1/me/profile", http.MethodPut, http.StatusOK},
		{"Me Profile Guest", member.RoleGuest, "/api/v1/me/profile", http.MethodPut, http.StatusOK},

		{"Me Member Profile Owner", member.RoleOwner, "/api/v1/me/member/profile", http.MethodPut, http.StatusOK},
		{"Me Member Profile Admin", member.RoleAdmin, "/api/v1/me/member/profile", http.MethodPut, http.StatusOK},
		{"Me Member Profile Member", member.RoleMember, "/api/v1/me/member/profile", http.MethodPut, http.StatusOK},
		{"Me Member Profile Guest", member.RoleGuest, "/api/v1/me/member/profile", http.MethodPut, http.StatusOK},

		{"Accept invitation Owner", member.RoleOwner, "/api/v1/members/invitations/{invitationId}/accept", http.MethodPost, http.StatusOK},
		{"Accept invitation Admin", member.RoleAdmin, "/api/v1/members/invitations/{invitationId}/accept", http.MethodPost, http.StatusOK},
		{"Accept invitation Member", member.RoleMember, "/api/v1/members/invitations/{invitationId}/accept", http.MethodPost, http.StatusOK},
		{"Accept invitation Guest", member.RoleGuest, "/api/v1/members/invitations/{invitationId}/accept", http.MethodPost, http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(tc.method, tc.path, nil)
			rr := httptest.NewRecorder()

			middleware := NewAuthorizationMiddleware(nil)
			handler := middleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handler.ServeHTTP(rr, req)
			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}

func TestAuthorizationMiddleware_Forbidden(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockContextOperator := shared.NewMockContextOperator(ctrl)

	tests := []struct {
		name           string
		role           member.Role
		path           string
		method         string
		expectedStatus int
	}{
		{"Invite Bulk Member", member.RoleMember, "/api/v1/members/invitations/bulk", http.MethodPost, http.StatusForbidden},
		{"Invite Bulk Guest", member.RoleGuest, "/api/v1/members/invitations/bulk", http.MethodPost, http.StatusForbidden},

		{"Revoke invitation Member", member.RoleMember, "/api/v1/members/invitations/{invitationId}/revoke", http.MethodPost, http.StatusForbidden},
		{"Revoke invitation Guest", member.RoleGuest, "/api/v1/members/invitations/{invitationId}/revoke", http.MethodPost, http.StatusForbidden},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockContextOperator.EXPECT().GetRole(gomock.Any()).Return(tc.role, nil)

			req, _ := http.NewRequest(tc.method, tc.path, nil)
			rr := httptest.NewRecorder()

			middleware := NewAuthorizationMiddleware(mockContextOperator)
			handler := middleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handler.ServeHTTP(rr, req)
			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}

func TestAuthorizationMiddleware_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockContextOperator := shared.NewMockContextOperator(ctrl)

	tests := []struct {
		name           string
		role           member.Role
		path           string
		method         string
		expectedStatus int
	}{

		{"Get Members Owner", member.RoleOwner, "/api/v1/members", http.MethodGet, http.StatusOK},
		{"Get Members Admin", member.RoleAdmin, "/api/v1/members", http.MethodGet, http.StatusOK},
		{"Get Members Member", member.RoleMember, "/api/v1/members", http.MethodGet, http.StatusOK},
		{"Get Members Guest", member.RoleGuest, "/api/v1/members", http.MethodGet, http.StatusOK},

		{"Get invitations Owner", member.RoleOwner, "/api/v1/invitations", http.MethodGet, http.StatusOK},
		{"Get invitations Admin", member.RoleAdmin, "/api/v1/invitations", http.MethodGet, http.StatusOK},
		{"Get invitations Member", member.RoleMember, "/api/v1/invitations", http.MethodGet, http.StatusOK},
		{"Get invitations Guest", member.RoleGuest, "/api/v1/invitations", http.MethodGet, http.StatusOK},

		{"Invite Bulk Owner", member.RoleOwner, "/api/v1/members/invitations/bulk", http.MethodPost, http.StatusOK},
		{"Invite Bulk Admin", member.RoleAdmin, "/api/v1/members/invitations/bulk", http.MethodPost, http.StatusOK},

		{"Revoke invitation Owner", member.RoleOwner, "/api/v1/members/invitations/{invitationId}/revoke", http.MethodPost, http.StatusOK},
		{"Revoke invitation Admin", member.RoleAdmin, "/api/v1/members/invitations/{invitationId}/revoke", http.MethodPost, http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockContextOperator.EXPECT().GetRole(gomock.Any()).Return(tc.role, nil)

			req, _ := http.NewRequest(tc.method, tc.path, nil)
			rr := httptest.NewRecorder()

			middleware := NewAuthorizationMiddleware(mockContextOperator)
			handler := middleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handler.ServeHTTP(rr, req)
			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}
