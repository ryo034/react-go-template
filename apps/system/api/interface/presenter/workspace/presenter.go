package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	memberPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	invitationPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

func NewPresenter(wa Adapter, inva invitationPresenter.Adapter, ma memberPresenter.Adapter) workspaceUc.OutputPort {
	return &presenter{wa, inva, ma}
}

type presenter struct {
	wa   Adapter
	inva invitationPresenter.Adapter
	ma   memberPresenter.Adapter
}

func (p *presenter) Create(w *workspace.Workspace) (openapi.APIV1WorkspacesPostRes, error) {
	res := p.wa.Adapt(w)
	return &openapi.CreateWorkspaceResponse{Workspace: res}, nil
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

func (p *presenter) FindAllMembers(ms member.Members) (openapi.APIV1MembersGetRes, error) {
	res := p.ma.AdaptAll(ms)
	return &openapi.MembersResponse{Members: res}, nil
}

func (p *presenter) RevokeInvitation(is invitation.Invitations) (openapi.RevokeInvitationRes, error) {
	rs, err := p.inva.AdaptAll(is)
	if err != nil {
		return nil, err
	}
	return &rs, err
}

func (p *presenter) ResendInvitation(i *invitation.Invitation) (openapi.ResendInvitationRes, error) {
	rs, err := p.inva.Adapt(i)
	if err != nil {
		return nil, err
	}
	return &rs, err
}

func (p *presenter) FindAllInvitation(is invitation.Invitations) (openapi.APIV1InvitationsGetRes, error) {
	rs, err := p.inva.AdaptAll(is)
	if err != nil {
		return nil, err
	}
	return &openapi.InvitationsResponse{Invitations: rs}, err
}

func (p *presenter) UpdateMemberRole(m *member.Member) (openapi.APIV1MembersMemberIdRolePutRes, error) {
	res := p.ma.Adapt(m)
	return &openapi.UpdateMemberRoleResponse{Member: res}, nil
}
