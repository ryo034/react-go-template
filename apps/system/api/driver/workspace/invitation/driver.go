package invitation

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Find(ctx context.Context, exec bun.IDB, id invitation.ID) (*models.Invitation, error)
	FindByToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*models.Invitation, error)
	FindActiveByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.Invitation, error)
	FindAllReceivedByEmail(ctx context.Context, exec bun.IDB, email account.Email) ([]*models.Invitation, error)
	FindActiveByToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*models.Invitation, error)
	FindAllByWorkspace(ctx context.Context, exec bun.IDB, wID workspace.ID) ([]*models.Invitation, error)
	VerifyByToken(ctx context.Context, exec bun.IDB, token invitation.Token) error
	Accept(ctx context.Context, exec bun.IDB, id invitation.ID) error
	Revoke(ctx context.Context, exec bun.IDB, id invitation.ID) error
	Resend(ctx context.Context, exec bun.IDB, id invitation.ID) (*models.Invitation, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (d *driver) Find(ctx context.Context, exec bun.IDB, id invitation.ID) (*models.Invitation, error) {
	inv := &models.Invitation{InvitationID: id.Value()}
	err := exec.
		NewSelect().
		Model(inv).
		Relation("InviteeName").
		Relation("Event").
		Relation("Event.InvitationEvent").
		Relation("Token").
		Relation("Token.InvitationToken").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Profile.MemberProfile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
		Relation("InvitationUnit.Member.MembershipEvent").
		Relation("InvitationUnit.Member.MembershipEvent.MembershipEvent").
		Relation("InvitationUnit.Member.Account").
		Relation("InvitationUnit.Member.Account.AuthProviders").
		Relation("InvitationUnit.Member.Account.Name").
		Relation("InvitationUnit.Member.Account.Name.AccountName").
		Relation("InvitationUnit.Member.Account.Email").
		Relation("InvitationUnit.Member.Account.Email.AccountEmail").
		Relation("InvitationUnit.Member.Account.PhoneNumber").
		Relation("InvitationUnit.Member.Account.PhoneNumber.AccountPhoneNumber").
		Relation("InvitationUnit.Member.Account.PhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent.Photo").
		Relation("Invitee").
		Where(fmt.Sprintf("%s.invitation_id = ?", models.InvitationTableAliasName), id.Value()).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. id: %s", id.Value()))
		}
		return nil, err
	}
	return inv, nil
}

func (d *driver) FindByToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*models.Invitation, error) {
	invt := &models.InvitationToken{}
	err := exec.
		NewSelect().
		Model(invt).
		Where("token = ?", token.Value()).
		Order("expired_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. token: %s", token.String()))
		}
		return nil, err
	}
	return d.Find(ctx, exec, invitation.NewID(invt.InvitationID))

}

func (d *driver) FindActiveByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.Invitation, error) {
	inv := &models.Invitation{}
	err := exec.
		NewSelect().
		Model(inv).
		Relation("InviteeName").
		Relation("Event").
		Relation("Event.InvitationEvent").
		Relation("Token").
		Relation("Token.InvitationToken").
		Relation("Invitee").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Profile.MemberProfile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
		Relation("InvitationUnit.Member.MembershipEvent").
		Relation("InvitationUnit.Member.MembershipEvent.MembershipEvent").
		Relation("InvitationUnit.Member.Account").
		Relation("InvitationUnit.Member.Account.AuthProviders").
		Relation("InvitationUnit.Member.Account.Name").
		Relation("InvitationUnit.Member.Account.Name.AccountName").
		Relation("InvitationUnit.Member.Account.Email").
		Relation("InvitationUnit.Member.Account.Email.AccountEmail").
		Relation("InvitationUnit.Member.Account.PhoneNumber").
		Relation("InvitationUnit.Member.Account.PhoneNumber.AccountPhoneNumber").
		Relation("InvitationUnit.Member.Account.PhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent.Photo").
		Where("invte.email = ?", email.ToString()).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("linve__inve.event_type != ?", "verified").WhereOr("linve__inve.event_type IS NULL")
		}).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. email: %s", email.ToString()))
		}
		return nil, err
	}
	return inv, nil
}

func (d *driver) FindAllReceivedByEmail(ctx context.Context, exec bun.IDB, email account.Email) ([]*models.Invitation, error) {
	var invs []*models.Invitation
	err := exec.
		NewSelect().
		Model(&invs).
		Relation("InviteeName").
		Relation("Event").
		Relation("Event.InvitationEvent").
		Relation("Token").
		Relation("Token.InvitationToken").
		Relation("Invitee").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Profile.MemberProfile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
		Relation("InvitationUnit.Member.MembershipEvent").
		Relation("InvitationUnit.Member.MembershipEvent.MembershipEvent").
		Relation("InvitationUnit.Member.Account").
		Relation("InvitationUnit.Member.Account.AuthProviders").
		Relation("InvitationUnit.Member.Account.Name").
		Relation("InvitationUnit.Member.Account.Name.AccountName").
		Relation("InvitationUnit.Member.Account.Email").
		Relation("InvitationUnit.Member.Account.Email.AccountEmail").
		Relation("InvitationUnit.Member.Account.PhoneNumber").
		Relation("InvitationUnit.Member.Account.PhoneNumber.AccountPhoneNumber").
		Relation("InvitationUnit.Member.Account.PhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent.Photo").
		Where("invte.email = ?", email.ToString()).
		Where("linvt__invt.expired_at > ?", time.Now()).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Where("linve__inve.event_type = ?", "verified").
				Where("linve__inve.event_type = ?", "reissued").
				WhereOr("linve__inve.event_type IS NULL")
		}).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. email: %s", email.ToString()))
		}
		return nil, err
	}
	return invs, nil
}

func (d *driver) FindActiveByToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*models.Invitation, error) {
	inv := &models.Invitation{}
	err := exec.
		NewSelect().
		Model(inv).
		Relation("InviteeName").
		Relation("Event").
		Relation("Event.InvitationEvent").
		Relation("Token").
		Relation("Token.InvitationToken").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Profile.MemberProfile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
		Relation("InvitationUnit.Member.MembershipEvent").
		Relation("InvitationUnit.Member.MembershipEvent.MembershipEvent").
		Relation("InvitationUnit.Member.Account").
		Relation("InvitationUnit.Member.Account.AuthProviders").
		Relation("InvitationUnit.Member.Account.Name").
		Relation("InvitationUnit.Member.Account.Name.AccountName").
		Relation("InvitationUnit.Member.Account.Email").
		Relation("InvitationUnit.Member.Account.Email.AccountEmail").
		Relation("InvitationUnit.Member.Account.PhoneNumber").
		Relation("InvitationUnit.Member.Account.PhoneNumber.AccountPhoneNumber").
		Relation("InvitationUnit.Member.Account.PhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent.Photo").
		Relation("Invitee").
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("linve__inve.event_type != ?", "verified").WhereOr("linve__inve.event_type IS NULL")
		}).
		Where("linvt__invt.token = ?", token.String()).
		Where("linvt__invt.expired_at > ?", time.Now()).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. token: %s", token.String()))
		}
		return nil, err
	}
	return inv, nil
}

func (d *driver) FindAllByWorkspace(ctx context.Context, exec bun.IDB, wID workspace.ID) ([]*models.Invitation, error) {
	var invs []*models.Invitation = nil
	err := exec.
		NewSelect().
		Model(&invs).
		Relation("InviteeName").
		Relation("Event").
		Relation("Event.InvitationEvent").
		Relation("Token").
		Relation("Token.InvitationToken").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Profile.MemberProfile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
		Relation("InvitationUnit.Member.MembershipEvent").
		Relation("InvitationUnit.Member.MembershipEvent.MembershipEvent").
		Relation("InvitationUnit.Member.Account").
		Relation("InvitationUnit.Member.Account.AuthProviders").
		Relation("InvitationUnit.Member.Account.Name").
		Relation("InvitationUnit.Member.Account.Name.AccountName").
		Relation("InvitationUnit.Member.Account.Email").
		Relation("InvitationUnit.Member.Account.Email.AccountEmail").
		Relation("InvitationUnit.Member.Account.PhoneNumber").
		Relation("InvitationUnit.Member.Account.PhoneNumber.AccountPhoneNumber").
		Relation("InvitationUnit.Member.Account.PhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent").
		Relation("InvitationUnit.Member.Account.PhotoEvent.AccountPhotoEvent.Photo").
		Relation("Invitee").
		Where(fmt.Sprintf("%s.workspace_id = ?", models.InvitationUnitTableAliasName), wID.Value()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invs, nil
}

func (d *driver) VerifyByToken(ctx context.Context, exec bun.IDB, token invitation.Token) error {
	res, err := d.FindActiveByToken(ctx, exec, token)
	if err != nil {
		return err
	}
	eid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	_, err = exec.NewInsert().Model(&models.InvitationEvent{
		InvitationEventID: eid,
		InvitationID:      res.InvitationID,
		EventType:         "verified",
	}).Exec(ctx)
	return err
}

func (d *driver) Accept(ctx context.Context, exec bun.IDB, id invitation.ID) error {
	if _, err := exec.NewDelete().Model(&models.LatestInvitationEvent{}).Where("invitation_id = ?", id.Value()).Exec(ctx); err != nil {
		return err
	}
	eid, _ := uuid.NewV7()
	if _, err := exec.NewInsert().Model(&models.InvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
		EventType:         "accepted",
	}).Exec(ctx); err != nil {
		return err
	}
	if _, err := exec.NewInsert().Model(&models.LatestInvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (d *driver) Revoke(ctx context.Context, exec bun.IDB, id invitation.ID) error {
	if _, err := exec.NewDelete().Model(&models.LatestInvitationEvent{}).Where("invitation_id = ?", id.Value()).Exec(ctx); err != nil {
		return err
	}
	eid, _ := uuid.NewV7()
	if _, err := exec.NewInsert().Model(&models.InvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
		EventType:         "revoked",
	}).Exec(ctx); err != nil {
		return err
	}
	if _, err := exec.NewInsert().Model(&models.LatestInvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (d *driver) Resend(ctx context.Context, exec bun.IDB, id invitation.ID) (*models.Invitation, error) {
	if _, err := exec.NewDelete().Model(&models.LatestInvitationEvent{}).Where("invitation_id = ?", id.Value()).Exec(ctx); err != nil {
		return nil, err
	}
	eid, _ := uuid.NewV7()
	if _, err := exec.NewInsert().Model(&models.InvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
		EventType:         "reissued",
	}).Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := exec.NewInsert().Model(&models.LatestInvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
	}).Exec(ctx); err != nil {
		return nil, err
	}
	inv, err := d.Find(ctx, exec, id)
	if err != nil {
		return nil, err
	}
	return inv, nil
}
