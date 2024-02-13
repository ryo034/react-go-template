package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
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
	InviteMember(ctx context.Context, exec bun.IDB, wID workspace.ID, invitedBy member.InvitedBy, m *member.InvitedMember) error
	VerifyInvitedMember(ctx context.Context, exec bun.IDB, token uuid.UUID) (*models.InvitedMember, error)
	FindInviteeWorkspaceFromToken(ctx context.Context, exec bun.IDB, token uuid.UUID) (*models.Workspace, error)
	FindActiveInvitation(ctx context.Context, exec bun.IDB, email account.Email) (*models.InvitedMember, error)
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

	mp := &models.MemberProfile{
		MemberID:       m.ID().Value(),
		MemberIDNumber: "",
		DisplayName:    m.DisplayName().ToString(),
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

func (p *driver) InviteMember(ctx context.Context, exec bun.IDB, wID workspace.ID, invitedBy member.InvitedBy, m *member.InvitedMember) error {
	im := models.InvitedMember{
		InvitedMemberID: m.ID(),
		WorkspaceID:     wID.Value(),
		Email:           m.Email().ToString(),
		Used:            false,
		Token:           m.Token(),
		ExpiredAt:       m.ExpiredAt(),
		InvitedBy:       invitedBy.ID().Value(),
	}
	if _, err := exec.NewInsert().Model(&im).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (p *driver) VerifyInvitedMember(ctx context.Context, exec bun.IDB, token uuid.UUID) (*models.InvitedMember, error) {
	im := &models.InvitedMember{}
	err := exec.
		NewSelect().
		Model(im).
		Where("token = ?", token).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return im, nil
}

func (p *driver) FindInviteeWorkspaceFromToken(ctx context.Context, exec bun.IDB, token uuid.UUID) (*models.Workspace, error) {
	im := &models.InvitedMember{}
	err := exec.
		NewSelect().
		Model(im).
		Relation("Workspace").
		Relation("Workspace.Detail").
		Where("token = ?", token).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return im.Workspace, nil
}

func (p *driver) FindActiveInvitation(ctx context.Context, exec bun.IDB, email account.Email) (*models.InvitedMember, error) {
	im := &models.InvitedMember{}
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
		return nil, err
	}
	return im, nil
}
