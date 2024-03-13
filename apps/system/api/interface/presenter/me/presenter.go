package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type presenter struct {
	a Adapter
}

func NewPresenter(a Adapter) meUc.OutputPort {
	return &presenter{a}
}

func (p *presenter) Find(m *me.Me) (openapi.APIV1MeGetRes, error) {
	am, err := p.a.Adapt(m)
	if err != nil {
		return nil, err
	}
	return &openapi.MeResponse{Me: am}, nil
}

func (p *presenter) AcceptInvitation(m *me.Me) (openapi.AcceptInvitationRes, error) {
	am, err := p.a.Adapt(m)
	if err != nil {
		return nil, err
	}
	return &openapi.InvitationsAcceptResponse{Me: am}, nil
}

func (p *presenter) UpdateProfile(m *me.Me) (openapi.APIV1MeProfilePutRes, error) {
	am, err := p.a.Adapt(m)
	if err != nil {
		return nil, err
	}
	return &openapi.UpdateProfileResponse{Me: am}, nil
}

func (p *presenter) UpdateMemberProfile(m *me.Me) (openapi.APIV1MeMemberProfilePutRes, error) {
	am, err := p.a.Adapt(m)
	if err != nil {
		return nil, err
	}
	return &openapi.UpdateMeMemberProfileResponse{Me: am}, nil
}

func (p *presenter) UpdateProfilePhoto(m *me.Me) (openapi.APIV1MeProfilePhotoPutRes, error) {
	am, err := p.a.Adapt(m)
	if err != nil {
		return nil, err
	}
	return &openapi.UpdateProfilePhotoResponse{Me: am}, nil
}

func (p *presenter) RemoveProfilePhoto(m *me.Me) (openapi.APIV1MeProfilePhotoDeleteRes, error) {
	am, err := p.a.Adapt(m)
	if err != nil {
		return nil, err
	}
	return &openapi.RemoveProfilePhotoResponse{Me: am}, nil
}

func (p *presenter) LeaveWorkspace() (openapi.APIV1MeWorkspaceLeavePostRes, error) {
	return &openapi.APIV1MeWorkspaceLeavePostNoContent{}, nil
}
