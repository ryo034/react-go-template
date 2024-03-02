package me

import (
	"slices"

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

func (m *Me) CheckJoined(wID workspace.ID) bool {
	ids := make([]workspace.ID, m.joinedWorkspaces.Size())
	for i, w := range m.joinedWorkspaces.AsSlice() {
		ids[i] = w.ID()
	}
	return slices.Contains(ids, wID)
}

func (m *Me) CheckNotJoined(wID workspace.ID) bool {
	return !m.CheckJoined(wID)
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
