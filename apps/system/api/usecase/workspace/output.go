package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	Create(w *workspace.Workspace) (openapi.APIV1WorkspacesPostRes, error)
	FindAllMembers(ms member.Members) (openapi.APIV1MembersGetRes, error)
	InviteMembers(is invitation.Invitations, registered invitation.Invitations, success invitation.Invitations, failed invitation.Invitations) (*openapi.InvitationsBulkResponse, error)
	RevokeInvitation(is invitation.Invitations) (openapi.RevokeInvitationRes, error)
	FindAllInvitation(is invitation.Invitations) (openapi.APIV1InvitationsGetRes, error)
}
