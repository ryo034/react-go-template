package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
)

type Repository interface {
	Find(ctx context.Context, aID account.ID) (*Me, error)
}
