package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"slices"
)

type Me struct {
	self             *user.User
	workspace        *workspace.Workspace
	member           *member.Member
	joinedWorkspaces workspace.Workspaces
}

func NewMe(self *user.User, workspace *workspace.Workspace, member *member.Member, joinedWorkspaces workspace.Workspaces) *Me {
	return &Me{self, workspace, member, joinedWorkspaces}
}

func (m *Me) Self() *user.User {
	return m.self
}

func (m *Me) Workspace() *workspace.Workspace {
	return m.workspace
}

func (m *Me) HasWorkspace() bool {
	return m.workspace != nil
}

func (m *Me) NotJoined() bool {
	return m.joinedWorkspaces == nil
}

func (m *Me) Member() *member.Member {
	return m.member
}

func (m *Me) HasMember() bool {
	return m.member != nil
}

func (m *Me) JoinedWorkspaces() workspace.Workspaces {
	return m.joinedWorkspaces
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
	m.UpdateSelf(m.self.UpdateName(name))
	return m
}
