package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	Create(w *workspace.Workspace) *openapi.Workspace
	FindAllMembers(ms member.Members) *openapi.Members
	InviteMembers(is invitation.Invitations, registered invitation.Invitations, success invitation.Invitations, failed invitation.Invitations) (*openapi.InvitationsBulkResponse, error)
	VerifyInvitationToken(w *workspace.Workspace, i *invitation.Invitation) openapi.VerifyInvitationRes
}
