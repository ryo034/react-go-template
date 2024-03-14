package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type OutputPort interface {
	Create(w *workspace.Workspace) (openapi.APIV1CreateWorkspaceRes, error)
	FindAllMembers(ms member.Members) (openapi.APIV1GetMembersRes, error)
	InviteMembers(is invitation.Invitations, registered invitation.Invitations, success invitation.Invitations, failed invitation.Invitations) (*openapi.InvitationsBulkResponse, error)
	RevokeInvitation(is invitation.Invitations) (openapi.APIV1RevokeInvitationRes, error)
	ResendInvitation(i *invitation.Invitation) (openapi.APIV1ResendInvitationRes, error)
	FindAllInvitation(is invitation.Invitations) (openapi.APIV1GetInvitationsRes, error)
	UpdateMemberRole(m *member.Member) (openapi.APIV1UpdateMemberRoleRes, error)
	UpdateWorkspace(w *workspace.Workspace) (openapi.APIV1UpdateWorkspaceRes, error)
	Leave() (openapi.APIV1RemoveMemberRes, error)
}
