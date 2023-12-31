package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/store"
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	meDriver "github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/me"
	businessEntityGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/business_entity"
	storeGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/account"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/phone"
	"github.com/ryo034/react-go-template/packages/go/infrastructure/database/sqlboiler/core"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type repository struct {
	ud  meDriver.Driver
	fd  firebaseDriver.Driver
	a   Adapter
	sa  storeGateway.Adapter
	bea businessEntityGateway.Adapter
}

func NewRepository(ud meDriver.Driver, fd firebaseDriver.Driver, a Adapter, sa storeGateway.Adapter, bea businessEntityGateway.Adapter) me.Repository {
	return &repository{ud, fd, a, sa, bea}
}

func (r *repository) SaveFromTemporary(ctx context.Context, aID account.ID, firstName account.FirstName, lastName account.LastName) (*me.Me, error) {
	fu, err := r.fd.GetUser(ctx, aID)
	if err != nil {
		return nil, err
	}
	m, err := r.a.AdaptFirebaseUser(fu, firstName, lastName, nil)
	if err != nil {
		return nil, err
	}
	if err := r.fd.UpdateMe(ctx, m); err != nil {
		return nil, err
	}
	res, err := r.ud.Save(ctx, core.ToExec(ctx, false), m)
	if err != nil {
		return nil, err
	}
	return r.a.Adapt(res, fu)
}

func (r *repository) VerifyEmail(ctx context.Context, aID account.ID) error {
	return r.ud.VerifyEmail(ctx, core.ToExec(ctx, false), aID)
}

func (r *repository) Find(ctx context.Context, exec boil.ContextExecutor, aID account.ID) (*me.Me, error) {
	res, err := r.ud.Find(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	fu, err := r.fd.GetUser(ctx, aID)
	if err != nil {
		return nil, err
	}
	return r.a.Adapt(res, fu)
}

func (r *repository) Update(ctx context.Context, me *me.Me) error {
	return r.ud.Update(ctx, core.ToExec(ctx, false), me)
}

func (r *repository) ExistOnStore(ctx context.Context, aID account.ID, lID store.ID) (bool, error) {
	return r.ud.ExistOnLibrary(ctx, core.ToExec(ctx, true), aID, lID)
}

func (r *repository) EmailVerified(ctx context.Context, aID account.ID) (bool, string, error) {
	return r.fd.EmailVerified(ctx, aID)
}

func (r *repository) UpdateEmail(ctx context.Context, aID account.ID, em account.Email) (*me.Me, error) {
	if err := r.fd.UpdateEmail(ctx, aID, em); err != nil {
		return nil, err
	}
	m, err := r.ud.UpdateEmail(ctx, core.ToExec(ctx, false), aID, em)
	if err != nil {
		return nil, err
	}
	fu, err := r.fd.GetUser(ctx, aID)
	if err != nil {
		return nil, err
	}
	return r.a.Adapt(m, fu)
}

func (r *repository) UpdateName(ctx context.Context, aID account.ID, fin account.FirstName, ln account.LastName) (*me.Me, error) {
	if err := r.fd.UpdateName(ctx, aID, fin, ln); err != nil {
		return nil, err
	}
	m, err := r.ud.UpdateName(ctx, core.ToExec(ctx, false), aID, fin, ln)
	if err != nil {
		return nil, err
	}
	fu, err := r.fd.GetUser(ctx, aID)
	if err != nil {
		return nil, err
	}
	return r.a.Adapt(m, fu)
}

func (r *repository) UpdatePhoneNumber(ctx context.Context, aID account.ID, ph phone.Number) (*me.Me, error) {
	if err := r.fd.UpdatePhoneNumber(ctx, aID, ph); err != nil {
		return nil, err
	}
	m, err := r.ud.UpdatePhoneNumber(ctx, core.ToExec(ctx, false), aID, ph)
	if err != nil {
		return nil, err
	}
	fu, err := r.fd.GetUser(ctx, aID)
	if err != nil {
		return nil, err
	}
	return r.a.Adapt(m, fu)
}
