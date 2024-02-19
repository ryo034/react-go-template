package me

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Controller interface {
	Find(ctx context.Context) (openapi.APIV1MeGetRes, error)
	UpdateProfile(ctx context.Context, i openapi.User) (openapi.APIV1MeProfilePutRes, error)
	AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.AcceptInvitationRes, error)
}

type controller struct {
	uc   meUc.UseCase
	resl shared.Resolver
	co   infraShared.ContextOperator
}

type AcceptInvitationInput struct {
	InvitationID uuid.UUID
}

func NewController(uc meUc.UseCase, resl shared.Resolver, co infraShared.ContextOperator) Controller {
	return &controller{uc, resl, co}
}

func (c *controller) Find(ctx context.Context) (openapi.APIV1MeGetRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := c.uc.Find(ctx, aID)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeGetRes), nil
	}
	return res, nil
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
	res, err := c.uc.UpdateProfile(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePutRes), nil
	}
	return res, nil
}

func (c *controller) AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.AcceptInvitationRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.AcceptInvitationRes), nil
	}
	in := meUc.AcceptInvitationInput{AccountID: aID, InvitationID: invitation.NewID(i.InvitationID)}
	res, err := c.uc.AcceptInvitation(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.AcceptInvitationRes), nil
	}
	return res, nil
}
