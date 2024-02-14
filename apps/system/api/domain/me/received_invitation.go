package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
)

type ReceivedInvitation struct {
	invitation *invitation.Invitation
	inviter    workspace.Inviter
}

func NewReceivedInvitation(invitation *invitation.Invitation, inviter workspace.Inviter) *ReceivedInvitation {
	return &ReceivedInvitation{invitation, inviter}
}

func (r *ReceivedInvitation) Invitation() *invitation.Invitation {
	return r.invitation
}

func (r *ReceivedInvitation) Inviter() workspace.Inviter {
	return r.inviter
}
