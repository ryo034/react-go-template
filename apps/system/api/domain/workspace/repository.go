package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (Workspaces, error)
}
