package core

import (
	"context"
	"database/sql"

	"github.com/ryo034/react-go-template/packages/go/util/reflect/function"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ToExec(ctx context.Context, isRead bool) boil.ContextExecutor {
	if isRead {
		return readDB()
	}
	if result, ok := ctx.(*providedResult); ok {
		return result
	}
	return boil.GetContextDB()
}

type provider struct{}

func NewTransactionProvider() Provider {
	return &provider{}
}

func (i *provider) Provide(ctx context.Context) (ProvidedResult, error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &providedResult{ctx, tx}, nil
}

type providedResult struct {
	context.Context
	*sql.Tx
}

func (p *providedResult) Transactional(fn interface{}) function.AnyFunc {
	return Decorate(fn, p)
}

type Provider interface {
	Provide(ctx context.Context) (ProvidedResult, error)
}

type ProvidedResult interface {
	context.Context
	Transactional(fn interface{}) function.AnyFunc
}
