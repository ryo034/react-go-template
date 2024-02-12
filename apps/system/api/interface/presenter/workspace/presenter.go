package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
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

func (p *presenter) InviteMembers(ms member.InvitedMembers, registeredMembers member.InvitedMembers, failedMembers member.InvitedMembers) *openapi.BulkInvitedResult {
	rims := make([]openapi.InvitedMember, 0, registeredMembers.Size())
	for _, ivm := range registeredMembers.AsSlice() {
		rims = append(rims, openapi.InvitedMember{
			Email: ivm.Email().ToString(),
		})
	}

	fims := make([]openapi.InvitedMember, 0, failedMembers.Size())
	for _, ivm := range failedMembers.AsSlice() {
		fims = append(fims, openapi.InvitedMember{
			Email: ivm.Email().ToString(),
		})
	}
	return &openapi.BulkInvitedResult{
		InvitedCount:      ms.Size(),
		RegisteredMembers: rims,
		FailedMembers:     fims,
	}
}

func (p *presenter) VerifyInvitationToken(m *member.InvitedMember) openapi.VerifyInvitationRes {
	return &openapi.InvitedMember{
		Email:    m.Email().ToString(),
		Verified: m.Verified(),
	}
}
