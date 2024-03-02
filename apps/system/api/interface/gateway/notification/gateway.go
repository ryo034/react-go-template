package notification

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"

	"github.com/ryo034/react-go-template/apps/system/api/domain/notification"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
)

type gateway struct {
	emailDriver email.Driver
}

func NewGateway(emailDriver email.Driver) notification.Repository {
	return &gateway{emailDriver}
}

func (g *gateway) NotifyOtpByEmail(ctx context.Context, email account.Email, otp string) error {
	return g.emailDriver.SendOTP(ctx, email, otp)
}

func (g *gateway) NotifyMembersInvited(ctx context.Context, inviter workspace.Inviter, is invitation.Invitations) (invitation.Invitations, invitation.Invitations) {
	return g.emailDriver.SendInvitations(ctx, inviter, is)
}
