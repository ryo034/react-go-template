package error

import (
	"fmt"
	"github.com/google/uuid"
)

type InvalidInviteToken struct {
	token uuid.UUID
}

func NewInvalidInviteToken(token uuid.UUID) *InvalidInviteToken {
	return &InvalidInviteToken{token}
}

func (e *InvalidInviteToken) Token() uuid.UUID {
	return e.token
}

func (e *InvalidInviteToken) Error() string {
	return fmt.Sprintf("invalid invite token:%s", e.Token())
}

func (e *InvalidInviteToken) MessageKey() MessageKey {
	return InvalidInviteTokenMessageKey
}

func (e *InvalidInviteToken) Code() string {
	return "400-000"
}
