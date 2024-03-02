package firebase

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/google/uuid"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"

	"firebase.google.com/go/v4/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
)

type Driver interface {
	CustomToken(ctx context.Context) (string, error)
	FindProviderData(ctx context.Context) (*provider.Provider, error)
	RevokeRefreshTokens(ctx context.Context) error
	GetUser(ctx context.Context) (*auth.UserRecord, error)
	CreateUser(ctx context.Context, email account.Email) error
	SetCurrentWorkspaceToCustomClaim(ctx context.Context, wID workspace.ID) error
	GetCurrentWorkspaceFromCustomClaim(ctx context.Context) (*workspace.ID, error)
	SetAccountIDToCustomClaim(ctx context.Context, aID account.ID) error
	MustGetCurrentWorkspaceFromCustomClaim(ctx context.Context) (workspace.ID, error)
	UpdateProfile(ctx context.Context, usr *user.User) error
	UpdateEmail(ctx context.Context, em account.Email) error
	UpdateName(ctx context.Context, n account.Name) error
	UpdatePhoneNumber(ctx context.Context, ph phone.Number) error
}

const (
	CustomClaimCurrentWorkspaceIDKey string = "current_workspace_id"
	CustomClaimAccountIDKey          string = "account_id"
)

type driver struct {
	f  *firebase.Firebase
	co shared.ContextOperator
}

func NewDriver(f *firebase.Firebase, co shared.ContextOperator) Driver {
	return &driver{f, co}
}

func (d *driver) CustomToken(ctx context.Context) (string, error) {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return "", err
	}
	return d.f.Auth.CustomToken(ctx, apUID.ToString())
}

func (d *driver) DeleteUser(ctx context.Context, apUID provider.UID) error {
	return d.f.Auth.DeleteUser(ctx, apUID.ToString())
}

func (d *driver) FindProviderData(ctx context.Context) (*provider.Provider, error) {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return nil, err
	}
	id, err := provider.GenerateID()
	if err != nil {
		return nil, err
	}
	return provider.NewProvider(id, provider.Google, provider.ProvidedByFirebase, apUID), nil
}

func (d *driver) RevokeRefreshTokens(ctx context.Context) error {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return err
	}
	return d.f.Auth.RevokeRefreshTokens(ctx, apUID.ToString())
}

func (d *driver) GetUser(ctx context.Context) (*auth.UserRecord, error) {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return nil, err
	}
	ur, err := d.f.Auth.GetUser(ctx, apUID.ToString())
	if err != nil {
		log.Printf("Error GetUser from authentication provider UID: %v\n error :%v", apUID, err)
		return nil, err
	}
	return ur, err
}

func (d *driver) UpdateProfile(ctx context.Context, usr *user.User) error {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return err
	}
	u, err := d.GetUser(ctx)
	if err != nil {
		return err
	}
	params := &auth.UserToUpdate{}
	if usr.HasName() {
		if u.DisplayName != usr.Name().ToString() {
			params = params.DisplayName(usr.Name().ToString())
		}
	}
	if usr.HasPhoneNumber() {
		if u.PhoneNumber != usr.PhoneNumber().ToE164() {
			params = params.PhoneNumber(usr.PhoneNumber().ToE164())
		}
	}

	// if params is empty, it will return nil
	if reflect.DeepEqual(params, &auth.UserToUpdate{}) {
		return nil
	}

	_, err = d.f.Auth.UpdateUser(ctx, apUID.ToString(), params)
	return err
}

func (d *driver) SetCurrentWorkspaceToCustomClaim(ctx context.Context, wID workspace.ID) error {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return err
	}
	u, err := d.f.Auth.GetUser(ctx, apUID.ToString())
	if err != nil {
		return err
	}
	cms := u.CustomClaims
	if cms == nil {
		cms = map[string]interface{}{}
	}
	cms[CustomClaimCurrentWorkspaceIDKey] = wID.Value()
	return d.f.Auth.SetCustomUserClaims(ctx, apUID.ToString(), cms)
}

func (d *driver) SetAccountIDToCustomClaim(ctx context.Context, aID account.ID) error {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return err
	}
	u, err := d.f.Auth.GetUser(ctx, apUID.ToString())
	if err != nil {
		return err
	}
	cms := u.CustomClaims
	if cms == nil {
		cms = map[string]interface{}{}
	}
	cms[CustomClaimAccountIDKey] = aID.Value().String()
	return d.f.Auth.SetCustomUserClaims(ctx, apUID.ToString(), cms)
}

func (d *driver) GetCurrentWorkspaceFromCustomClaim(ctx context.Context) (*workspace.ID, error) {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return nil, err
	}
	ur, err := d.f.Auth.GetUser(ctx, apUID.ToString())
	if err != nil {
		return nil, err
	}
	claims := ur.CustomClaims
	if claims == nil {
		return nil, nil
	}
	ccWID, ok := claims[CustomClaimCurrentWorkspaceIDKey].(string)
	if !ok {
		return nil, nil
	}
	wID := workspace.NewIDFromUUID(uuid.MustParse(ccWID))
	return &wID, err
}

func (d *driver) MustGetCurrentWorkspaceFromCustomClaim(ctx context.Context) (workspace.ID, error) {
	wID, err := d.GetCurrentWorkspaceFromCustomClaim(ctx)
	if err != nil {
		return workspace.ID{}, err
	}
	if wID == nil {
		return workspace.ID{}, domainErr.NewUnauthenticated(fmt.Sprintf("account does not have current workspace"))
	}
	return *wID, nil
}

func (d *driver) CreateUser(ctx context.Context, email account.Email) error {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return err
	}
	params := (&auth.UserToCreate{}).Email(email.ToString()).UID(apUID.ToString()).EmailVerified(true)
	if _, err := d.f.Auth.CreateUser(ctx, params); err != nil {
		return err
	}
	return nil
}

func (d *driver) UpdateEmail(ctx context.Context, em account.Email) error {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return err
	}
	params := (&auth.UserToUpdate{}).Email(em.ToString())
	_, err = d.f.Auth.UpdateUser(ctx, apUID.ToString(), params)
	if err != nil {
		if auth.IsEmailAlreadyExists(err) {
			return domainErr.NewEmailAlreadyInUse(em.ToString())
		}
		return domainErr.NewInvalidEmail(em.ToString())
	}
	return nil
}

func (d *driver) UpdateName(ctx context.Context, n account.Name) error {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return err
	}
	params := (&auth.UserToUpdate{}).DisplayName(n.ToString())
	_, err = d.f.Auth.UpdateUser(ctx, apUID.ToString(), params)
	return err
}

func (d *driver) UpdatePhoneNumber(ctx context.Context, ph phone.Number) error {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return err
	}
	params := (&auth.UserToUpdate{}).PhoneNumber(ph.ToE164())
	_, err = d.f.Auth.UpdateUser(ctx, apUID.ToString(), params)
	if err != nil {
		if auth.IsPhoneNumberAlreadyExists(err) {
			return domainErr.NewPhoneNumberAlreadyInUse(ph.ToE164())
		}
		return err
	}
	return nil
}
