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
	invitationDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace/invitation"
	"github.com/uptrace/bun"
)

type gateway struct {
	md   meDr.Driver
	fd   fbDr.Driver
	wd   workspaceDr.Driver
	invd invitationDr.Driver
	a    Adapter
}

func NewGateway(md meDr.Driver, fd fbDr.Driver, wd workspaceDr.Driver, invd invitationDr.Driver, a Adapter) me.Repository {
	return &gateway{md, fd, wd, invd, a}
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
	p, err := g.md.FindProfile(ctx, exec, account.NewIDFromUUID(res.SystemAccountID))
	if err != nil {
		return nil, err
	}
	em, err := account.NewEmail(p.Profile.Email)
	if err != nil {
		return nil, err
	}
	ris, err := g.invd.FindActiveAllByEmail(ctx, exec, em)
	if err != nil {
		return nil, err
	}
	return g.a.Adapt(res, ws, ris)
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
	m, err := g.a.AdaptSystemAccount(res)
	if err != nil {
		return nil, err
	}
	invs, err := g.invd.FindActiveAllByEmail(ctx, exec, m.Self().Email())
	if err != nil {
		return nil, err
	}
	ainvs, err := g.a.AdaptAllReceivedInvitation(invs)
	if err != nil {
		return nil, err
	}
	return m.UpdateReceivedInvitations(ainvs), nil
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
	return g.invd.Accept(ctx, exec, id)
}

func (g *gateway) FindAllActiveReceivedInvitations(ctx context.Context, exec bun.IDB, aID account.ID) (me.ReceivedInvitations, error) {
	p, err := g.md.FindProfile(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	em, err := account.NewEmail(p.Profile.Email)
	if err != nil {
		return nil, err
	}
	res, err := g.invd.FindActiveAllByEmail(ctx, exec, em)
	if err != nil {
		return nil, err
	}
	return g.a.AdaptAllReceivedInvitation(res)
}
