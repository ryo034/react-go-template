package member

import (
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"time"
)

type InvitedMember struct {
	id        uuid.UUID
	email     account.Email
	token     uuid.UUID
	verified  bool
	expiredAt time.Time
}

func NewInvitedMemberFromEmail(email account.Email, expiredAt time.Time) (*InvitedMember, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	token, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &InvitedMember{id, email, token, false, expiredAt}, nil
}

func NewInvitedMember(id uuid.UUID, email account.Email, token uuid.UUID, verified bool, expiredAt time.Time) *InvitedMember {
	return &InvitedMember{id, email, token, verified, expiredAt}
}

func (m *InvitedMember) Email() account.Email {
	return m.email
}

func (m *InvitedMember) ID() uuid.UUID {
	return m.id
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
