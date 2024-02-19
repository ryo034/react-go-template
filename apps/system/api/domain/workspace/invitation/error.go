package invitation

import (
	"fmt"
	"github.com/google/uuid"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
)

const (
	ExpiredInviteTokenMessageKey        domainErr.MessageKey = "ExpiredInvitation"
	InvalidInviteTokenMessageKey        domainErr.MessageKey = "InvalidInviteToken"
	AlreadyExpiredInvitationMessageKey  domainErr.MessageKey = "AlreadyExpiredInvitation"
	AlreadyRevokedInvitationMessageKey  domainErr.MessageKey = "AlreadyRevokedInvitation"
	AlreadyVerifiedInvitationMessageKey domainErr.MessageKey = "AlreadyVerifiedInvitation"
)

type ExpiredInvitation struct {
	Token uuid.UUID
}

func NewExpiredInvitation(token uuid.UUID) *ExpiredInvitation {
	return &ExpiredInvitation{token}
}

func (e *ExpiredInvitation) Error() string {
	return fmt.Sprintf("token:%s is expired", e.Token.String())
}

func (e *ExpiredInvitation) MessageKey() domainErr.MessageKey {
	return ExpiredInviteTokenMessageKey
}

type AlreadyExpiredInvitation struct {
	ID    ID
	Token uuid.UUID
}

func NewAlreadyExpiredInvitation(id ID, token uuid.UUID) *AlreadyExpiredInvitation {
	return &AlreadyExpiredInvitation{id, token}
}

func (e *AlreadyExpiredInvitation) Error() string {
	return fmt.Sprintf("id:%s token:%s is already expired", e.ID.ToString(), e.Token.String())
}

func (e *AlreadyExpiredInvitation) MessageKey() domainErr.MessageKey {
	return AlreadyExpiredInvitationMessageKey
}

type AlreadyRevokedInvitation struct {
	ID    ID
	Token uuid.UUID
}

func NewAlreadyRevokedInvitation(id ID, token uuid.UUID) *AlreadyRevokedInvitation {
	return &AlreadyRevokedInvitation{id, token}
}

func (e *AlreadyRevokedInvitation) Error() string {
	return fmt.Sprintf("id:%s token:%s is already revoked", e.ID.ToString(), e.Token.String())
}

func (e *AlreadyRevokedInvitation) MessageKey() domainErr.MessageKey {
	return AlreadyRevokedInvitationMessageKey
}

type AlreadyVerifiedInvitation struct {
	ID    ID
	Token uuid.UUID
}

func NewAlreadyVerifiedInvitation(id ID, token uuid.UUID) *AlreadyVerifiedInvitation {
	return &AlreadyVerifiedInvitation{id, token}
}

func (e *AlreadyVerifiedInvitation) Error() string {
	return fmt.Sprintf("id:%s token:%s is already verified", e.ID.ToString(), e.Token.String())
}

func (e *AlreadyVerifiedInvitation) MessageKey() domainErr.MessageKey {
	return AlreadyVerifiedInvitationMessageKey
}

type ExpiredInviteToken struct {
	ID    ID
	Token uuid.UUID
}

func NewExpiredInviteToken(id ID, token uuid.UUID) *ExpiredInviteToken {
	return &ExpiredInviteToken{id, token}
}

func (e *ExpiredInviteToken) Error() string {
	return fmt.Sprintf("id:%s token:%s is expired", e.ID.ToString(), e.Token.String())
}

type InvalidInviteToken struct {
	Token uuid.UUID
}

func NewInvalidInviteToken(token uuid.UUID) *InvalidInviteToken {
	return &InvalidInviteToken{token}
}

func (e *InvalidInviteToken) Error() string {
	return fmt.Sprintf("token:%s is invalid", e.Token.String())
}

func (e *InvalidInviteToken) MessageKey() domainErr.MessageKey {
	return InvalidInviteTokenMessageKey
}
