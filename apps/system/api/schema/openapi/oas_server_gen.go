// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// APIV1AuthOAuthPost implements POST /api/v1/auth/oauth operation.
	//
	// Auth by OAuth.
	//
	// POST /api/v1/auth/oauth
	APIV1AuthOAuthPost(ctx context.Context) (APIV1AuthOAuthPostRes, error)
	// APIV1AuthOtpPost implements POST /api/v1/auth/otp operation.
	//
	// One Time Password (OTP) to user.
	//
	// POST /api/v1/auth/otp
	APIV1AuthOtpPost(ctx context.Context, req *APIV1AuthOtpPostReq) (APIV1AuthOtpPostRes, error)
	// APIV1AuthOtpVerifyPost implements POST /api/v1/auth/otp/verify operation.
	//
	// Verify OTP sent by user.
	//
	// POST /api/v1/auth/otp/verify
	APIV1AuthOtpVerifyPost(ctx context.Context, req *APIV1AuthOtpVerifyPostReq) (APIV1AuthOtpVerifyPostRes, error)
	// APIV1MeGet implements GET /api/v1/me operation.
	//
	// Returns the admin user.
	//
	// GET /api/v1/me
	APIV1MeGet(ctx context.Context) (APIV1MeGetRes, error)
	// APIV1MeProfilePut implements PUT /api/v1/me/profile operation.
	//
	// Updates the user profile.
	//
	// PUT /api/v1/me/profile
	APIV1MeProfilePut(ctx context.Context, req *APIV1MeProfilePutReq) (APIV1MeProfilePutRes, error)
	// APIV1MembersGet implements GET /api/v1/members operation.
	//
	// Returns the members of the workspace.
	//
	// GET /api/v1/members
	APIV1MembersGet(ctx context.Context) (APIV1MembersGetRes, error)
	// APIV1PingGet implements GET /api/v1/ping operation.
	//
	// Checks if the server is running.
	//
	// GET /api/v1/ping
	APIV1PingGet(ctx context.Context) (APIV1PingGetRes, error)
	// APIV1WorkspacesGet implements GET /api/v1/workspaces operation.
	//
	// Returns the workspaces the user is a member of.
	//
	// GET /api/v1/workspaces
	APIV1WorkspacesGet(ctx context.Context) (APIV1WorkspacesGetRes, error)
	// APIV1WorkspacesPost implements POST /api/v1/workspaces operation.
	//
	// Creates a new workspace.
	//
	// POST /api/v1/workspaces
	APIV1WorkspacesPost(ctx context.Context, req *APIV1WorkspacesPostReq) (APIV1WorkspacesPostRes, error)
	// AcceptInvitation implements acceptInvitation operation.
	//
	// Accept an invitation to join a workspace.
	//
	// POST /api/v1/members/invitations/accept
	AcceptInvitation(ctx context.Context) (AcceptInvitationRes, error)
	// InviteMultipleUsersToWorkspace implements inviteMultipleUsersToWorkspace operation.
	//
	// Invite multiple users to the workspace by email.
	//
	// POST /api/v1/members/invitations/bulk
	InviteMultipleUsersToWorkspace(ctx context.Context, req *InviteMultipleUsersToWorkspaceReq) (InviteMultipleUsersToWorkspaceRes, error)
	// Login implements login operation.
	//
	// Login.
	//
	// POST /api/v1/login
	Login(ctx context.Context) (LoginRes, error)
	// ProcessInvitation implements processInvitation operation.
	//
	// Process an invitation by verifying token and email.
	//
	// POST /api/v1/auth/invitations/process
	ProcessInvitation(ctx context.Context, req *ProcessInvitationReq) (ProcessInvitationRes, error)
	// VerifyInvitation implements verifyInvitation operation.
	//
	// Verify Invitation.
	//
	// GET /api/v1/members/invitations/verify
	VerifyInvitation(ctx context.Context, params VerifyInvitationParams) (VerifyInvitationRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
