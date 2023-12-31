package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/store"
	models "github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/sqlboiler/api"
	dbErr "github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/sqlboiler/error"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type (
	Driver interface {
		Save(ctx context.Context, exec boil.ContextExecutor, me *me.Me) (*models.User, error)
		Find(ctx context.Context, exec boil.ContextExecutor, aID account.ID) (*models.User, error)
		VerifyEmail(ctx context.Context, exec boil.ContextExecutor, aID account.ID) error
		Update(ctx context.Context, exec boil.ContextExecutor, me *me.Me) error
		UpdateEmail(ctx context.Context, exec boil.ContextExecutor, aID account.ID, em account.Email) (*models.User, error)
		UpdateName(ctx context.Context, exec boil.ContextExecutor, aID account.ID, fn account.FirstName, ln account.LastName) (*models.User, error)
		UpdatePhoneNumber(ctx context.Context, exec boil.ContextExecutor, aID account.ID, ph phone.Number) (*models.User, error)
		ExistOnLibrary(ctx context.Context, exec boil.ContextExecutor, aID account.ID, sID store.ID) (bool, error)
	}
)

type driver struct{}

func NewDriver() Driver {
	return &driver{}
}

func (d *driver) Save(ctx context.Context, exec boil.ContextExecutor, me *me.Me) (*models.User, error) {
	u := me.User()
	st := &models.User{
		UserID:    u.AccountID().ToString(),
		Email:     u.Email().ToString(),
		FirstName: u.FirstName().ToString(),
		LastName:  u.LastName().ToString(),
	}
	if err := st.Insert(ctx, exec, boil.Infer()); err != nil {
		return nil, err
	}
	return st, nil
}

func (d *driver) Find(ctx context.Context, exec boil.ContextExecutor, aID account.ID) (*models.User, error) {
	return models.Users(
		models.UserWhere.UserID.EQ(aID.ToString()),
		qm.Load(models.UserRels.UserPhoneNumber),
		qm.Load(models.UserRels.Representatives),
		qm.Load(qm.Rels(models.UserRels.Employees, models.EmployeeRels.EmployeePermission, models.EmployeePermissionRels.AnalyticsPermission)),
		qm.Load(qm.Rels(models.UserRels.Employees, models.EmployeeRels.EmployeePermission, models.EmployeePermissionRels.EmployeeOperationPermission)),
		qm.Load(qm.Rels(models.UserRels.Employees, models.EmployeeRels.EmployeePermission, models.EmployeePermissionRels.StorePermission)),
		qm.Load(qm.Rels(models.UserRels.Employees, models.EmployeeRels.EmployeePermission, models.EmployeePermissionRels.TransactionHistoryPermission)),
	).One(ctx, exec)
}

func (d *driver) VerifyEmail(ctx context.Context, exec boil.ContextExecutor, aID account.ID) error {
	u, err := models.FindUser(ctx, exec, aID.ToString())
	if err != nil {
		return err
	}
	u.EmailVerified = true
	_, err = u.Update(ctx, exec, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (d *driver) Update(ctx context.Context, exec boil.ContextExecutor, me *me.Me) error {
	u := me.User()
	st, err := d.Find(ctx, exec, u.AccountID())
	if err != nil {
		return err
	}
	st.Email = u.Email().ToString()
	st.FirstName = u.FirstName().ToString()
	st.LastName = u.LastName().ToString()
	st.EmailVerified = me.EmailVerified()
	if u.HasPhoneNumber() {
		st.R.UserPhoneNumber.PhoneNumber = u.PhoneNumber().ToString()
	}
	_, err = st.Update(ctx, exec, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (d *driver) UpdateEmail(ctx context.Context, exec boil.ContextExecutor, aID account.ID, em account.Email) (*models.User, error) {
	u, err := d.Find(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	u.Email = em.ToString()
	u.EmailVerified = false
	_, err = u.Update(ctx, exec, boil.Infer())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (d *driver) UpdateName(ctx context.Context, exec boil.ContextExecutor, aID account.ID, fn account.FirstName, ln account.LastName) (*models.User, error) {
	u, err := d.Find(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	u.FirstName = fn.ToString()
	u.LastName = ln.ToString()
	_, err = u.Update(ctx, exec, boil.Infer())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (d *driver) UpdatePhoneNumber(ctx context.Context, exec boil.ContextExecutor, aID account.ID, ph phone.Number) (*models.User, error) {
	u, err := d.Find(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	if u.R.UserPhoneNumber != nil {
		if _, err = u.R.UserPhoneNumber.Delete(ctx, exec); err != nil {
			return nil, err
		}
	}
	if err = u.SetUserPhoneNumber(ctx, exec, true, &models.UserPhoneNumber{PhoneNumber: ph.ToString()}); err != nil {
		if dbErr.IsDuplicateError(err) {
			return nil, domainErr.NewPhoneNumberAlreadyInUse(ph.ToString())
		}
		return nil, err
	}
	return u, nil
}

func (d *driver) ExistOnLibrary(ctx context.Context, exec boil.ContextExecutor, aID account.ID, sID store.ID) (bool, error) {
	//return models.Staffs(
	//	models.StaffWhere.UserID.EQ(aID.ToString()),
	//	models.StaffWhere.LibraryID.EQ(sID.ToString()),
	//).Exists(ctx, exec)
	return false, nil
}
