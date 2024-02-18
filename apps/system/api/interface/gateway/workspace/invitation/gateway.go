package invitation

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	invitationDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace/invitation"
	"github.com/uptrace/bun"
)

type gateway struct {
	d  invitationDr.Driver
	ia Adapter
}

func NewGateway(d invitationDr.Driver, ia Adapter) invitation.Repository {
	return &gateway{d, ia}
}

func (g *gateway) Find(ctx context.Context, exec bun.IDB, iID invitation.ID) (*invitation.Invitation, error) {
	res, err := g.d.Find(ctx, exec, iID)
	if err != nil {
		return nil, err
	}
	return g.ia.Adapt(res)
}

func (g *gateway) FindByToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*invitation.Invitation, error) {
	res, err := g.d.FindByToken(ctx, exec, token)
	if err != nil {
		return nil, err
	}
	return g.ia.Adapt(res)
}

func (g *gateway) FindActiveByToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*invitation.Invitation, error) {
	res, err := g.d.FindActiveByToken(ctx, exec, token)
	if err != nil {
		return nil, err
	}
	return g.ia.Adapt(res)
}

func (g *gateway) VerifyByToken(ctx context.Context, exec bun.IDB, token invitation.Token) error {
	return g.d.VerifyByToken(ctx, exec, token)
}

func (g *gateway) FindActiveByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*invitation.Invitation, error) {
	res, err := g.d.FindActiveByEmail(ctx, exec, email)
	if err != nil {
		return nil, err
	}
	return g.ia.Adapt(res)
}

func (g *gateway) FindActiveAllByEmail(ctx context.Context, exec bun.IDB, email account.Email) (invitation.Invitations, error) {
	res, err := g.d.FindActiveAllByEmail(ctx, exec, email)
	if err != nil {
		return nil, err
	}
	return g.ia.AdaptAll(res)
}

func (g *gateway) Accept(ctx context.Context, exec bun.IDB, iID invitation.ID) error {
	return g.d.Accept(ctx, exec, iID)
}
