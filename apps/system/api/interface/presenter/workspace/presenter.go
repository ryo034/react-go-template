package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	memberPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

func NewPresenter(wa Adapter, ma memberPresenter.Adapter) workspaceUc.OutputPort {
	return &presenter{wa, ma}
}

type presenter struct {
	wa Adapter
	ma memberPresenter.Adapter
}

func (p *presenter) Create(w *workspace.Workspace) *openapi.Workspace {
	res := p.wa.Adapt(w)
	return &res
}

func (p *presenter) FindAllMembers(ms member.Members) *openapi.Members {
	res := p.ma.AdaptAll(ms)
	return &res
}

func (p *presenter) InviteMembers(ms invitation.Invitations, successList invitation.Invitations, registeredList invitation.Invitations, failedList invitation.Invitations) *openapi.InvitationsBulkResponse {
	rims := make([]openapi.Invitation, 0, registeredList.Size())
	for _, ivm := range registeredList.AsSlice() {
		rims = append(rims, openapi.Invitation{
			ID:           ivm.ID().Value(),
			Verified:     ivm.Verified(),
			ExpiredAt:    ivm.ExpiredAt().Value().ToTime(),
			InviteeEmail: ivm.InviteeEmail().ToString(),
			DisplayName:  ivm.DisplayName().ToString(),
		})
	}

	fims := make([]openapi.Invitation, 0, failedList.Size())
	for _, ivm := range failedList.AsSlice() {
		fims = append(fims, openapi.Invitation{
			ID:           ivm.ID().Value(),
			Verified:     ivm.Verified(),
			ExpiredAt:    ivm.ExpiredAt().Value().ToTime(),
			InviteeEmail: ivm.InviteeEmail().ToString(),
			DisplayName:  ivm.DisplayName().ToString(),
		})
	}

	sis := make([]openapi.Invitation, 0, successList.Size())
	for _, ivm := range successList.AsSlice() {
		sis = append(sis, openapi.Invitation{
			ID:           ivm.ID().Value(),
			Verified:     ivm.Verified(),
			ExpiredAt:    ivm.ExpiredAt().Value().ToTime(),
			InviteeEmail: ivm.InviteeEmail().ToString(),
			DisplayName:  ivm.DisplayName().ToString(),
		})
	}
	return &openapi.InvitationsBulkResponse{
		Total:                 ms.Size(),
		SuccessfulInvitations: sis,
		RegisteredInvitations: rims,
		FailedInvitations:     fims,
	}
}

func (p *presenter) VerifyInvitationToken(w *workspace.Workspace, i *invitation.Invitation) openapi.VerifyInvitationRes {
	d := w.Detail()
	return &openapi.InvitationInfoResponse{
		WorkspaceName: d.Name().ToString(),
		Verified:      i.Verified(),
	}
}
