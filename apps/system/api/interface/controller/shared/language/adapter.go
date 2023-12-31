package language

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"golang.org/x/text/language"
)

type Adapter interface {
	Adapt(ctx context.Context) language.Tag
}

type adapter struct {
	co       shared.ContextOperator
	fallback language.Tag
}

func (a *adapter) Adapt(ctx context.Context) language.Tag {
	v, err := a.co.GetLang(ctx)
	if err != nil {
		return v
	}
	return a.fallback
}

func NewAdapter(fallback language.Tag, co shared.ContextOperator) Adapter {
	return &adapter{co, fallback}
}
