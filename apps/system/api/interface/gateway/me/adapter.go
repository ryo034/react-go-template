package me

import (
	"firebase.google.com/go/v4/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	memberGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
	userGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
)

type Adapter interface {
	Adapt(m *models.Member, fu *auth.UserRecord) (*me.Me, error)
	AdaptFirebaseUser(fu *auth.UserRecord, name account.Name, phoneNumber *phone.Number) (*me.Me, error)
}

type adapter struct {
	uga userGw.Adapter
	mga memberGw.Adapter
}

func NewAdapter(uga userGw.Adapter, mga memberGw.Adapter) Adapter {
	return &adapter{uga, mga}
}

func (a *adapter) Adapt(m *models.Member, fu *auth.UserRecord) (*me.Me, error) {
	mem, err := a.mga.Adapt(m)
	if err != nil {
		return nil, err
	}
	return me.NewMe(nil, mem), nil
}

func (a *adapter) AdaptFirebaseUser(fu *auth.UserRecord, name account.Name, phoneNumber *phone.Number) (*me.Me, error) {
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
