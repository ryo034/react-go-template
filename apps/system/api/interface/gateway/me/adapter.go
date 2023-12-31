package me

import (
	"firebase.google.com/go/v4/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/business_entity/employee"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/multi_factor"
	employeeGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/business_entity/employee"
	storeGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store"
	staffGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/staff"
	userGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/account"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/datetime"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/phone"
	models "github.com/ryo034/react-go-template/packages/go/infrastructure/database/sqlboiler/api"
)

type Adapter interface {
	Adapt(s *models.User, fu *auth.UserRecord) (*me.Me, error)
	AdaptFirebaseUser(fu *auth.UserRecord, firstName account.FirstName, lastName account.LastName, phoneNumber *phone.Number) (*me.Me, error)
}

type adapter struct {
	a  userGateway.Adapter
	ea employeeGateway.Adapter
	la storeGateway.Adapter
	sa staffGateway.Adapter
}

func NewAdapter(a userGateway.Adapter, ea employeeGateway.Adapter, la storeGateway.Adapter, sa staffGateway.Adapter) Adapter {
	return &adapter{a, ea, la, sa}
}

func (a *adapter) Adapt(usr *models.User, fu *auth.UserRecord) (*me.Me, error) {
	mfs := make([]multi_factor.MultiFactor, 0)
	if fu.MultiFactor.EnrolledFactors != nil {
		for _, f := range fu.MultiFactor.EnrolledFactors {
			ph, err := phone.NewInternationalPhoneNumber(f.PhoneNumber)
			if err != nil {
				return nil, err
			}
			dt, _ := datetime.NewDateFromInt64(f.EnrollmentTimestamp)
			mfs = append(mfs, multi_factor.NewMultiFactor(ph, dt))
		}
	}

	var role *me.Role = nil

	if usr.R.Representatives != nil {
		au, err := a.a.Adapt(usr)
		if err != nil {
			return nil, err
		}
		role = me.NewAsRepresentative(employee.NewRepresentative(au))
	} else {
		ae, err := a.ea.Adapt(usr.R.Employees[0])
		if err != nil {
			return nil, err
		}
		role = me.NewAsEmployee(ae)
	}
	return me.NewMe(fu.EmailVerified && usr.EmailVerified, multi_factor.NewMultiFactors(mfs), role), nil
}

func (a *adapter) AdaptFirebaseUser(fu *auth.UserRecord, firstName account.FirstName, lastName account.LastName, phoneNumber *phone.Number) (*me.Me, error) {
	return nil, nil
	//id, err := account.NewID(fu.UserInfo.UID)
	//if err != nil {
	//	return nil, err
	//}
	//email, err := account.NewEmail(fu.UserInfo.Email)
	//if err != nil {
	//	return nil, err
	//}
	//mfs := make([]multi_factor.MultiFactor, 0)
	//if fu.MultiFactor.EnrolledFactors != nil {
	//	for _, f := range fu.MultiFactor.EnrolledFactors {
	//		ph, err := phone.NewInternationalPhoneNumber(f.PhoneNumber)
	//		if err != nil {
	//			return nil, err
	//		}
	//		dt, _ := datetime.NewDateFromInt64(f.EnrollmentTimestamp)
	//		mfs = append(mfs, multi_factor.NewMultiFactor(ph, dt))
	//	}
	//}
	//return me.NewMe(user.NewUser(id, email, phoneNumber, firstName, lastName), fu.EmailVerified, multi_factor.NewMultiFactors(mfs)), nil
}
