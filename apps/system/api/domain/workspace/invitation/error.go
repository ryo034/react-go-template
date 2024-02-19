package invitation

import (
	"fmt"
	"github.com/google/uuid"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
)

const (
	ExpiredInviteTokenMessageKey        domainErr.MessageKey = "ExpiredInvitation"
	AlreadyExpiredInvitationMessageKey  domainErr.MessageKey = "AlreadyExpiredInvitation"
	AlreadyRevokedInvitationMessageKey  domainErr.MessageKey = "AlreadyRevokedInvitation"
	AlreadyVerifiedInvitationMessageKey domainErr.MessageKey = "AlreadyVerifiedInvitation"
	ExpiredInvitationCodeKey            domainErr.Code       = "001"
	AlreadyExpiredInvitationCodeKey     domainErr.Code       = "000"
	AlreadyRevokedInvitationCodeKey     domainErr.Code       = "000"
	AlreadyVerifiedInvitationCodeKey    domainErr.Code       = "000"
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

func (e *ExpiredInvitation) Code() string {
	return "400-" + string(ExpiredInvitationCodeKey)
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

func (e *AlreadyExpiredInvitation) Code() string {
	return "410-" + string(AlreadyExpiredInvitationCodeKey)
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

func (e *AlreadyRevokedInvitation) Code() string {
	return "410-" + string(AlreadyRevokedInvitationCodeKey)
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

func (e *AlreadyVerifiedInvitation) Code() string {
	return "409-" + string(AlreadyVerifiedInvitationCodeKey)
}
