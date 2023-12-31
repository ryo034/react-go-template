package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
)

type repository struct {
	fd firebaseDriver.Driver
}

func NewRepository(fd firebaseDriver.Driver) me.Repository {
	return &repository{fd}
}

func (r *repository) Find(ctx context.Context, aID account.ID) (*me.Me, error) {
	return nil, nil
}
