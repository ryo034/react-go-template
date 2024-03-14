// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// APIV1AcceptInvitation implements APIV1AcceptInvitation operation.
//
// Accept an invitation to join a workspace.
//
// POST /api/v1/members/invitations/{invitationId}/accept
func (UnimplementedHandler) APIV1AcceptInvitation(ctx context.Context, params APIV1AcceptInvitationParams) (r APIV1AcceptInvitationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1AuthByOAuth implements APIV1AuthByOAuth operation.
//
// Auth by OAuth.
//
// POST /api/v1/auth/oauth
func (UnimplementedHandler) APIV1AuthByOAuth(ctx context.Context) (r APIV1AuthByOAuthRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1AuthByOtp implements APIV1AuthByOtp operation.
//
// One Time Password (OTP) to user.
//
// POST /api/v1/auth/otp
func (UnimplementedHandler) APIV1AuthByOtp(ctx context.Context, req *APIV1AuthByOtpReq) (r APIV1AuthByOtpRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1CreateWorkspace implements APIV1CreateWorkspace operation.
//
// Creates a new workspace.
//
// POST /api/v1/workspaces
func (UnimplementedHandler) APIV1CreateWorkspace(ctx context.Context, req *APIV1CreateWorkspaceReq) (r APIV1CreateWorkspaceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1GetInvitationByToken implements APIV1GetInvitationByToken operation.
//
// Get Invitation by token.
//
// GET /api/v1/auth/invitations
func (UnimplementedHandler) APIV1GetInvitationByToken(ctx context.Context, params APIV1GetInvitationByTokenParams) (r APIV1GetInvitationByTokenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1GetInvitations implements APIV1GetInvitations operation.
//
// Returns the pending invitations (not used yet).
//
// GET /api/v1/invitations
func (UnimplementedHandler) APIV1GetInvitations(ctx context.Context, params APIV1GetInvitationsParams) (r APIV1GetInvitationsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1GetMe implements APIV1GetMe operation.
//
// Returns the admin user.
//
// GET /api/v1/me
func (UnimplementedHandler) APIV1GetMe(ctx context.Context) (r APIV1GetMeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1GetMembers implements APIV1GetMembers operation.
//
// Returns the members of the workspace.
//
// GET /api/v1/members
func (UnimplementedHandler) APIV1GetMembers(ctx context.Context) (r APIV1GetMembersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1GetWorkspaces implements APIV1GetWorkspaces operation.
//
// Returns the workspaces the user is a member of.
//
// GET /api/v1/workspaces
func (UnimplementedHandler) APIV1GetWorkspaces(ctx context.Context) (r APIV1GetWorkspacesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1InviteMultipleUsers implements APIV1InviteMultipleUsers operation.
//
// Invite multiple users to the workspace by email.
//
// POST /api/v1/members/invitations/bulk
func (UnimplementedHandler) APIV1InviteMultipleUsers(ctx context.Context, req *APIV1InviteMultipleUsersReq) (r APIV1InviteMultipleUsersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1LeaveWorkspace implements APIV1LeaveWorkspace operation.
//
// Leaves the workspace.
//
// POST /api/v1/me/workspace/leave
func (UnimplementedHandler) APIV1LeaveWorkspace(ctx context.Context) (r APIV1LeaveWorkspaceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1ProcessInvitationEmail implements APIV1ProcessInvitationEmail operation.
//
// Process an invitation by verifying token and email.
//
// POST /api/v1/auth/invitations/process/email
func (UnimplementedHandler) APIV1ProcessInvitationEmail(ctx context.Context, req *APIV1ProcessInvitationEmailReq) (r APIV1ProcessInvitationEmailRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1ProcessInvitationOAuth implements APIV1ProcessInvitationOAuth operation.
//
// Process an invitation by verifying token and OAuth, and register or add user to workspace.
//
// POST /api/v1/auth/invitations/process/oauth
func (UnimplementedHandler) APIV1ProcessInvitationOAuth(ctx context.Context, req *APIV1ProcessInvitationOAuthReq) (r APIV1ProcessInvitationOAuthRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1RemoveMember implements APIV1RemoveMember operation.
//
// Removes a member from the workspace.
//
// DELETE /api/v1/members/{memberId}
func (UnimplementedHandler) APIV1RemoveMember(ctx context.Context, params APIV1RemoveMemberParams) (r APIV1RemoveMemberRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1RemoveProfilePhoto implements APIV1RemoveProfilePhoto operation.
//
// Deletes the user profile photo.
//
// DELETE /api/v1/me/profile/photo
func (UnimplementedHandler) APIV1RemoveProfilePhoto(ctx context.Context) (r APIV1RemoveProfilePhotoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1ResendInvitation implements APIV1ResendInvitation operation.
//
// Resend invitation.
//
// POST /api/v1/members/invitations/{invitationId}/resend
func (UnimplementedHandler) APIV1ResendInvitation(ctx context.Context, params APIV1ResendInvitationParams) (r APIV1ResendInvitationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1RevokeInvitation implements APIV1RevokeInvitation operation.
//
// Revoke invitation.
//
// POST /api/v1/members/invitations/{invitationId}/revoke
func (UnimplementedHandler) APIV1RevokeInvitation(ctx context.Context, params APIV1RevokeInvitationParams) (r APIV1RevokeInvitationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1UpdateMeMemberProfile implements APIV1UpdateMeMemberProfile operation.
//
// Updates Me the member profile.
//
// PUT /api/v1/me/member/profile
func (UnimplementedHandler) APIV1UpdateMeMemberProfile(ctx context.Context, req *APIV1UpdateMeMemberProfileReq) (r APIV1UpdateMeMemberProfileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1UpdateMemberRole implements APIV1UpdateMemberRole operation.
//
// Updates the role of a member.
//
// PUT /api/v1/members/{memberId}/role
func (UnimplementedHandler) APIV1UpdateMemberRole(ctx context.Context, req *APIV1UpdateMemberRoleReq, params APIV1UpdateMemberRoleParams) (r APIV1UpdateMemberRoleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1UpdateProfile implements APIV1UpdateProfile operation.
//
// Updates the user profile.
//
// PUT /api/v1/me/profile
func (UnimplementedHandler) APIV1UpdateProfile(ctx context.Context, req *APIV1UpdateProfileReq) (r APIV1UpdateProfileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1UpdateProfilePhoto implements APIV1UpdateProfilePhoto operation.
//
// Updates the user profile photo.
//
// PUT /api/v1/me/profile/photo
func (UnimplementedHandler) APIV1UpdateProfilePhoto(ctx context.Context, req *APIV1UpdateProfilePhotoReq) (r APIV1UpdateProfilePhotoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1UpdateWorkspace implements APIV1UpdateWorkspace operation.
//
// Updates the workspace.
//
// PUT /api/v1/workspaces/{workspaceId}
func (UnimplementedHandler) APIV1UpdateWorkspace(ctx context.Context, req *APIV1UpdateWorkspaceReq, params APIV1UpdateWorkspaceParams) (r APIV1UpdateWorkspaceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIV1VerifyOTP implements APIV1VerifyOTP operation.
//
// Verify OTP sent by user.
//
// POST /api/v1/auth/otp/verify
func (UnimplementedHandler) APIV1VerifyOTP(ctx context.Context, req *APIV1VerifyOTPReq) (r APIV1VerifyOTPRes, _ error) {
	return r, ht.ErrNotImplemented
}

// Ping implements Ping operation.
//
// Checks if the server is running.
//
// GET /ping
func (UnimplementedHandler) Ping(ctx context.Context) (r PingRes, _ error) {
	return r, ht.ErrNotImplemented
}
