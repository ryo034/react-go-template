package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Controller interface {
	Find(ctx context.Context, aID string, wID string) (openapi.MeGetRes, error)
}

type controller struct {
	uc   meUc.UseCase
	resl shared.Resolver
}

func NewController(uc meUc.UseCase, resl shared.Resolver) Controller {
	return &controller{uc, resl}
}

func (c *controller) Find(ctx context.Context, aID string, wID string) (openapi.MeGetRes, error) {
	accountID, err := account.NewID(aID)
	wsID := workspace.ID{}
	if err != nil {
		return nil, err
	}
	res, err := c.uc.Find(ctx, accountID, wsID)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.MeGetRes), nil
	}
	return res, nil
}
