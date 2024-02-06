package me

import (
	"context"
	"fmt"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Controller interface {
	Find(ctx context.Context) (openapi.APIV1MeGetRes, error)
	UpdateProfile(ctx context.Context, i openapi.User) (openapi.APIV1MeProfilePutRes, error)
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

func (c *controller) UpdateProfile(ctx context.Context, i openapi.User) (openapi.APIV1MeProfilePutRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePutRes), nil
	}
	if i.UserId != aID.Value() {
		// TODO: Return BadRequest
		return c.resl.Error(ctx, fmt.Errorf("Invalid Input")).(openapi.APIV1MeProfilePutRes), nil
	}
	in, err := meUc.NewUpdateProfileInput(i)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePutRes), nil
	}
	return c.uc.UpdateProfile(ctx, in)
}
