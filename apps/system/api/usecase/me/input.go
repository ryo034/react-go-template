package me

import (
	"io"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type UpdateInput struct {
	Me *me.Me
}

type UpdateNameInput struct {
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

type UpdateProfilePhotoInput struct {
	AccountID account.ID
	PhotoID   media.ID
	File      io.Reader
	Ext       media.AvatarExt
	Size      int64
}

type RemoveProfilePhotoInput struct {
	AccountID account.ID
}

type LeaveWorkspaceInput struct {
	AccountID account.ID
}
