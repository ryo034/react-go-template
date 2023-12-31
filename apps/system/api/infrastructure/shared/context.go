package shared

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/spf13/cast"
	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
)

type ContextOperator interface {
	SetClaim(ctx context.Context, claim map[string]interface{}) context.Context
	SetUID(ctx context.Context, uID string) context.Context
	GetUID(ctx context.Context) (account.ID, error)
	GetLang(ctx context.Context) (language.Tag, error)
}

type contextOperator struct{}

func NewContextOperator() ContextOperator {
	return contextOperator{}
}

type key string

const (
	ContextTokenKey    key = "token"
	ContextUIDKey      key = "uid"
	ContextLanguageKey key = "lang"
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

func SetLang(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		panic("metadata.FromIncomingContext(ctx)")
	}
	langKey := md.Get(string(ContextLanguageKey))
	if len(langKey) == 0 {
		// Return Japanese if lang is not included in the request
		return context.WithValue(ctx, ContextLanguageKey, language.Japanese.String())
	}
	la := langKey[0]
	if la == "" {
		return context.WithValue(ctx, ContextLanguageKey, language.Japanese.String())
	}

	// Accept only Japanese / English
	switch la {
	case language.Japanese.String():
		la = language.Japanese.String()
	case language.English.String():
		la = language.English.String()
	}
	return context.WithValue(ctx, ContextLanguageKey, la)
}
