package workspace

import "github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"

type Inviter struct {
	*member.Member
	workspace *Workspace
}

func NewInviter(m *member.Member, w *Workspace) Inviter {
	return Inviter{m, w}
}

func (i *Inviter) Workspace() *Workspace {
	return i.workspace
}
