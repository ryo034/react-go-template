package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type CreateInput struct {
	Workspace *workspace.Workspace
	AccountID account.ID
}

func NewCreateInput(subdomain string, aID account.ID) (CreateInput, error) {
	s, err := workspace.NewSubdomain(subdomain)
	if err != nil {
		return CreateInput{}, err
	}
	n, err := workspace.NewName(s.ToString())
	if err != nil {
		return CreateInput{}, err
	}
	id, err := workspace.GenerateID()
	if err != nil {
		return CreateInput{}, err
	}
	d := workspace.NewDetail(n, s)
	return CreateInput{
		workspace.NewWorkspace(id, d),
		aID,
	}, nil
}

type FindAllMembersInput struct {
	CurrentWorkspaceID workspace.ID
}

type CreateInvitation struct {
	InviteeEmail       string
	InviteeDisplayName string
}

type InviteMembersInput struct {
	InviteeEmail       account.Email
	InviteeDisplayName member.DisplayName
	CurrentWorkspaceID workspace.ID
	AccountID          account.ID
	Invitations        []CreateInvitation
}

type VerifyInvitationTokenInput struct {
	Token invitation.Token
}

type RevokeInvitationInput struct {
	AccountID          account.ID
	CurrentWorkspaceID workspace.ID
	InvitationID       invitation.ID
}

type ResendInvitationInput struct {
	AccountID          account.ID
	CurrentWorkspaceID workspace.ID
	InvitationID       invitation.ID
}

type FindAllInvitationInput struct {
	CurrentWorkspaceID workspace.ID
	IsAccepted         bool
}

type UpdateMemberRoleInput struct {
	AccountID account.ID
	MemberID  member.ID
	Role      member.Role
}

type UpdateWorkspaceInput struct {
	AccountID   account.ID
	WorkspaceID workspace.ID
	Name        workspace.Name
	Subdomain   workspace.Subdomain
}
