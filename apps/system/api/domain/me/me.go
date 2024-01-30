package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type Me struct {
	self      *user.User
	workspace *workspace.Workspace
	member    *member.Member
}

func NewMe(self *user.User, workspace *workspace.Workspace, member *member.Member) *Me {
	return &Me{self, workspace, member}
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

func (m *Me) Member() *member.Member {
	return m.member
}

func (m *Me) HasMember() bool {
	return m.member != nil
}
