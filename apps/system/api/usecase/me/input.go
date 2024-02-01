package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
)

type UpdateNameInput struct {
	accountID account.ID
	name      account.Name
}

func NewUpdateNameInput(aID account.ID, name string) (*UpdateNameInput, error) {
	n, err := account.NewName(name)
	if err != nil {
		return nil, err
	}
	return &UpdateNameInput{
		accountID: aID,
		name:      n,
	}, nil
}
