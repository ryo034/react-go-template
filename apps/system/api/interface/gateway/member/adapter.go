package member

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	userGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
)

type Adapter interface {
	Adapt(m *models.Member) (*member.Member, error)
}

type adapter struct {
	uga userGw.Adapter
}

func NewAdapter(uga userGw.Adapter) Adapter {
	return &adapter{uga}
}

func (a *adapter) Adapt(m *models.Member) (*member.Member, error) {
	u, err := a.uga.Adapt(m.SystemAccount)
	if err != nil {
		return nil, err
	}
	idNumber, err := member.NewIDNumber(m.Profile.MemberIDNumber)
	if err != nil {
		return nil, err
	}

	dn, err := member.NewDisplayName(m.Profile.DisplayName)
	if err != nil {
		return nil, err
	}
	id := member.NewIDFromUUID(m.MemberID)
	return member.NewMember(id, u, dn, &idNumber), nil
}
