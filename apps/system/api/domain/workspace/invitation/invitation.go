package invitation

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type Invitation struct {
	id           ID
	token        Token
	events       Events
	expiredAt    ExpiredAt
	inviteeEmail account.Email
	displayName  *member.DisplayName
}

func NewInvitation(id ID, token Token, events Events, expiredAt ExpiredAt, inviteeEmail account.Email, displayName *member.DisplayName) *Invitation {
	return &Invitation{id, token, events, expiredAt, inviteeEmail, displayName}
}

func GenInvitation(inviteeEmail string, displayName string) (*Invitation, error) {
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
	return &Invitation{id, token, nil, GenerateExpiredAt(), em, dn}, nil
}

func (i *Invitation) ID() ID {
	return i.id
}

func (i *Invitation) Token() Token {
	return i.token
}

func (i *Invitation) Events() Events {
	return i.events
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

func (i *Invitation) CheckActive() error {
	if i.IsVerified() {
		return NewAlreadyVerifiedInvitation(i.ID(), i.Token().Value())
	}
	if i.IsExpired() {
		return NewAlreadyExpiredInvitation(i.ID(), i.Token().Value())
	}
	if i.IsRevoked() {
		return NewAlreadyRevokedInvitation(i.ID(), i.Token().Value())
	}
	return nil
}

func (i *Invitation) IsVerified() bool {
	return i.Events() != nil &&
		i.Events().IsNotEmpty() &&
		i.Events().Latest() != nil &&
		i.Events().Latest().IsVerified()
}

func (i *Invitation) IsExpired() bool {
	return i.expiredAt.IsExpired()
}

func (i *Invitation) IsNotExpired() bool {
	return i.expiredAt.IsNotExpired()
}

func (i *Invitation) IsRevoked() bool {
	return i.Events() != nil &&
		i.Events().IsNotEmpty() &&
		i.Events().Latest() != nil &&
		i.Events().Latest().IsRevoked()
}
