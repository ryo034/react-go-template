package invitation

import (
	"fmt"
	memberGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"

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
	ma memberGw.Adapter
}

func NewAdapter(ma memberGw.Adapter) Adapter {
	return &adapter{ma}
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
		} else if ev.EventType == "accepted" {
			evs = append(evs, invitation.NewAsAccepted(datetime.NewDatetime(ev.CreatedAt)))
		} else if ev.EventType == "reissued" {
			evs = append(evs, invitation.NewAsReissued(datetime.NewDatetime(ev.CreatedAt)))
		} else {
			return nil, fmt.Errorf("unknown event type: %s", ev.EventType)
		}
	}
	inviter, err := a.ma.Adapt(i.InvitationUnit.Member)
	if err != nil {
		return nil, err
	}
	return invitation.NewInvitation(id, token, invitation.NewEvents(evs), ex, ema, dn, inviter), nil
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
