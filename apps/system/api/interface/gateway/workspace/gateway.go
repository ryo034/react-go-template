package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	workspaceDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace"
	memberDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace/member"
	memberGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
	"github.com/uptrace/bun"
)

type gateway struct {
	d   workspaceDr.Driver
	md  memberDr.Driver
	adp Adapter
	ma  memberGw.Adapter
}

func NewGateway(d workspaceDr.Driver, md memberDr.Driver, adp Adapter, ma memberGw.Adapter) workspace.Repository {
	return &gateway{d, md, adp, ma}
}

func (g *gateway) FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (workspace.Workspaces, error) {
	res, err := g.d.FindAll(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	return g.adp.AdaptAll(res)
}

func (g *gateway) Create(ctx context.Context, exec bun.IDB, w *workspace.Workspace) (*workspace.Workspace, error) {
	res, err := g.d.Create(ctx, exec, w)
	if err != nil {
		return nil, err
	}
	return g.adp.Adapt(res)
}

func (g *gateway) AddMember(ctx context.Context, exec bun.IDB, w *workspace.Workspace, m *member.Member) (*member.Member, error) {
	_, err := g.d.AddMember(ctx, exec, w, m)
	if err != nil {
		return nil, err
	}
	res, err := g.md.Find(ctx, exec, m.ID())
	if err != nil {
		return nil, err
	}
	return g.ma.Adapt(res)
}