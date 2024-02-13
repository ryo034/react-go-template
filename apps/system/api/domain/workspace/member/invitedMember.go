package member

import (
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"time"
)

type InvitedMember struct {
	id          uuid.UUID
	email       account.Email
	displayName DisplayName
	token       uuid.UUID
	verified    bool
	expiredAt   time.Time
}

func NewInvitedMemberFromEmail(email account.Email, displayName string, expiredAt time.Time) (*InvitedMember, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	token, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	dn, err := NewDisplayName(displayName)
	if err != nil {
		return nil, err
	}
	return &InvitedMember{id, email, dn, token, false, expiredAt}, nil
}

func NewInvitedMember(id uuid.UUID, email account.Email, displayName DisplayName, token uuid.UUID, verified bool, expiredAt time.Time) (*InvitedMember, error) {
	if expiredAt.Before(time.Now()) {
		return nil, domainError.NewExpiredInviteToken(token)
	}
	return &InvitedMember{id, email, displayName, token, verified, expiredAt}, nil
}

func (m *InvitedMember) ID() uuid.UUID {
	return m.id
}

func (m *InvitedMember) Email() account.Email {
	return m.email
}

func (m *InvitedMember) DisplayName() DisplayName {
	return m.displayName
}

func (m *InvitedMember) Token() uuid.UUID {
	return m.token
}

func (m *InvitedMember) Verified() bool {
	return m.verified
}

func (m *InvitedMember) ExpiredAt() time.Time {
	return m.expiredAt
}
