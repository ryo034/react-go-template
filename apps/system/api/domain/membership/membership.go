package membership

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type Membership struct {
	workspace *workspace.Workspace
	member    *member.Member
	period    Period
}

func NewMembership() *Membership {
	return &Membership{}
}

func (m *Membership) Workspace() *workspace.Workspace {
	return m.workspace
}

func (m *Membership) Member() *member.Member {
	return m.member
}

func (m *Membership) Period() Period {
	return m.period
}
