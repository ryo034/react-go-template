//go:generate gomockhandler -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package shared

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/spf13/cast"
	"golang.org/x/text/language"
)

type ContextOperator interface {
	SetClaim(ctx context.Context, claim map[string]interface{}) context.Context
	SetUID(ctx context.Context, uID string) context.Context
	GetUID(ctx context.Context) (account.ID, error)
	SetAuthProviderUID(ctx context.Context, apUID provider.UID) context.Context
	GetAuthProviderUID(ctx context.Context) (provider.UID, error)
	SetRole(ctx context.Context, role member.Role) context.Context
	GetRole(ctx context.Context) (member.Role, error)
	GetUIDWithNil(ctx context.Context) (*account.ID, error)
	GetLang(ctx context.Context) (language.Tag, error)
	SetLang(ctx context.Context, lang language.Tag) context.Context
}

type contextOperator struct{}

func NewContextOperator() ContextOperator {
	return contextOperator{}
}

type Key string

const (
	ContextTokenKey           Key = "token"
	ContextUIDKey             Key = "uid"
	ContextAuthProviderUIDKey Key = "auth-provider-uid"
	ContextRoleKey            Key = "role"
	ContextLanguageKey        Key = "lang"
	ContextRequestIDKey       Key = "request-id"
)

func (co contextOperator) SetClaim(ctx context.Context, claim map[string]interface{}) context.Context {
	return context.WithValue(ctx, ContextTokenKey, claim)
}

func (co contextOperator) SetUID(ctx context.Context, uID string) context.Context {
	return context.WithValue(ctx, ContextUIDKey, uID)
}

func (co contextOperator) GetUID(ctx context.Context) (account.ID, error) {
	value := ctx.Value(ContextUIDKey)
	if value == nil {
		return account.ID{}, domainErr.NewUnauthenticated("uid is not found")
	}
	return account.NewID(value.(string))
}

func (co contextOperator) GetUIDWithNil(ctx context.Context) (*account.ID, error) {
	value := ctx.Value(ContextUIDKey)
	if value == nil {
		return nil, nil
	}
	id, err := account.NewID(value.(string))
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (co contextOperator) SetAuthProviderUID(ctx context.Context, apUID provider.UID) context.Context {
	return context.WithValue(ctx, ContextAuthProviderUIDKey, apUID.ToString())
}

func (co contextOperator) GetAuthProviderUID(ctx context.Context) (provider.UID, error) {
	value := ctx.Value(ContextAuthProviderUIDKey)
	v := cast.ToString(value)
	if value == nil || v == "" {
		return provider.UID{}, domainErr.NewUnauthenticated("auth provider uid is not found")
	}
	return provider.NewUID(v)
}

func (co contextOperator) SetRole(ctx context.Context, role member.Role) context.Context {
	return context.WithValue(ctx, ContextRoleKey, role.ToString())
}

func (co contextOperator) GetRole(ctx context.Context) (member.Role, error) {
	value := ctx.Value(ContextRoleKey)
	v := cast.ToString(value)
	if value == nil || v == "" {
		return "", domainErr.NewNoSuchData("role")
	}
	return member.NewRole(v)
}

func (co contextOperator) GetLang(ctx context.Context) (language.Tag, error) {
	v := ctx.Value(ContextLanguageKey)
	if v == nil {
		return language.Parse(language.Japanese.String())
	}
	return language.Parse(cast.ToString(v))
}

func (co contextOperator) SetLang(ctx context.Context, lang language.Tag) context.Context {
	return context.WithValue(ctx, ContextLanguageKey, lang)
}
