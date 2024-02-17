package invitation

import (
	"context"
	"github.com/google/uuid"
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

func (g *gateway) FindByToken(ctx context.Context, exec bun.IDB, token uuid.UUID) (*invitation.Invitation, error) {
	res, err := g.d.FindByToken(ctx, exec, token)
	if err != nil {
		return nil, err
	}
	return g.ia.Adapt(res)

}

func (g *gateway) VerifyByToken(ctx context.Context, exec bun.IDB, token uuid.UUID) error {
	return g.d.VerifyByToken(ctx, exec, token)
}
