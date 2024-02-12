package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (Workspaces, error)
	Create(ctx context.Context, exec bun.IDB, w *Workspace) (*Workspace, error)
	AddMember(ctx context.Context, exec bun.IDB, w *Workspace, m *member.Member) (*member.Member, error)
	FindMember(ctx context.Context, exec bun.IDB, aID account.ID, wID ID) (*member.Member, error)
	FindAllMembers(ctx context.Context, exec bun.IDB, wID ID) (member.Members, error)
	BulkInviteMembers(ctx context.Context, exec bun.IDB, wID ID, ms member.InvitedMembers) error
	InviteMember(ctx context.Context, exec bun.IDB, wID ID, invitedBy member.InvitedBy, m *member.InvitedMember) error
	VerifyInvitedMember(ctx context.Context, exec bun.IDB, token uuid.UUID) (*member.InvitedMember, error)
}
