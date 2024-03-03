//go:generate gomockhandler -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package core

import (
	"context"
	"database/sql"

	"github.com/ryo034/react-go-template/apps/system/api/util/reflect/function"
	"github.com/uptrace/bun"
)

type TransactionProvider interface {
	Provide(ctx context.Context) (TransactionResult, error)
}

type transactionProvider struct {
	primary *bun.DB
}

func NewTransactionProvider(primary *bun.DB) TransactionProvider {
	return &transactionProvider{primary}
}

func (p *transactionProvider) Provide(ctx context.Context) (TransactionResult, error) {
	tx, err := p.primary.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	return &transactionResult{ctx, tx}, nil
}

type TransactionResult interface {
	context.Context
	Transactional(fn interface{}) function.AnyFunc
}

type transactionResult struct {
	context.Context
	bun.Tx
}

func (r *transactionResult) Transactional(fn interface{}) function.AnyFunc {
	return Decorate(fn, r)
}
