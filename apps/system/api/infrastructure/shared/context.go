package shared

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/spf13/cast"
	"golang.org/x/text/language"
)

type ContextOperator interface {
	SetClaim(ctx context.Context, claim map[string]interface{}) context.Context
	SetUID(ctx context.Context, uID string) context.Context
	GetUID(ctx context.Context) (account.ID, error)
	GetLang(ctx context.Context) (language.Tag, error)
	SetLang(ctx context.Context, lang language.Tag) context.Context
}

type contextOperator struct{}

func NewContextOperator() ContextOperator {
	return contextOperator{}
}

type Key string

const (
	ContextTokenKey     Key = "token"
	ContextUIDKey       Key = "uid"
	ContextLanguageKey  Key = "lang"
	ContextRequestIDKey Key = "request-id"
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
		return account.ID{}, domainErr.NewUnauthenticated()
	}
	return account.NewID(value.(string))
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
