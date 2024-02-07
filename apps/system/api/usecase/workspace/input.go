package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
)

type CreateInput struct {
	workspace *workspace.Workspace
	accountID account.ID
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
		workspace: workspace.NewWorkspace(id, d),
		accountID: aID,
	}, nil
}

func (i *CreateInput) Workspace() *workspace.Workspace {
	return i.workspace
}

func (i *CreateInput) AccountID() account.ID {
	return i.accountID
}
