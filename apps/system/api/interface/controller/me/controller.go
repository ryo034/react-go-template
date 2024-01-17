package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Controller interface {
	Find(ctx context.Context, aID string) (openapi.MeGetRes, error)
}

type controller struct {
	uc meUc.UseCase
}

func NewController(uc meUc.UseCase) Controller {
	return &controller{uc}
}

func (c *controller) Find(ctx context.Context, aID string) (openapi.MeGetRes, error) {
	accountID, err := account.NewID(aID)
	if err != nil {
		return nil, err
	}
	return c.uc.Find(ctx, accountID)
}
