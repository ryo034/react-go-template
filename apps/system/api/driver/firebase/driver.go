package firebase

import (
	"context"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
)

type Driver interface {
	DeleteUser(ctx context.Context, aID account.ID) error
	RevokeRefreshTokens(ctx context.Context, aID account.ID) error
	GetUser(ctx context.Context, aID account.ID) (*auth.UserRecord, error)
	EmailVerified(ctx context.Context, aID account.ID) (bool, string, error)
	UpdateMe(ctx context.Context, me *me.Me) error
	UpdateEmail(ctx context.Context, aID account.ID, em account.Email) error
	UpdateName(ctx context.Context, aID account.ID, n account.Name) error
	UpdatePhoneNumber(ctx context.Context, aID account.ID, ph phone.Number) error
}

const currentStoreIdKey = "currentStoreID"

type driver struct {
	f *firebase.Firebase
}

func NewDriver(f *firebase.Firebase) Driver {
	return &driver{f}
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

func (d *driver) EmailVerified(ctx context.Context, aID account.ID) (bool, string, error) {
	ur, err := d.f.Auth.GetUser(ctx, aID.ToString())
	if err != nil {
		return false, "", err
	}
	return ur.EmailVerified, ur.Email, nil
}

func (d *driver) UpdateMe(ctx context.Context, me *me.Me) error {
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
