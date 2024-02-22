package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
)

// ReceivedInvitation is only active verified invitation
type ReceivedInvitation struct {
	invitation *invitation.Invitation
	inviter    workspace.Inviter
}

func NewReceivedInvitation(inv *invitation.Invitation, inviter workspace.Inviter) (ReceivedInvitation, error) {
	if inv.IsRevoked() {
		return ReceivedInvitation{}, invitation.NewAlreadyRevokedInvitation(inv.ID(), inv.Token().Value())
	}
	if inv.IsExpired() {
		return ReceivedInvitation{}, invitation.NewAlreadyExpiredInvitation(inv.ID(), inv.Token().Value())
	}
	if inv.IsAccepted() {
		return ReceivedInvitation{}, invitation.NewAlreadyAcceptedInvitation(inv.ID(), inv.Token().Value())
	}
	return ReceivedInvitation{inv, inviter}, nil
}

func (r *ReceivedInvitation) Invitation() *invitation.Invitation {
	return r.invitation
}

func (r *ReceivedInvitation) Inviter() workspace.Inviter {
	return r.inviter
}
