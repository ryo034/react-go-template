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
	inv := &models.Invitation{
		InvitationID: id.Value(),
	}
	err := exec.
		NewSelect().
		Model(inv).
		Relation("InviteeName").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
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
		Relation("Tokens").
		Relation("Events").
		Where(fmt.Sprintf("%s.invitation_id = ?", models.InvitationTableAliasName), id.Value()).
		//WherePK().
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
	invt, err := d.findTokenByEmail(ctx, exec, email)
	if err != nil {
		return nil, err
	}

	inve, err := d.findEvent(ctx, exec, invt.InvitationID)
	if err != nil {
		return nil, err
	}

	//has-manyかつWhereを使う場合うまくいかないので独自で実装する
	inv := &models.Invitation{}
	err = exec.
		NewSelect().
		Model(inv).
		Relation("InviteeName").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
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
		Where(fmt.Sprintf("%s.invitation_id = ?", models.InvitationTableAliasName), invt.InvitationID).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. email: %s", email.ToString()))
		}
		return nil, err
	}
	inv.Tokens = append(make([]*models.InvitationToken, 0, 1), invt)
	if inve != nil {
		inv.Events = []*models.InvitationEvent{inve}
	}
	return inv, nil
}

func (d *driver) findActiveToken(ctx context.Context, exec bun.IDB, token uuid.UUID) (*models.InvitationToken, error) {
	invt := &models.InvitationToken{}
	err := exec.
		NewSelect().
		Model(invt).
		Where("token = ?", token.String()).
		Where("expired_at > ?", time.Now()).
		Order("expired_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("active invitation not found. token: %s", token.String()))
		}
		return nil, err
	}
	return invt, nil
}

func (d *driver) findTokenByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.InvitationToken, error) {
	invitee := &models.Invitee{}
	if err := exec.
		NewSelect().
		Model(invitee).
		Column("invitation_id").
		Where("email = ?", email.ToString()).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. email: %s", email.ToString()))
		}
		return nil, err
	}

	invt := &models.InvitationToken{}
	err := exec.
		NewSelect().
		Model(invt).
		Where("invitation_id = ?", invitee.InvitationID).
		Where("expired_at > ?", time.Now()).
		Order("expired_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("active invitation not found. email: %s", email.ToString()))
		}
		return nil, err
	}
	return invt, nil
}

func (d *driver) FindAllReceivedByEmail(ctx context.Context, exec bun.IDB, email account.Email) ([]*models.Invitation, error) {
	var invitees []*models.Invitee = nil
	if err := exec.
		NewSelect().
		Model(&invitees).
		Column("invitation_id").
		Where("email = ?", email.ToString()).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. email: %s", email.ToString()))
		}
		return nil, err
	}
	if invitees == nil {
		return nil, nil
	}

	inviteeInvIDs := make([]uuid.UUID, 0, len(invitees))
	for _, inv := range invitees {
		inviteeInvIDs = append(inviteeInvIDs, inv.InvitationID)
	}

	// get not expired tokens
	var ts []*models.InvitationToken = nil
	if err := exec.
		NewSelect().
		Model(&ts).
		Where("expired_at > ?", time.Now()).
		Where("invitation_id IN (?)", bun.In(inviteeInvIDs)).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("active invitation not found. email: %s", email.ToString()))
		}
		return nil, err
	}
	if ts == nil {
		return nil, nil
	}

	invtInvIDs := make([]uuid.UUID, 0, len(ts))
	for _, t := range ts {
		invtInvIDs = append(invtInvIDs, t.InvitationID)
	}

	var invs []*models.Invitation
	if err := exec.
		NewSelect().
		Model(&invs).
		Relation("InviteeName").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
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
		Relation("Tokens", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("expired_at DESC").Limit(1)
		}).
		Relation("Events", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("created_at DESC")
		}).
		Where(fmt.Sprintf("%s.invitation_id IN (?)", models.InvitationTableAliasName), bun.In(invtInvIDs)).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. email: %s", email.ToString()))
		}
		return nil, err
	}

	result := make([]*models.Invitation, 0)
	for _, inv := range invs {
		if inv.Tokens[0].ExpiredAt.After(time.Now()) && inv.Events != nil && inv.Events[0].EventType == "verified" {
			result = append(result, inv)
		}
	}
	return result, nil
}

func (d *driver) findEvent(ctx context.Context, exec bun.IDB, id uuid.UUID) (*models.InvitationEvent, error) {
	inve := &models.InvitationEvent{}
	err := exec.
		NewSelect().
		Model(inve).
		Where("invitation_id = ?", id).
		Where("event_type != ?", "verified").
		Order("created_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			inve = nil
		} else {
			return nil, err
		}
	}
	return inve, nil
}

func (d *driver) FindActiveByToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*models.Invitation, error) {
	invt, err := d.findActiveToken(ctx, exec, token.Value())
	if err != nil {
		return nil, err
	}

	inve, err := d.findEvent(ctx, exec, invt.InvitationID)
	if err != nil {
		return nil, err
	}

	//has-manyかつWhereを使う場合うまくいかないので独自で実装する
	inv := &models.Invitation{}
	err = exec.
		NewSelect().
		Model(inv).
		Relation("InviteeName").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
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
		Where(fmt.Sprintf("%s.invitation_id = ?", models.InvitationTableAliasName), invt.InvitationID).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invitation not found. token: %s", token.String()))
		}
		return nil, err
	}
	inv.Tokens = append(make([]*models.InvitationToken, 0, 1), invt)
	if inve != nil {
		inv.Events = []*models.InvitationEvent{inve}
	}
	return inv, nil
}

func (d *driver) FindAllByWorkspace(ctx context.Context, exec bun.IDB, wID workspace.ID) ([]*models.Invitation, error) {
	var invs []*models.Invitation = nil
	err := exec.
		NewSelect().
		Model(&invs).
		Relation("InviteeName").
		Relation("InvitationUnit").
		Relation("InvitationUnit.Workspace").
		Relation("InvitationUnit.Workspace.Detail").
		Relation("InvitationUnit.Workspace.Detail.WorkspaceDetail").
		Relation("InvitationUnit.Member").
		Relation("InvitationUnit.Member.Profile").
		Relation("InvitationUnit.Member.Role").
		Relation("InvitationUnit.Member.Role.MemberRole").
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
		Relation("Tokens").
		Relation("Events").
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
	eid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	_, err = exec.NewInsert().Model(&models.InvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
		EventType:         "accepted",
	}).Exec(ctx)
	return err
}

func (d *driver) Revoke(ctx context.Context, exec bun.IDB, id invitation.ID) error {
	eid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	_, err = exec.NewInsert().Model(&models.InvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
		EventType:         "revoked",
	}).Exec(ctx)
	return err
}

func (d *driver) Resend(ctx context.Context, exec bun.IDB, id invitation.ID) (*models.Invitation, error) {
	inv, err := d.Find(ctx, exec, id)
	if err != nil {
		return nil, err
	}
	eid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	ev := &models.InvitationEvent{
		InvitationEventID: eid,
		InvitationID:      id.Value(),
		EventType:         "reissued",
	}
	_, err = exec.NewInsert().Model(ev).Exec(ctx)
	if err != nil {
		return nil, err
	}
	inv.Events = append(inv.Events, ev)
	return inv, nil
}
