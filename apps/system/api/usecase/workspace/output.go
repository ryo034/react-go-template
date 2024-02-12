package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	Create(w *workspace.Workspace) *openapi.Workspace
	FindAllMembers(ms member.Members) *openapi.Members
	InviteMembers(ms member.InvitedMembers, registeredMembers member.InvitedMembers, failedMembers member.InvitedMembers) *openapi.BulkInvitedResult
	VerifyInvitationToken(m *member.InvitedMember) openapi.VerifyInvitationRes
}
