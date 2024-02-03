package me

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	meDr "github.com/ryo034/react-go-template/apps/system/api/driver/me"
	workspaceDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace"
	"github.com/uptrace/bun"
)

type gateway struct {
	md meDr.Driver
	fd fbDr.Driver
	wd workspaceDr.Driver
	a  Adapter
}

func NewGateway(md meDr.Driver, fd fbDr.Driver, wd workspaceDr.Driver, a Adapter) me.Repository {
	return &gateway{md, fd, wd, a}
}

func (g *gateway) Find(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*me.Me, error) {
	res, err := g.md.Find(ctx, exec, aID, wID)
	if err != nil {
		return nil, err
	}
	ws, err := g.wd.FindAll(ctx, exec, aID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return g.a.Adapt(res, ws)
}

func (g *gateway) FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*me.Me, error) {
	res, err := g.md.FindLastLogin(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	return g.Find(ctx, exec, aID, workspace.NewIDFromUUID(res.Member.WorkspaceID))
}

func (g *gateway) LastLogin(ctx context.Context, exec bun.IDB, m *me.Me) error {
	if err := g.md.LastLogin(ctx, exec, m.Member().ID()); err != nil {
		return err
	}
	return g.fd.SetCurrentWorkspaceToCustomClaim(ctx, m.Self().AccountID(), m.Workspace().ID())
}

func (g *gateway) FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*me.Me, error) {
	res, err := g.md.FindBeforeOnboard(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	return g.a.AdaptSystemAccount(res)
}

func (g *gateway) Update(ctx context.Context, exec bun.IDB, m *me.Me) (*me.Me, error) {
	res, err := g.md.Update(ctx, exec, m)
	if err != nil {
		return nil, err
	}
	return g.a.Adapt(res, nil)
}

func (g *gateway) UpdateName(ctx context.Context, exec bun.IDB, aID account.ID, name account.Name) (*me.Me, error) {
	res, err := g.md.UpdateName(ctx, exec, aID, name)
	if err != nil {
		return nil, err
	}
	if err = g.fd.UpdateName(ctx, aID, name); err != nil {
		return nil, err
	}
	return g.a.AdaptSystemAccount(res)
}
