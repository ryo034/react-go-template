package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type UpdateInput struct {
	Me *me.Me
}

type UpdateProfileInput struct {
	AccountID account.ID
	Name      account.Name
}

type UpdateMemberProfileInput struct {
	AccountID account.ID
	Profile   member.Profile
}

type AcceptInvitationInput struct {
	AccountID    account.ID
	InvitationID invitation.ID
}
