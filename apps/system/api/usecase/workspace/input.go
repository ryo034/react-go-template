package workspace

import (
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
)

type CreateInput struct {
	Workspace *workspace.Workspace
	AccountID account.ID
}

func NewCreateInput(subdomain string, aID account.ID) (*CreateInput, error) {
	s, err := workspace.NewSubdomain(subdomain)
	if err != nil {
		return nil, err
	}
	n, err := workspace.NewName(s.ToString())
	if err != nil {
		return nil, err
	}
	id, err := workspace.GenerateID()
	if err != nil {
		return nil, err
	}
	d := workspace.NewDetail(n, s)
	return &CreateInput{
		workspace.NewWorkspace(id, d),
		aID,
	}, nil
}

type FindAllMembersInput struct {
	accountID account.ID
}

func NewFindAllMembersInput(aID account.ID) *FindAllMembersInput {
	return &FindAllMembersInput{aID}
}

type InviteMembersInput struct {
	AccountID   account.ID
	Invitations invitation.Invitations
}

type VerifyInvitationTokenInput struct {
	Token uuid.UUID
}
