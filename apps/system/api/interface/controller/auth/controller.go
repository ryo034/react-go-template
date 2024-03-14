package auth

import (
	"context"
	"fmt"

	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"github.com/ryo034/react-go-template/apps/system/api/driver/firebase"

	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	authUc "github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
)

type Controller interface {
	AuthByOTP(ctx context.Context, req *openapi.APIV1AuthByOtpReq) (openapi.APIV1AuthByOtpRes, error)
	VerifyOTP(ctx context.Context, req *openapi.APIV1VerifyOTPReq) (openapi.APIV1VerifyOTPRes, error)
	AuthByOAuth(ctx context.Context) (openapi.APIV1AuthByOAuthRes, error)
	APIV1ProcessInvitationEmail(ctx context.Context, i ProcessInvitationInput) (openapi.APIV1ProcessInvitationEmailRes, error)
	APIV1ProcessInvitationOAuth(ctx context.Context, i APIV1ProcessInvitationOAuth) (openapi.APIV1ProcessInvitationOAuthRes, error)
	InvitationByToken(ctx context.Context, i InvitationByTokenInput) (openapi.APIV1GetInvitationByTokenRes, error)
}

type controller struct {
	auc  authUc.UseCase
	resl shared.Resolver
	co   infraShared.ContextOperator
	fbDr firebase.Driver
}

type ProcessInvitationInput struct {
	Token uuid.UUID
	Email string
}

type APIV1ProcessInvitationOAuth struct {
	Token uuid.UUID
}

type InvitationByTokenInput struct {
	Token uuid.UUID
}

func NewController(auc authUc.UseCase, resl shared.Resolver, co infraShared.ContextOperator, fbDr firebase.Driver) Controller {
	return &controller{auc, resl, co, fbDr}
}

func (c *controller) AuthByOTP(ctx context.Context, req *openapi.APIV1AuthByOtpReq) (openapi.APIV1AuthByOtpRes, error) {
	em, err := account.NewEmail(req.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthByOtpRes), nil
	}
	inp := authUc.ByOTPInput{Email: em}
	res, err := c.auc.AuthByOTP(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthByOtpRes), nil
	}
	return res, nil
}

func (c *controller) VerifyOTP(ctx context.Context, req *openapi.APIV1VerifyOTPReq) (openapi.APIV1VerifyOTPRes, error) {
	em, err := account.NewEmail(req.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1VerifyOTPRes), nil
	}
	inp := authUc.VerifyOTPInput{Email: em, Otp: req.Otp}
	return c.auc.VerifyOTP(ctx, inp)
}

func (c *controller) AuthByOAuth(ctx context.Context) (openapi.APIV1AuthByOAuthRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthByOAuthRes), nil
	}
	ci, err := c.createUser(ctx, aID)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthByOAuthRes), nil
	}
	inp := authUc.ByOAuthInput{AccountID: aID, CreateInfo: ci}
	res, err := c.auc.AuthByOAuth(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthByOAuthRes), nil
	}
	return res, nil
}

func (c *controller) APIV1ProcessInvitationEmail(ctx context.Context, i ProcessInvitationInput) (openapi.APIV1ProcessInvitationEmailRes, error) {
	em, err := account.NewEmail(i.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1ProcessInvitationEmailRes), nil
	}
	inp := authUc.APIV1ProcessInvitationEmailInput{Token: invitation.NewToken(i.Token), Email: em}
	res, err := c.auc.APIV1ProcessInvitationEmail(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1ProcessInvitationEmailRes), nil
	}
	return res, nil
}

func (c *controller) createUser(ctx context.Context, aID account.ID) (authUc.CreateInfo, error) {
	prov, err := c.fbDr.GenProviderData(ctx)
	if err != nil {
		return authUc.CreateInfo{}, err
	}
	pi, err := c.fbDr.GetProviderInfo(ctx, firebase.GetProviderInfoRequiredOption{CurrentWorkspaceID: false})
	if err != nil {
		return authUc.CreateInfo{}, err
	}

	var na *account.Name = nil
	if pi.UserInfo.DisplayName == nil {
		// if not match name format, just ignore
		fmt.Printf("failed to create name: %s", err)
	} else {
		tmpNa, err := account.NewName(pi.UserInfo.DisplayName.ToString())
		if err != nil {
			// if not match name format, just ignore
			fmt.Printf("failed to create name: %s", err)
		}
		na = &tmpNa
	}

	if pi.UserInfo.Email == nil {
		return authUc.CreateInfo{}, fmt.Errorf("email is required")
	}
	em, err := account.NewEmail(pi.UserInfo.Email.ToString())
	if err != nil {
		return authUc.CreateInfo{}, err
	}

	var pho *user.Photo = nil
	if pi.UserInfo.Photo != nil {
		pho = user.NewPhotoFromFirebase(pi.UserInfo.Photo.URL())
	}

	return authUc.CreateInfo{User: user.NewUser(aID, em, na, pi.UserInfo.PhoneNumber, pho), Provider: prov}, nil
}

func (c *controller) APIV1ProcessInvitationOAuth(ctx context.Context, i APIV1ProcessInvitationOAuth) (openapi.APIV1ProcessInvitationOAuthRes, error) {
	pi, err := c.fbDr.GetProviderInfo(ctx, firebase.GetProviderInfoRequiredOption{CurrentWorkspaceID: false})
	if err != nil {
		return nil, err
	}
	var ci *authUc.CreateInfo = nil
	if pi.CustomClaim.AccountID != nil {
		tmpCI, err := c.createUser(ctx, *pi.CustomClaim.AccountID)
		if err != nil {
			return nil, err
		}
		ci = &tmpCI
	}

	if pi.UserInfo.Email == nil {
		return c.resl.Error(ctx, fmt.Errorf("email is required")).(openapi.APIV1ProcessInvitationOAuthRes), nil
	}
	inp := authUc.APIV1ProcessInvitationOAuthInput{
		Token:      invitation.NewToken(i.Token),
		Email:      *pi.UserInfo.Email,
		CreateInfo: ci,
	}
	res, err := c.auc.APIV1ProcessInvitationOAuth(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1ProcessInvitationOAuthRes), nil
	}
	return res, nil
}

func (c *controller) InvitationByToken(ctx context.Context, i InvitationByTokenInput) (openapi.APIV1GetInvitationByTokenRes, error) {
	inp := authUc.InvitationByTokenInput{Token: invitation.NewToken(i.Token)}
	res, err := c.auc.InvitationByToken(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1GetInvitationByTokenRes), nil
	}
	return res, nil
}
