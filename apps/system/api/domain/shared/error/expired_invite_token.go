package error

import (
	"fmt"
	"github.com/google/uuid"
)

type ExpiredInviteToken struct {
	token uuid.UUID
}

func NewExpiredInviteToken(token uuid.UUID) *ExpiredInviteToken {
	return &ExpiredInviteToken{token}
}

func (e *ExpiredInviteToken) Token() uuid.UUID {
	return e.token
}

func (e *ExpiredInviteToken) Error() string {
	return fmt.Sprintf("token:%s is expired", e.Token().String())
}

func (e *ExpiredInviteToken) MessageKey() MessageKey {
	return ExpiredInviteTokenMessageKey
}

func (e *ExpiredInviteToken) Code() string {
	return "400-" + string(ExpiredInviteTokenCodeKey)
}
