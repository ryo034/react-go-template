package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (Workspaces, error)
	Create(ctx context.Context, exec bun.IDB, w *Workspace) (*Workspace, error)
	AddMember(ctx context.Context, exec bun.IDB, w *Workspace, m *member.Member) (*member.Member, error)
	FindMember(ctx context.Context, exec bun.IDB, aID account.ID, wID ID) (*member.Member, error)
}
