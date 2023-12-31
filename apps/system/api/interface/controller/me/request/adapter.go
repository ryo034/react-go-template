package request

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/multi_factor"
	userRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/user/request"
	mePb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/me/v1"
)

type Adapter interface {
	Adapt(m *mePb.Me) (*me.Me, error)
}

func NewAdapter(userAdapter userRequest.Adapter) Adapter {
	return &adapter{userAdapter}
}

type adapter struct {
	userAdapter userRequest.Adapter
}

func (a *adapter) Adapt(m *mePb.Me) (*me.Me, error) {
	mfs := make([]multi_factor.MultiFactor, 0)
	return me.NewMe(m.EmailVerified, multi_factor.NewMultiFactors(mfs)), nil
}
