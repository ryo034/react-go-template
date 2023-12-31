package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/sqlboiler/core"
)

type UseCase interface {
	SignUp(ctx context.Context, meID account.ID, firstName account.FirstName, lastName account.LastName) (*me.Me, error)
	Login(ctx context.Context, aID account.ID) (*me.Me, error)
	RegisterComplete(ctx context.Context, aID account.ID) (*me.Me, error)
	Find(ctx context.Context, aID account.ID, isRead bool) (*me.Me, error)
	Update(ctx context.Context, me *me.Me) error
	UpdateEmail(ctx context.Context, aID account.ID, em account.Email) (*me.Me, error)
	UpdateName(ctx context.Context, aID account.ID, fin account.FirstName, ln account.LastName) (*me.Me, error)
	UpdatePhoneNumber(ctx context.Context, aID account.ID, ph phone.Number) (*me.Me, error)
}

type useCase struct {
	isLocal bool
	txp     core.Provider
	repo    me.Repository
}

func NewUseCase(isLocal bool, txp core.Provider, acRepo me.Repository) UseCase {
	return &useCase{isLocal, txp, acRepo}
}

func (u *useCase) SignUp(ctx context.Context, meID account.ID, firstName account.FirstName, lastName account.LastName) (*me.Me, error) {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		return u.repo.SaveFromTemporary(pr, meID, firstName, lastName)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return result.Value(0).(*me.Me), nil
}

func (u *useCase) Login(ctx context.Context, aID account.ID) (*me.Me, error) {
	return u.repo.Find(ctx, core.ToExec(ctx, false), aID)
}

func (u *useCase) RegisterComplete(ctx context.Context, aID account.ID) (*me.Me, error) {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		verified, email, err := u.repo.EmailVerified(ctx, aID)
		if err != nil {
			return nil, err
		}
		if !verified {
			return nil, domainErr.NewEmailNotVerified(email)
		} else {
			if err = u.repo.VerifyEmail(pr, aID); err != nil {
				return nil, err
			}
		}
		return u.Find(pr, aID, false)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return result.Value(0).(*me.Me), nil
}

func (u *useCase) Find(ctx context.Context, aID account.ID, isRead bool) (*me.Me, error) {
	return u.repo.Find(ctx, core.ToExec(ctx, isRead), aID)
}

func (u *useCase) Update(ctx context.Context, me *me.Me) error {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return err
	}
	fn := func() (err error) {
		return u.repo.Update(pr, me)
	}
	result := pr.Transactional(fn)()
	return result.Error()
}

func (u *useCase) UpdateEmail(ctx context.Context, aID account.ID, em account.Email) (*me.Me, error) {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		m, err := u.repo.UpdateEmail(pr, aID, em)
		if err != nil {
			return nil, err
		}
		return m, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return result.Value(0).(*me.Me), nil
}

func (u *useCase) UpdateName(ctx context.Context, aID account.ID, fin account.FirstName, ln account.LastName) (*me.Me, error) {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		m, err := u.repo.UpdateName(pr, aID, fin, ln)
		if err != nil {
			return nil, err
		}
		return m, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return result.Value(0).(*me.Me), nil
}

func (u *useCase) UpdatePhoneNumber(ctx context.Context, aID account.ID, ph phone.Number) (*me.Me, error) {
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		m, err := u.repo.UpdatePhoneNumber(pr, aID, ph)
		if err != nil {
			return nil, err
		}
		return m, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return result.Value(0).(*me.Me), nil
}
