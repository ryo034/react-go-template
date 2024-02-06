package firebase

import (
	"context"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
)

type Driver interface {
	CustomToken(ctx context.Context, aID account.ID) (string, error)
	DeleteUser(ctx context.Context, aID account.ID) error
	RevokeRefreshTokens(ctx context.Context, aID account.ID) error
	GetUser(ctx context.Context, aID account.ID) (*auth.UserRecord, error)
	CreateUser(ctx context.Context, aID account.ID, email account.Email) error
	SetCurrentWorkspaceToCustomClaim(ctx context.Context, aID account.ID, wID workspace.ID) error
	UpdateProfile(ctx context.Context, usr *user.User) error
	UpdateEmail(ctx context.Context, aID account.ID, em account.Email) error
	UpdateName(ctx context.Context, aID account.ID, n account.Name) error
	UpdatePhoneNumber(ctx context.Context, aID account.ID, ph phone.Number) error
}

type driver struct {
	f *firebase.Firebase
}

func NewDriver(f *firebase.Firebase) Driver {
	return &driver{f}
}

func (d *driver) CustomToken(ctx context.Context, aID account.ID) (string, error) {
	return d.f.Auth.CustomToken(ctx, aID.ToString())
}

func (d *driver) DeleteUser(ctx context.Context, aID account.ID) error {
	return d.f.Auth.DeleteUser(ctx, aID.ToString())
}

func (d *driver) RevokeRefreshTokens(ctx context.Context, aID account.ID) error {
	return d.f.Auth.RevokeRefreshTokens(ctx, aID.ToString())
}

func (d *driver) GetUser(ctx context.Context, aID account.ID) (*auth.UserRecord, error) {
	ur, err := d.f.Auth.GetUser(ctx, aID.ToString())
	if err != nil {
		log.Fatalf("Error get firebase userID: %v\n error :%v", aID.ToString(), err)
		return nil, err
	}
	return ur, err
}

func (d *driver) UpdateProfile(ctx context.Context, usr *user.User) error {
	params := &auth.UserToUpdate{}
	if usr.HasName() {
		params = params.DisplayName(usr.Name().ToString())
	}
	if usr.HasPhoneNumber() {
		params = params.PhoneNumber(usr.PhoneNumber().ToInternationalNumberString())
	}
	_, err := d.f.Auth.UpdateUser(ctx, usr.AccountID().ToString(), params)
	return err
}

func (d *driver) SetCurrentWorkspaceToCustomClaim(ctx context.Context, aID account.ID, wID workspace.ID) error {
	return d.f.Auth.SetCustomUserClaims(ctx, aID.ToString(), map[string]interface{}{"current_workspace_id": wID.ToFriendlyString()})
}

func (d *driver) CreateUser(ctx context.Context, aID account.ID, email account.Email) error {
	params := (&auth.UserToCreate{}).Email(email.ToString()).UID(aID.ToString()).EmailVerified(true)
	if _, err := d.f.Auth.CreateUser(ctx, params); err != nil {
		return err
	}
	return nil
}

func (d *driver) UpdateEmail(ctx context.Context, aID account.ID, em account.Email) error {
	params := (&auth.UserToUpdate{}).Email(em.ToString())
	_, err := d.f.Auth.UpdateUser(ctx, aID.ToString(), params)
	if err != nil {
		if auth.IsEmailAlreadyExists(err) {
			return domainErr.NewEmailAlreadyInUse(em.ToString())
		}
		return domainErr.NewInvalidEmail(em.ToString())
	}
	return nil
}

func (d *driver) UpdateName(ctx context.Context, aID account.ID, n account.Name) error {
	params := (&auth.UserToUpdate{}).DisplayName(n.ToString())
	_, err := d.f.Auth.UpdateUser(ctx, aID.ToString(), params)
	return err
}

func (d *driver) UpdatePhoneNumber(ctx context.Context, aID account.ID, ph phone.Number) error {
	params := (&auth.UserToUpdate{}).PhoneNumber(ph.ToInternationalNumberString())
	_, err := d.f.Auth.UpdateUser(ctx, aID.ToString(), params)
	if err != nil {
		if auth.IsPhoneNumberAlreadyExists(err) {
			return domainErr.NewPhoneNumberAlreadyInUse(ph.ToString())
		}
		return err
	}
	return nil
}
