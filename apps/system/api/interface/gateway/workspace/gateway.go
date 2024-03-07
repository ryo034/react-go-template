package workspace

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	workspaceDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace"
	invitationDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace/invitation"
	memberDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace/member"
	memberGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
	invitationGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace/invitation"
	"github.com/uptrace/bun"
)

type gateway struct {
	d    workspaceDr.Driver
	md   memberDr.Driver
	invd invitationDr.Driver
	adp  Adapter
	ma   memberGw.Adapter
	ia   invitationGw.Adapter
}

func NewGateway(d workspaceDr.Driver, md memberDr.Driver, invd invitationDr.Driver, adp Adapter, ma memberGw.Adapter, ia invitationGw.Adapter) workspace.Repository {
	return &gateway{d, md, invd, adp, ma, ia}
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

func (g *gateway) UpdateMemberRole(ctx context.Context, exec bun.IDB, m *member.Member) (*member.Member, error) {
	return m, g.d.UpdateMemberRole(ctx, exec, m)
}

func (g *gateway) FindMember(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*member.Member, error) {
	res, err := g.d.FindMember(ctx, exec, aID, wID)
	if err != nil {
		return nil, err
	}
	return g.ma.Adapt(res)
}

func (g *gateway) FindAllMembers(ctx context.Context, exec bun.IDB, wID workspace.ID) (member.Members, error) {
	res, err := g.d.FindAllMembers(ctx, exec, wID)
	if err != nil {
		return nil, err
	}
	return g.ma.AdaptAll(res)
}

func (g *gateway) InviteMembers(ctx context.Context, exec bun.IDB, inviter workspace.Inviter, is invitation.Invitations) error {
	return g.d.InviteMembers(ctx, exec, inviter, is)
}

func (g *gateway) FindInviterFromToken(ctx context.Context, exec bun.IDB, token invitation.Token) (workspace.Inviter, error) {
	res, err := g.invd.FindActiveByToken(ctx, exec, token)
	if err != nil {
		return workspace.Inviter{}, err
	}
	return g.adp.AdaptInviter(res.InvitationUnit.Workspace, res.InvitationUnit.Member)
}

func (g *gateway) FindActiveInvitation(ctx context.Context, exec bun.IDB, id invitation.ID) (*invitation.Invitation, *workspace.Workspace, error) {
	res, err := g.invd.Find(ctx, exec, id)
	if err != nil {
		return nil, nil, err
	}
	im, err := g.ia.Adapt(res)
	if err != nil {
		return nil, nil, err
	}
	w, err := g.adp.Adapt(res.InvitationUnit.Workspace)
	if err != nil {
		return nil, nil, err
	}
	return im, w, nil
}

func (g *gateway) FindAllInvitations(ctx context.Context, exec bun.IDB, wID workspace.ID) (invitation.Invitations, error) {
	res, err := g.invd.FindAllByWorkspace(ctx, exec, wID)
	if err != nil {
		return nil, err
	}
	return g.ia.AdaptAll(res)
}
