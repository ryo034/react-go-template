package me

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
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

func (g *gateway) Find(ctx context.Context, exec bun.IDB, mID member.ID) (*me.Me, error) {
	res, err := g.md.Find(ctx, exec, mID)
	if err != nil {
		return nil, err
	}
	ws, err := g.wd.FindAll(ctx, exec, account.NewIDFromUUID(res.SystemAccountID))
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
	return g.Find(ctx, exec, member.NewIDFromUUID(res.Member.MemberID))
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

func (g *gateway) FindProfile(ctx context.Context, exec bun.IDB, aID account.ID) (*me.Me, error) {
	res, err := g.md.FindProfile(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	return g.a.AdaptSystemAccount(res)
}

func (g *gateway) FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*me.Me, error) {
	res, err := g.md.FindByEmail(ctx, exec, email)
	if err != nil {
		return nil, err
	}
	return g.a.AdaptSystemAccount(res)
}

func (g *gateway) UpdateProfile(ctx context.Context, exec bun.IDB, usr *user.User) error {
	if err := g.md.UpdateProfile(ctx, exec, usr); err != nil {
		return err
	}
	return g.fd.UpdateProfile(ctx, usr)
}

func (g *gateway) UpdateMember(ctx context.Context, exec bun.IDB, m *me.Me) error {
	return g.md.UpdateMember(ctx, exec, m)
}

func (g *gateway) AcceptInvitation(ctx context.Context, exec bun.IDB, id invitation.ID) error {
	return g.md.AcceptInvitation(ctx, exec, id)
}
