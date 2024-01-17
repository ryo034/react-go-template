package shared

import (
	"context"
	"golang.org/x/text/language"
)

type LanguageAdapter interface {
	Adapt(ctx context.Context) language.Tag
}

type languageAdapter struct {
	fallback language.Tag
}

func (l *languageAdapter) Adapt(ctx context.Context) language.Tag {
	v := ctx.Value("language")
	if result, ok := v.(language.Tag); ok {
		return result
	}
	return l.fallback
}

func NewLanguageAdapter(fallback language.Tag) LanguageAdapter {
	return &languageAdapter{fallback}
}
