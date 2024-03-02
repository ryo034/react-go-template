package notification

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
)

type Repository interface {
	NotifyOtpByEmail(ctx context.Context, email account.Email, otp string) error
	NotifyMembersInvited(ctx context.Context, inviter workspace.Inviter, is invitation.Invitations) (invitation.Invitations, invitation.Invitations)
}
