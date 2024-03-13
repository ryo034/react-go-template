package firebase

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"reflect"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me"

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
	VerifyIDToken(ctx context.Context, token string) (*auth.Token, error)
	CustomToken(ctx context.Context) (string, error)
	GenProviderData(ctx context.Context) (*provider.Provider, error)
	RevokeRefreshTokens(ctx context.Context) error

	CreateUser(ctx context.Context, email account.Email) error
	GetProviderInfo(ctx context.Context, option GetProviderInfoRequiredOption) (ProviderInfo, error)

	SetMeToCustomClaim(ctx context.Context, me *me.Me) error
	SetAccountIDToCustomClaim(ctx context.Context, aID account.ID) error
	ClearCustomClaim(ctx context.Context) error

	UpdateProfile(ctx context.Context, usr *user.User) error
	UpdateEmail(ctx context.Context, em account.Email) error
	UpdateName(ctx context.Context, n account.Name) error
	UpdatePhoneNumber(ctx context.Context, ph phone.Number) error
}

const (
	CustomClaimCurrentWorkspaceIDKey string = "current_workspace_id"
	CustomClaimAccountIDKey          string = "account_id"
	CustomClaimRoleKey               string = "role"
)

type CustomClaim struct {
	CurrentWorkspaceID *workspace.ID
	AccountID          *account.ID
	Role               *member.Role
}

type UserInfo struct {
	Email       *account.Email
	DisplayName *member.DisplayName
	PhoneNumber *phone.Number
	Photo       *user.Photo
}

type ProviderInfo struct {
	CustomClaim CustomClaim
	UserInfo    UserInfo
}

type driver struct {
	f  *firebase.Firebase
	co shared.ContextOperator
	sh storage.Handler
}

func NewDriver(f *firebase.Firebase, co shared.ContextOperator, sh storage.Handler) Driver {
	return &driver{f, co, sh}
}

func (d *driver) VerifyIDToken(ctx context.Context, token string) (*auth.Token, error) {
	return d.f.Auth.VerifyIDToken(ctx, token)
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

func (d *driver) GenProviderData(ctx context.Context) (*provider.Provider, error) {
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

type GetProviderInfoRequiredOption struct {
	CurrentWorkspaceID bool
}

func (d *driver) GetProviderInfo(ctx context.Context, option GetProviderInfoRequiredOption) (ProviderInfo, error) {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return ProviderInfo{}, err
	}
	u, err := d.f.Auth.GetUser(ctx, apUID.ToString())
	if err != nil {
		return ProviderInfo{}, err
	}
	claims := u.CustomClaims
	if claims == nil {
		claims = map[string]interface{}{}
	}
	var clmAID *account.ID = nil
	var clmWID *workspace.ID = nil
	var clmRole *member.Role = nil

	aIDInClaims, _ := claims[CustomClaimAccountIDKey].(string)
	if aIDInClaims != "" {
		aID := account.NewIDFromUUID(uuid.MustParse(aIDInClaims))
		if err != nil {
			return ProviderInfo{}, err
		}
		clmAID = &aID
	}

	wIDInClaims, _ := claims[CustomClaimCurrentWorkspaceIDKey].(string)
	if wIDInClaims == "" && option.CurrentWorkspaceID {
		return ProviderInfo{}, domainErr.NewUnauthenticated(fmt.Sprintf("account does not have current workspace"))
	}
	if wIDInClaims != "" {
		wID := workspace.NewIDFromUUID(uuid.MustParse(wIDInClaims))
		clmWID = &wID
	}

	roleInClaims, _ := claims[CustomClaimRoleKey].(string)
	if roleInClaims != "" {
		role, err := member.NewRole(roleInClaims)
		if err != nil {
			return ProviderInfo{}, err
		}
		clmRole = &role
	}

	var em *account.Email = nil
	var dn *member.DisplayName = nil
	var ph *phone.Number = nil
	var pho *user.Photo = nil

	if u.UserInfo.Email != "" {
		tmpEm, err := account.NewEmail(u.Email)
		if err != nil {
			return ProviderInfo{}, err
		}
		em = &tmpEm
	}

	if u.UserInfo.DisplayName != "" {
		dn = member.NewDisplayName(u.DisplayName)
	}

	if u.UserInfo.PhoneNumber != "" {
		tmpPh, err := phone.NewInternationalPhoneNumber(u.PhoneNumber, "JP")
		if err != nil {
			return ProviderInfo{}, err
		}
		ph = &tmpPh
	}

	if u.UserInfo.PhotoURL != "" {
		phoID, err := uuid.NewV7()
		if err != nil {
			return ProviderInfo{}, err
		}
		uri, err := url.Parse(u.UserInfo.PhotoURL)
		if err != nil {
			return ProviderInfo{}, err
		}
		pho = user.NewPhoto(media.NewIDFromUUID(phoID), media.HostingToFirebase, uri)
	}

	return ProviderInfo{
		CustomClaim: CustomClaim{clmWID, clmAID, clmRole},
		UserInfo:    UserInfo{em, dn, ph, pho},
	}, nil

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
	if usr.HasPhoto() {
		photoPath, err := d.sh.CreateAvatarPath(usr.Photo().HostingTo(), usr.AccountID(), usr.Photo().ID())
		if err != nil {
			return err
		}
		if u.PhotoURL != photoPath.String() {
			params = params.PhotoURL(photoPath.String())
		}
	}

	// if params is empty, it will return nil
	if reflect.DeepEqual(params, &auth.UserToUpdate{}) {
		return nil
	}

	_, err = d.f.Auth.UpdateUser(ctx, apUID.ToString(), params)
	return err
}

func (d *driver) GetEmail(ctx context.Context) (account.Email, error) {
	apUID, err := d.co.GetAuthProviderUID(ctx)
	if err != nil {
		return account.Email{}, err
	}
	u, err := d.f.Auth.GetUser(ctx, apUID.ToString())
	if err != nil {
		return account.Email{}, err
	}
	return account.NewEmail(u.Email)
}

func (d *driver) SetMeToCustomClaim(ctx context.Context, me *me.Me) error {
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
	cms[CustomClaimAccountIDKey] = me.Self().AccountID().Value().String()
	if me.IsJoined() {
		cms[CustomClaimCurrentWorkspaceIDKey] = me.Workspace().ID().Value().String()
		cms[CustomClaimRoleKey] = me.Member().Role().ToString()
	}
	return d.f.Auth.SetCustomUserClaims(ctx, apUID.ToString(), cms)
}

func (d *driver) ClearCustomClaim(ctx context.Context) error {
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
		return nil
	}
	delete(cms, CustomClaimCurrentWorkspaceIDKey)
	delete(cms, CustomClaimAccountIDKey)
	delete(cms, CustomClaimRoleKey)
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
