package workspace

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	dbErr "github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/error"
	"github.com/uptrace/bun"
	"time"
)

type Driver interface {
	FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (models.Workspaces, error)
	Create(ctx context.Context, exec bun.IDB, w *workspace.Workspace) (*models.Workspace, error)
	AddMember(ctx context.Context, exec bun.IDB, w *workspace.Workspace, m *member.Member) (*models.Member, error)
	FindMember(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*models.Member, error)
	FindAllMembers(ctx context.Context, exec bun.IDB, wID workspace.ID) (models.Members, error)
	InviteMembers(ctx context.Context, exec bun.IDB, inviter workspace.Inviter, is invitation.Invitations) error
	FindInviterWorkspaceFromToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*models.Workspace, error)
	FindActiveInvitationByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.Invitation, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (p *driver) FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (models.Workspaces, error) {
	var ws models.Workspaces
	err := exec.
		NewSelect().
		Model(&ws).
		Relation("Detail").
		Join("JOIN members ms ON ms.workspace_id = ws.workspace_id").
		Where("ms.system_account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return ws, nil
}

func (p *driver) Create(ctx context.Context, exec bun.IDB, w *workspace.Workspace) (*models.Workspace, error) {
	d := w.Detail()
	m := &models.Workspace{
		WorkspaceID: w.ID().Value(),
	}
	md := &models.WorkspaceDetail{
		WorkspaceID: w.ID().Value(),
		Subdomain:   d.Subdomain().ToString(),
		Name:        d.Name().ToString(),
	}
	if _, err := exec.NewInsert().Model(m).Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := exec.NewInsert().Model(md).Exec(ctx); err != nil {
		if dbErr.IsDuplicateError(err) {
			return nil, domainErr.NewConflicted("workspace_details", "subdomain")
		}
		return nil, err
	}
	m.Detail = md
	return m, nil
}

func (p *driver) AddMember(ctx context.Context, exec bun.IDB, w *workspace.Workspace, m *member.Member) (*models.Member, error) {
	mm := &models.Member{
		MemberID:        m.ID().Value(),
		WorkspaceID:     w.ID().Value(),
		SystemAccountID: m.User().AccountID().Value(),
	}
	if _, err := exec.NewInsert().Model(mm).Exec(ctx); err != nil {
		return nil, err
	}

	dn := ""
	if m.HasDisplayName() {
		dn = m.DisplayName().ToString()
	}
	mp := &models.MemberProfile{
		MemberID:       m.ID().Value(),
		MemberIDNumber: "",
		DisplayName:    dn,
	}
	if _, err := exec.NewInsert().Model(mp).Exec(ctx); err != nil {
		return nil, err
	}
	return mm, nil
}

func (p *driver) FindMember(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*models.Member, error) {
	m := &models.Member{}
	err := exec.
		NewSelect().
		Model(m).
		Relation("Profile").
		Relation("SystemAccount").
		Relation("SystemAccount.Profile").
		Relation("SystemAccount.PhoneNumber").
		Relation("Workspace").
		Where("ms.system_account_id = ?", aID.ToString()).
		Where("ms.workspace_id = ?", wID.Value()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (p *driver) FindAllMembers(ctx context.Context, exec bun.IDB, wID workspace.ID) (models.Members, error) {
	var ms models.Members
	err := exec.
		NewSelect().
		Model(&ms).
		Relation("Profile").
		Relation("SystemAccount").
		Relation("SystemAccount.Profile").
		Relation("SystemAccount.PhoneNumber").
		Relation("Workspace").
		Where("ms.workspace_id = ?", wID.Value()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (p *driver) InviteMembers(ctx context.Context, exec bun.IDB, inviter workspace.Inviter, is invitation.Invitations) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	invu := &models.InvitationUnit{
		InvitationUnitID: uid,
		WorkspaceID:      inviter.Workspace().ID().Value(),
		InvitedBy:        inviter.Member.ID().Value(),
	}
	if _, err = exec.NewInsert().Model(invu).Exec(ctx); err != nil {
		return err
	}

	invs := make([]*models.Invitation, 0, is.Size())
	invitees := make([]*models.Invitee, 0, is.Size())
	invns := make([]*models.InviteeName, 0)
	for _, i := range is.AsSlice() {
		invs = append(invs, &models.Invitation{
			InvitationID:     i.ID().Value(),
			InvitationUnitID: invu.InvitationUnitID,
		})

		invitees = append(invitees, &models.Invitee{
			InvitationID: i.ID().Value(),
			Email:        i.InviteeEmail().ToString(),
		})

		if i.DisplayName() != nil {
			invns = append(invns, &models.InviteeName{
				InvitationID: i.ID().Value(),
				DisplayName:  i.DisplayName().ToString(),
			})
		}
	}
	//TODO: Parallel insert
	if _, err = exec.NewInsert().Model(&invs).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&invitees).Exec(ctx); err != nil {
		return err
	}
	if len(invns) > 0 {
		if _, err = exec.NewInsert().Model(&invns).Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (p *driver) FindInviterWorkspaceFromToken(ctx context.Context, exec bun.IDB, token invitation.Token) (*models.Workspace, error) {
	invt := &models.InvitationToken{}
	err := exec.
		NewSelect().
		Model(invt).
		Relation("Invitation").
		Relation("Invitation.InvitationUnit").
		Relation("Invitation.InvitationUnit.Workspace").
		Relation("Invitation.InvitationUnit.Workspace.Detail").
		Where(fmt.Sprintf("%s.token = ?", models.InvitationTokenTableAliasName), token.Value()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invt.Invitation.InvitationUnit.Workspace, nil
}

func (p *driver) FindActiveInvitationByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.Invitation, error) {
	im := &models.Invitation{}
	err := exec.
		NewSelect().
		Model(im).
		Where("email = ?", email.ToString()).
		Where("used = ?", false).
		Where("expired_at > ?", time.Now()).
		Relation("Workspace").
		Relation("Workspace.Detail").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData(fmt.Sprintf("invited_member: email=%s and used=false and expired_at > now()", email.ToString()))
		}
		return nil, err
	}
	return im, nil
}
