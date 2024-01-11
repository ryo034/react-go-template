package request

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/multi_factor"
	"github.com/ryo034/react-go-template/apps/system/api/domain/member"
	userRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/user/request"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(m *openapi.Me) (*me.Me, error)
}

func NewAdapter(userAdapter userRequest.Adapter) Adapter {
	return &adapter{userAdapter}
}

type adapter struct {
	userAdapter userRequest.Adapter
}

func (a *adapter) Adapt(m *openapi.Me) (*me.Me, error) {
	mfs := make([]multi_factor.MultiFactor, 0)
	return me.NewMe(
		m.EmailVerified,
		multi_factor.NewMultiFactors(mfs),
		member.Member{},
	), nil
}
