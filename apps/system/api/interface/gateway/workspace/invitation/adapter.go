package invitation

import (
	"fmt"

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
	token := invitation.NewToken(i.Tokens[0].Token)
	ex := invitation.NewExpiredAt(datetime.NewDatetime(i.Tokens[0].ExpiredAt))
	ema, err := account.NewEmail(i.Invitee.Email)
	if err != nil {
		return nil, err
	}
	var dn *member.DisplayName
	if i.InviteeName != nil {
		dn = member.NewDisplayName(i.InviteeName.DisplayName)
	}
	evs := make([]invitation.Event, 0, len(i.Events))
	for _, ev := range i.Events {
		if ev.EventType == "verified" {
			evs = append(evs, invitation.NewAsVerified(datetime.NewDatetime(ev.CreatedAt)))
		} else if ev.EventType == "revoked" {
			evs = append(evs, invitation.NewAsRevoked(datetime.NewDatetime(ev.CreatedAt)))
		} else {
			//TODO: Error handling
			return nil, fmt.Errorf("unknown event type: %s", ev.EventType)
		}
	}
	return invitation.NewInvitation(id, token, invitation.NewEvents(evs), ex, ema, dn), nil
}

func (a *adapter) AdaptAll(is []*models.Invitation) (invitation.Invitations, error) {
	if is == nil {
		return nil, nil
	}
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
