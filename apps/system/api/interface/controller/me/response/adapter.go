package response

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	userResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/user/response"
	mePb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/me/v1"
)

type Adapter interface {
	Adapt(m *me.Me) (*mePb.Me, error)
}

type adapter struct {
	ua userResponse.Adapter
}

func NewAdapter(ua userResponse.Adapter) Adapter {
	return &adapter{ua}
}

func (a *adapter) Adapt(m *me.Me) (*mePb.Me, error) {
	mfs := make([]*mePb.MultiFactor, 0, m.MultiFactors().Size())
	for _, mf := range m.MultiFactors().AsSlice() {
		mfs = append(mfs, &mePb.MultiFactor{
			PhoneNumber: mf.PhoneNumber().ToString(),
		})
	}
	var mf *mePb.MultiFactor = nil
	if m.MultiFactors().IsNotEmpty() {
		lt := m.MultiFactors().Latest()
		mf = &mePb.MultiFactor{
			PhoneNumber: lt.PhoneNumber().ToString(),
		}
	}
	resultMe := &mePb.Me{
		EmailVerified: m.EmailVerified(),
		MultiFactor:   mf,
	}
	return resultMe, nil
}
