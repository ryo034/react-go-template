package invitation

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type Invitation struct {
	id           ID
	token        Token
	latestEvent  *Event
	expiredAt    ExpiredAt
	inviteeEmail account.Email
	displayName  *member.DisplayName
	inviter      *member.Member
}

func NewInvitation(id ID, token Token, latestEvent *Event, expiredAt ExpiredAt, inviteeEmail account.Email, displayName *member.DisplayName, inviter *member.Member) *Invitation {
	return &Invitation{id, token, latestEvent, expiredAt, inviteeEmail, displayName, inviter}
}

func GenInvitation(inviteeEmail string, displayName string, inviter *member.Member) (*Invitation, error) {
	em, err := account.NewEmail(inviteeEmail)
	if err != nil {
		return nil, err
	}
	dn := member.NewDisplayName(displayName)
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	token := GenerateToken()
	if err != nil {
		return nil, err
	}
	return &Invitation{id, token, nil, GenerateExpiredAt(), em, dn, inviter}, nil
}

func (i *Invitation) ID() ID {
	return i.id
}

func (i *Invitation) Token() Token {
	return i.token
}

func (i *Invitation) LatestEvent() *Event {
	return i.latestEvent
}

func (i *Invitation) ExpiredAt() ExpiredAt {
	return i.expiredAt
}

func (i *Invitation) InviteeEmail() account.Email {
	return i.inviteeEmail
}

func (i *Invitation) DisplayName() *member.DisplayName {
	return i.displayName
}

func (i *Invitation) Inviter() *member.Member {
	return i.inviter
}

func (i *Invitation) ValidateCanAccept() error {
	if i.IsAccepted() {
		return NewAlreadyAcceptedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsRevoked() {
		return NewAlreadyRevokedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsExpired() {
		return NewAlreadyExpiredInvitation(i.ID(), i.Token().Value())
	}
	return nil
}

func (i *Invitation) ValidateCanRevoke(aID account.ID) error {
	if i.Inviter().User().AccountID() != aID {
		return domainErr.NewForbidden("revoke can only be done by the inviter")
	}
	if i.IsAccepted() {
		return NewAlreadyAcceptedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsRevoked() {
		return NewAlreadyRevokedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsVerified() {
		return NewAlreadyVerifiedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsExpired() {
		return NewAlreadyExpiredInvitation(i.ID(), i.Token().Value())
	}
	return nil
}

func (i *Invitation) ValidateCanResend(aID account.ID) error {
	if i.Inviter().User().AccountID() != aID {
		return domainErr.NewForbidden("resend can only be done by the inviter")
	}
	if i.IsAccepted() {
		return NewAlreadyAcceptedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsRevoked() {
		return NewAlreadyRevokedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsVerified() {
		return NewAlreadyVerifiedInvitation(i.ID(), i.Token().Value())
	}
	return nil
}

func (i *Invitation) ValidateCanVerify(token Token) error {
	if i.Token().NotEquals(token) {
		return NewInvalidInviteToken(token.Value())
	}
	if i.IsAccepted() {
		return NewAlreadyAcceptedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsRevoked() {
		return NewAlreadyRevokedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsVerified() {
		return NewAlreadyVerifiedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsExpired() {
		return NewAlreadyExpiredInvitation(i.ID(), i.Token().Value())
	}
	return nil
}

func (i *Invitation) ValidateCanGetByToken() error {
	if i.IsAccepted() {
		return NewAlreadyAcceptedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsRevoked() {
		return NewAlreadyRevokedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsExpired() {
		return NewAlreadyExpiredInvitation(i.ID(), i.Token().Value())
	}
	return nil
}

func (i *Invitation) IsVerified() bool {
	return i.latestEvent != nil && i.latestEvent.IsVerified()
}

func (i *Invitation) IsExpired() bool {
	return i.expiredAt.IsExpired()
}

func (i *Invitation) IsNotExpired() bool {
	return i.expiredAt.IsNotExpired()
}

func (i *Invitation) IsRevoked() bool {
	return i.latestEvent != nil && i.latestEvent.IsRevoked()
}

func (i *Invitation) IsAccepted() bool {
	return i.latestEvent != nil && i.latestEvent.IsAccepted()
}
