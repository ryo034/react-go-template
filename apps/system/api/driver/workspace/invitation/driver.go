package invitation

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
	"time"
)

type Driver interface {
	Find(ctx context.Context, exec bun.IDB, id invitation.ID) (*models.Invitation, error)
	FindActiveInvitationByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.Invitation, error)
	FindByToken(ctx context.Context, exec bun.IDB, token uuid.UUID) (*models.Invitation, error)
	VerifyByToken(ctx context.Context, exec bun.IDB, token uuid.UUID) error
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (p *driver) Find(ctx context.Context, exec bun.IDB, id invitation.ID) (*models.Invitation, error) {
	inv := models.Invitation{
		InvitationID: id.Value(),
	}
	err := exec.
		NewSelect().
		Model(&inv).
		WherePK().
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. id: %s", id.Value()))
		}
		return nil, err
	}
	return &inv, nil
}

func (p *driver) FindActiveInvitationByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.Invitation, error) {
	im := &models.Invitation{}
	err := exec.
		NewSelect().
		Model(im).
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.SystemAccount").
		Relation("InvitationUnit.Member.SystemAccount.Profile").
		Relation("InvitationUnit.Member.SystemAccount.PhoneNumber").
		Relation("Invitee").
		Relation("InviteeName").
		Relation("Events").
		Relation("Tokens").
		Where("invitees.email = ?", email.ToString()).
		Where("invitation_tokens.expired_at > ?", time.Now()).
		Where("invitations_events.event_type != ?", "verified").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invited_member: email=%s and expired_at > now()", email.ToString()))
		}
		return nil, err
	}
	return im, nil
}

func (p *driver) FindByToken(ctx context.Context, exec bun.IDB, token uuid.UUID) (*models.Invitation, error) {
	inv := models.Invitation{}
	err := exec.
		NewSelect().
		Model(&inv).
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.SystemAccount").
		Relation("InvitationUnit.Member.SystemAccount.Profile").
		Relation("InvitationUnit.Member.SystemAccount.PhoneNumber").
		Relation("Invitee").
		Relation("InviteeName").
		Relation("Events").
		Relation("Tokens").
		Where("invitation_tokens.token = ?", token).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. token: %s", token.String()))
		}
		return nil, err
	}
	return &inv, nil
}

func (p *driver) VerifyByToken(ctx context.Context, exec bun.IDB, token uuid.UUID) error {
	res, err := p.FindByToken(ctx, exec, token)
	if err != nil {
		return err
	}
	_, err = exec.NewInsert().Model(&models.InvitationEvent{
		InvitationID: res.InvitationID,
		EventType:    "verified",
	}).Exec(ctx)
	return err
}
