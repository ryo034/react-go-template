package me

import (
	"firebase.google.com/go/v4/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/multi_factor"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	models "github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/sqlboiler/api"
	userGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
)

type Adapter interface {
	Adapt(s *models.User, fu *auth.UserRecord) (*me.Me, error)
	AdaptFirebaseUser(fu *auth.UserRecord, firstName account.FirstName, lastName account.LastName, phoneNumber *phone.Number) (*me.Me, error)
}

type adapter struct {
	a userGateway.Adapter
}

func NewAdapter(a userGateway.Adapter) Adapter {
	return &adapter{a}
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
	return me.NewMe(fu.EmailVerified && usr.EmailVerified, multi_factor.NewMultiFactors(mfs)), nil
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
