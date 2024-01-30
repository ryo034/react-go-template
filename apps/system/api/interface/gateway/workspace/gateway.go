package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	workspaceDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace"
	"github.com/uptrace/bun"
)

type gateway struct {
	d   workspaceDr.Driver
	adp Adapter
}

func NewGateway(d workspaceDr.Driver, adp Adapter) workspace.Repository {
	return &gateway{d, adp}
}

func (g *gateway) FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (workspace.Workspaces, error) {
	res, err := g.d.FindAll(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	return g.adp.AdaptAll(res)
}
