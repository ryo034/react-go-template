//go:generate gomockhandler -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package workspace

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (Workspaces, error)
	Create(ctx context.Context, exec bun.IDB, w *Workspace) (*Workspace, error)
	Update(ctx context.Context, exec bun.IDB, w *Workspace) error
	AddMember(ctx context.Context, exec bun.IDB, w *Workspace, m *member.Member) (*member.Member, error)
	UpdateMemberRole(ctx context.Context, exec bun.IDB, assignor *member.Member, m *member.Member) (*member.Member, error)
	FindMember(ctx context.Context, exec bun.IDB, memID member.ID) (*member.Member, error)
	FindAllMembers(ctx context.Context, exec bun.IDB, wID ID) (member.Members, error)
	InviteMembers(ctx context.Context, exec bun.IDB, inviter Inviter, is invitation.Invitations) error
	FindInviterFromToken(ctx context.Context, exec bun.IDB, token invitation.Token) (Inviter, error)
	FindActiveInvitation(ctx context.Context, exec bun.IDB, id invitation.ID) (*invitation.Invitation, *Workspace, error)
	FindAllInvitations(ctx context.Context, exec bun.IDB, wID ID) (invitation.Invitations, error)
	Leave(ctx context.Context, exec bun.IDB, executorID member.ID, mID member.ID) error
}
