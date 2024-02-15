package invitation

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

type Adapter interface {
	Adapt(m *models.Invitation) (*invitation.Invitation, error)
	AdaptAll(m []*models.Invitation) (invitation.Invitations, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(i *models.Invitation) (*invitation.Invitation, error) {
	id := invitation.NewID(i.InvitationID)
	token := invitation.NewToken(i.Token)
	ex := invitation.NewExpiredAt(datetime.NewDatetime(i.ExpiredAt))
	ema, err := account.NewEmail(i.Email)
	if err != nil {
		return nil, err
	}
	dn, err := member.NewDisplayName(i.DisplayName)
	if err != nil {
		return nil, err
	}
	return invitation.NewInvitation(id, token, i.Verified, i.Used, ex, ema, dn), nil
}

func (a *adapter) AdaptAll(is []*models.Invitation) (invitation.Invitations, error) {
	mws := make([]*invitation.Invitation, 0, len(is))
	for _, i := range is {
		aw, err := a.Adapt(i)
		if err != nil {
			return nil, err
		}
		mws = append(mws, aw)
	}
	return invitation.NewInvitations(mws), nil
}
