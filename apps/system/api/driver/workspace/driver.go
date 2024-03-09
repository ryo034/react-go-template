package workspace

import (
	"context"

	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	dbErr "github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/error"
	"github.com/uptrace/bun"
)

type Driver interface {
	FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (models.Workspaces, error)
	Create(ctx context.Context, exec bun.IDB, w *workspace.Workspace) (*models.Workspace, error)
	Update(ctx context.Context, exec bun.IDB, w *workspace.Workspace) error
	AddMember(ctx context.Context, exec bun.IDB, w *workspace.Workspace, m *member.Member) (*models.Member, error)
	UpdateMemberRole(ctx context.Context, exec bun.IDB, assignor *member.Member, m *member.Member) error
	FindMember(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*models.Member, error)
	FindAllMembers(ctx context.Context, exec bun.IDB, wID workspace.ID) (models.Members, error)
	InviteMembers(ctx context.Context, exec bun.IDB, inviter workspace.Inviter, is invitation.Invitations) error
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (d *driver) FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (models.Workspaces, error) {
	var ws models.Workspaces
	err := exec.
		NewSelect().
		Model(&ws).
		Relation("Detail").
		Relation("Detail.WorkspaceDetail").
		Join("JOIN members ms ON ms.workspace_id = ws.workspace_id").
		Where("ms.account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return ws, nil
}

func (d *driver) Create(ctx context.Context, exec bun.IDB, w *workspace.Workspace) (*models.Workspace, error) {
	wdt := w.Detail()
	wldID, _ := uuid.NewV7()
	wo := &models.Workspace{
		WorkspaceID: w.ID().Value(),
	}
	wd := &models.WorkspaceDetail{
		WorkspaceDetailID: wldID,
		WorkspaceID:       w.ID().Value(),
		Subdomain:         wdt.Subdomain().ToString(),
		Name:              wdt.Name().ToString(),
	}
	wld := &models.WorkspaceLatestDetail{
		WorkspaceDetailID: wldID,
		WorkspaceID:       w.ID().Value(),
	}
	if _, err := exec.NewInsert().Model(wo).Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := exec.NewInsert().Model(wd).Exec(ctx); err != nil {
		if dbErr.IsDuplicateError(err) {
			return nil, domainErr.NewConflicted("workspace_details", "subdomain")
		}
		return nil, err
	}
	if _, err := exec.NewInsert().Model(wld).Exec(ctx); err != nil {
		return nil, err
	}
	wld.WorkspaceDetail = wd
	wo.Detail = wld
	return wo, nil
}

func (d *driver) Update(ctx context.Context, exec bun.IDB, w *workspace.Workspace) error {
	wdID, _ := uuid.NewV7()
	wdt := w.Detail()
	if _, err := exec.NewDelete().Model(&models.WorkspaceLatestDetail{}).Where("workspace_id = ?", w.ID().Value()).Exec(ctx); err != nil {
		return err
	}
	if _, err := exec.NewInsert().Model(&models.WorkspaceDetail{
		WorkspaceDetailID: wdID,
		WorkspaceID:       w.ID().Value(),
		Subdomain:         wdt.Subdomain().ToString(),
		Name:              wdt.Name().ToString(),
	}).Exec(ctx); err != nil {
		if dbErr.IsDuplicateError(err) {
			return domainErr.NewConflicted("workspace_details", "subdomain")
		}
		return err
	}
	if _, err := exec.NewInsert().Model(&models.WorkspaceLatestDetail{
		WorkspaceDetailID: wdID,
		WorkspaceID:       w.ID().Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func adaptMemberRole(mr member.Role) string {
	switch mr {
	case member.RoleOwner:
		return "owner"
	case member.RoleAdmin:
		return "admin"
	case member.RoleMember:
		return "member"
	default:
		return "guest"
	}
}

func (d *driver) AddMember(ctx context.Context, exec bun.IDB, w *workspace.Workspace, m *member.Member) (*models.Member, error) {
	mm := &models.Member{
		MemberID:    m.ID().Value(),
		WorkspaceID: w.ID().Value(),
		AccountID:   m.User().AccountID().Value(),
	}
	if _, err := exec.NewInsert().Model(mm).Exec(ctx); err != nil {
		return nil, err
	}

	p := m.Profile()
	dn := ""
	if p.HasDisplayName() {
		dn = p.DisplayName().ToString()
	}
	mp := &models.MemberProfile{
		MemberID:       m.ID().Value(),
		MemberIDNumber: "",
		DisplayName:    dn,
		Bio:            "",
	}
	if _, err := exec.NewInsert().Model(mp).Exec(ctx); err != nil {
		return nil, err
	}

	mrID, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}
	mr := &models.MemberRole{
		MemberRoleID: mrID,
		MemberID:     m.ID().Value(),
		Role:         adaptMemberRole(m.Role()),
		AssignedBy:   m.ID().Value(),
	}
	if _, err := exec.NewInsert().Model(mr).Exec(ctx); err != nil {
		return nil, err
	}
	mlr := &models.MemberLatestRole{
		MemberRoleID: mrID,
		MemberID:     m.ID().Value(),
	}
	if _, err := exec.NewInsert().Model(mlr).Exec(ctx); err != nil {
		return nil, err
	}
	mlr.MemberRole = mr
	mm.Role = mlr
	return mm, nil
}

func (d *driver) UpdateMemberRole(ctx context.Context, exec bun.IDB, assignor *member.Member, m *member.Member) error {
	mrID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	mr := &models.MemberRole{
		MemberRoleID: mrID,
		MemberID:     m.ID().Value(),
		Role:         adaptMemberRole(m.Role()),
		AssignedBy:   assignor.ID().Value(),
	}

	if _, err = exec.NewDelete().Model(&models.MemberLatestRole{}).Where("member_id = ?", m.ID().Value()).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(mr).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.MemberLatestRole{
		MemberRoleID: mrID,
		MemberID:     m.ID().Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (d *driver) FindMember(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*models.Member, error) {
	m := &models.Member{}
	err := exec.
		NewSelect().
		Model(m).
		Relation("Profile").
		Relation("Role").
		Relation("Role.MemberRole").
		Relation("Account").
		Relation("Account.AuthProviders").
		Relation("Account.Name").
		Relation("Account.Name.AccountName").
		Relation("Account.Email").
		Relation("Account.Email.AccountEmail").
		Relation("Account.PhoneNumber").
		Relation("Account.PhoneNumber.AccountPhoneNumber").
		Relation("Account.PhotoEvent").
		Relation("Account.PhotoEvent.AccountPhotoEvent").
		Relation("Account.PhotoEvent.AccountPhotoEvent.Photo").
		Relation("Workspace").
		Where("ms.account_id = ?", aID.ToString()).
		Where("ms.workspace_id = ?", wID.Value()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d *driver) FindAllMembers(ctx context.Context, exec bun.IDB, wID workspace.ID) (models.Members, error) {
	var ms models.Members
	err := exec.
		NewSelect().
		Model(&ms).
		Relation("Profile").
		Relation("Role").
		Relation("Role.MemberRole").
		Relation("Account").
		Relation("Account.AuthProviders").
		Relation("Account.Name").
		Relation("Account.Name.AccountName").
		Relation("Account.Email").
		Relation("Account.Email.AccountEmail").
		Relation("Account.PhoneNumber").
		Relation("Account.PhoneNumber.AccountPhoneNumber").
		Relation("Account.PhotoEvent").
		Relation("Account.PhotoEvent.AccountPhotoEvent").
		Relation("Account.PhotoEvent.AccountPhotoEvent.Photo").
		Relation("Workspace").
		Where("ms.workspace_id = ?", wID.Value()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *driver) InviteMembers(ctx context.Context, exec bun.IDB, inviter workspace.Inviter, is invitation.Invitations) error {
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
