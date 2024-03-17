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
	token := invitation.NewToken(i.Token.InvitationToken.Token)
	ex := invitation.NewExpiredAt(datetime.NewDatetime(i.Token.InvitationToken.ExpiredAt))
	ema, err := account.NewEmail(i.Invitee.Email)
	if err != nil {
		return nil, err
	}
	var dn *member.DisplayName
	if i.InviteeName != nil {
		dn = member.NewDisplayName(i.InviteeName.DisplayName)
	}
	inviter, err := a.ma.Adapt(i.InvitationUnit.Member)
	if err != nil {
		return nil, err
	}

	if i.Event == nil {
		return invitation.NewInvitation(id, token, nil, ex, ema, dn, inviter), nil
	}

	var evt *invitation.Event = nil
	switch i.Event.InvitationEvent.EventType {
	case "verified":
		tmpEvt := invitation.NewAsVerified(datetime.NewDatetime(i.Event.InvitationEvent.CreatedAt))
		evt = &tmpEvt
	case "revoked":
		tmpEvt := invitation.NewAsRevoked(datetime.NewDatetime(i.Event.InvitationEvent.CreatedAt))
		evt = &tmpEvt
	case "accepted":
		tmpEvt := invitation.NewAsAccepted(datetime.NewDatetime(i.Event.InvitationEvent.CreatedAt))
		evt = &tmpEvt
	case "reissued":
		tmpEvt := invitation.NewAsReissued(datetime.NewDatetime(i.Event.InvitationEvent.CreatedAt))
		evt = &tmpEvt
	default:
		return nil, fmt.Errorf("unknown event type: %s", i.Event.InvitationEvent.EventType)
	}
	return invitation.NewInvitation(id, token, evt, ex, ema, dn, inviter), nil
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
