package invitation

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type Invitation struct {
	id           ID
	token        Token
	verifiedAt   *VerifiedAt
	expiredAt    ExpiredAt
	inviteeEmail account.Email
	displayName  *member.DisplayName
}

func NewInvitation(id ID, token Token, verified *VerifiedAt, expiredAt ExpiredAt, inviteeEmail account.Email, displayName *member.DisplayName) *Invitation {
	return &Invitation{id, token, verified, expiredAt, inviteeEmail, displayName}
}

func GenInvitation(inviteeEmail string, displayName string) (*Invitation, error) {
	em, err := account.NewEmail(inviteeEmail)
	if err != nil {
		return nil, err
	}
	var dn *member.DisplayName
	if displayName != "" {
		tmpDn, err := member.NewDisplayName(displayName)
		if err != nil {
			return nil, err
		}
		dn = &tmpDn
	}
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

func (i *Invitation) Verified() *VerifiedAt {
	return i.verifiedAt
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
