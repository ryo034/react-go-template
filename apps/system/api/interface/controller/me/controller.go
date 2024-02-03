package me

import (
	"context"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Controller interface {
	Find(ctx context.Context) (openapi.APIV1MeGetRes, error)
	UpdateName(ctx context.Context, i UpdateNameInput) (openapi.UpdateNameRes, error)
}

type controller struct {
	uc   meUc.UseCase
	resl shared.Resolver
	co   infraShared.ContextOperator
}

func NewController(uc meUc.UseCase, resl shared.Resolver, co infraShared.ContextOperator) Controller {
	return &controller{uc, resl, co}
}

func (c *controller) Find(ctx context.Context) (openapi.APIV1MeGetRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return nil, err
	}
	return c.uc.Find(ctx, aID)
}

type UpdateNameInput struct {
	Name string
}

func (c *controller) UpdateName(ctx context.Context, i UpdateNameInput) (openapi.UpdateNameRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.UpdateNameRes), nil
	}
	in, err := meUc.NewUpdateNameInput(aID, i.Name)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.UpdateNameRes), nil
	}
	return c.uc.UpdateName(ctx, in)
}
