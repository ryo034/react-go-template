//go:generate gomockhandler -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package me

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/uptrace/bun"
)

type Repository interface {
	Find(ctx context.Context, exec bun.IDB, mID member.ID) (*Me, error)
	FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error)
	RecordLogin(ctx context.Context, exec bun.IDB, m *Me) error
	SetCurrentProvider(ctx context.Context, p *provider.Provider) context.Context
	SetMe(ctx context.Context, m *Me) error
	FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error)
	FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*Me, error)
	UpdateName(ctx context.Context, exec bun.IDB, usr *user.User) error
	UpdateMemberProfile(ctx context.Context, exec bun.IDB, m *member.Member) (*member.Member, error)
	UpdateProfilePhoto(ctx context.Context, exec bun.IDB, m *Me, photo *media.UploadPhoto) error
	RemoveProfilePhoto(ctx context.Context, exec bun.IDB, m *Me) error
	AcceptInvitation(ctx context.Context, exec bun.IDB, id invitation.ID) error
	FindAllActiveReceivedInvitations(ctx context.Context, exec bun.IDB, aID account.ID) (ReceivedInvitations, error)
}
