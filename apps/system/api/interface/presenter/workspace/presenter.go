package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	memberPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	invitation2 "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

func NewPresenter(wa Adapter, inva invitation2.Adapter, ma memberPresenter.Adapter) workspaceUc.OutputPort {
	return &presenter{wa, inva, ma}
}

type presenter struct {
	wa   Adapter
	inva invitation2.Adapter
	ma   memberPresenter.Adapter
}

func (p *presenter) Create(w *workspace.Workspace) *openapi.Workspace {
	res := p.wa.Adapt(w)
	return &res
}

func (p *presenter) FindAllMembers(ms member.Members) *openapi.Members {
	res := p.ma.AdaptAll(ms)
	return &res
}

func (p *presenter) InviteMembers(ms invitation.Invitations, registeredList invitation.Invitations, successList invitation.Invitations, failedList invitation.Invitations) (*openapi.InvitationsBulkResponse, error) {
	rims, err := p.inva.AdaptAll(registeredList)
	if err != nil {
		return nil, err
	}
	fims, err := p.inva.AdaptAll(failedList)
	if err != nil {
		return nil, err
	}
	sis, err := p.inva.AdaptAll(successList)
	if err != nil {
		return nil, err
	}
	return &openapi.InvitationsBulkResponse{
		Total:                 ms.Size(),
		SuccessfulInvitations: sis,
		RegisteredInvitations: rims,
		FailedInvitations:     fims,
	}, nil
}

func (p *presenter) VerifyInvitationToken(w *workspace.Workspace, i *invitation.Invitation) openapi.VerifyInvitationRes {
	d := w.Detail()
	return &openapi.InvitationInfoResponse{
		WorkspaceName: d.Name().ToString(),
		Verified:      i.IsVerified(),
	}
}
