package response

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	userResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/user/response"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(m *me.Me) (*openapi.Me, error)
}

type adapter struct {
	ua userResponse.Adapter
}

func NewAdapter(ua userResponse.Adapter) Adapter {
	return &adapter{ua}
}

func (a *adapter) Adapt(m *me.Me) (*openapi.Me, error) {
	//mfs := make([]*openapi.MultiFactor, 0, m.MultiFactors().Size())
	//for _, mf := range m.MultiFactors().AsSlice() {
	//	mfs = append(mfs, &openapi.MultiFactor{
	//		PhoneNumber: mf.PhoneNumber().ToString(),
	//	})
	//}
	var mf = openapi.OptMultiFactor{
		Value: openapi.MultiFactor{},
		Set:   false,
	}
	if m.MultiFactors().IsNotEmpty() {
		lt := m.MultiFactors().Latest()
		mf.Value = openapi.MultiFactor{
			PhoneNumber: lt.PhoneNumber().ToString(),
		}
	}
	resultMe := &openapi.Me{
		EmailVerified: m.EmailVerified(),
		MultiFactor:   mf,
	}
	return resultMe, nil
}
