package me

import (
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type Me struct {
	self                *user.User
	workspace           *workspace.Workspace
	member              *member.Member
	joinedWorkspaces    workspace.Workspaces
	receivedInvitations ReceivedInvitations
	providers           provider.Providers
}

func NewMe(self *user.User, workspace *workspace.Workspace, member *member.Member, joinedWorkspaces workspace.Workspaces, receivedInvitations ReceivedInvitations, providers provider.Providers) *Me {
	return &Me{self, workspace, member, joinedWorkspaces, receivedInvitations, providers}
}

func (m *Me) Self() *user.User {
	return m.self
}

func (m *Me) Workspace() *workspace.Workspace {
	return m.workspace
}

func (m *Me) NotJoined() bool {
	return m.joinedWorkspaces == nil && m.member == nil && m.workspace == nil
}

func (m *Me) IsJoined() bool {
	return m.joinedWorkspaces != nil && m.member != nil && m.workspace != nil
}

func (m *Me) Member() *member.Member {
	return m.member
}

func (m *Me) UpdateMember(member *member.Member) *Me {
	m.member = member
	return m
}

func (m *Me) JoinedWorkspaces() workspace.Workspaces {
	return m.joinedWorkspaces
}

func (m *Me) ReceivedInvitations() ReceivedInvitations {
	return m.receivedInvitations
}

func (m *Me) UpdateReceivedInvitations(ris ReceivedInvitations) *Me {
	m.receivedInvitations = ris
	return m
}

func (m *Me) UpdateSelf(u *user.User) *Me {
	m.self = u
	return m
}

func (m *Me) UpdateName(name account.Name) *Me {
	tmpSelf := m.self.UpdateName(name)
	updated := m.UpdateSelf(tmpSelf)
	if m.member == nil {
		return updated
	}
	return m.UpdateMember(m.member.UpdateUser(tmpSelf))
}

func (m *Me) Providers() provider.Providers {
	return m.providers
}

func (m *Me) UpdateProfilePhoto(photo *user.Photo) *Me {
	return m.UpdateSelf(m.self.UpdateProfilePhoto(photo))
}

func (m *Me) RemoveProfilePhoto() *Me {
	return m.UpdateSelf(m.self.RemoveProfilePhoto())
}

func (m *Me) SameAs(t *Me) bool {
	return m.Self().AccountID().Value().String() == t.Self().AccountID().Value().String()
}

func (m *Me) ValidateCanUpdateWorkspace(wID workspace.ID) error {
	if m.Workspace().ID().ToString() != wID.ToString() {
		return domainErr.NewForbidden("Cannot update workspace")
	}
	ok := m.Member().Role().IsOwner() || m.Member().Role().IsAdmin()
	if !ok {
		return domainErr.NewForbidden("Can update only owner or admin")
	}
	return nil
}

func (m *Me) ValidateCanLeave() error {
	if m.NotJoined() {
		return domainErr.NewUnauthenticated("Not joined")
	}
	if m.Member().Role().IsOwner() {
		return domainErr.NewForbidden("Cannot leave owner")
	}
	if err := m.member.ValidateCanLeave(); err != nil {
		return err
	}
	return nil
}
