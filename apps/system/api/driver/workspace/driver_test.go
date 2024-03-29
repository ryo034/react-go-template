//go:build testcontainers

package workspace

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Test_driver_FindAll_OK(t *testing.T) {
	var systemAccountIDUUID = uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	aID := account.NewIDFromUUID(systemAccountIDUUID)
	defaultTime := test.GetDefaultTime()
	wantErr := false

	wID := uuid.MustParse("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
	wdID := uuid.MustParse("018e200b-9d01-70ed-8c5a-5a5df2a98f11")

	want := models.Workspaces{
		{
			WorkspaceID: wID,
			CreatedAt:   defaultTime,
			Detail: &models.WorkspaceLatestDetail{
				WorkspaceDetailID: wdID,
				WorkspaceID:       wID,
				WorkspaceDetail: &models.WorkspaceDetail{
					WorkspaceDetailID: wdID,
					WorkspaceID:       wID,
					Name:              "Example",
					Subdomain:         "example",
					CreatedAt:         defaultTime,
				},
			},
			Members: nil,
		},
	}

	ctx := context.Background()
	t.Run("FindAll", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		pr := core.NewDatabaseProvider(db, db)
		got, err := NewDriver().FindAll(ctx, pr.GetExecutor(ctx, true), aID)
		if (err != nil) != wantErr {
			t.Errorf("FindAll() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			assert.EqualValuesf(t, want, got, "%v failed", "Find")
		}
	})
}

func workspacesEqual(a, b *models.Workspace) bool {
	return a.WorkspaceID == b.WorkspaceID &&
		a.Detail.WorkspaceDetail.Name == b.Detail.WorkspaceDetail.Name &&
		a.Detail.WorkspaceID == b.Detail.WorkspaceID &&
		a.Detail.WorkspaceDetail.Subdomain == b.Detail.WorkspaceDetail.Subdomain
}

func Test_driver_Create_OK(t *testing.T) {
	defaultTime := test.GetDefaultTime()
	wantErr := false

	wID := workspace.NewIDFromUUID(uuid.MustParse("018d5d05-6061-7f67-ba81-1626c15622c7"))
	wdID := uuid.MustParse("018e2216-64a3-7e72-8758-a948a5b8296a")

	want := &models.Workspace{
		WorkspaceID: wID.Value(),
		CreatedAt:   defaultTime,
		Detail: &models.WorkspaceLatestDetail{
			WorkspaceDetailID: wdID,
			WorkspaceID:       wID.Value(),
			WorkspaceDetail: &models.WorkspaceDetail{
				WorkspaceDetailID: wdID,
				WorkspaceID:       wID.Value(),
				Name:              "Example",
				Subdomain:         "example-test",
				CreatedAt:         defaultTime,
			},
		},
		Members: nil,
	}

	wsub, _ := workspace.NewSubdomain("example-test")
	wn, _ := workspace.NewName("Example")
	wd := workspace.NewDetail(wn, wsub)
	w := workspace.NewWorkspace(wID, wd)

	ctx := context.Background()
	t.Run("Create", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		pr := core.NewDatabaseProvider(db, db)
		got, err := NewDriver().Create(ctx, pr.GetExecutor(ctx, false), w)
		if (err != nil) != wantErr {
			t.Errorf("Create() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !workspacesEqual(got, want) {
			assert.EqualValuesf(t, want, got, "%v failed", "Create")
		}
	})
}

func Test_driver_AddMember_OK(t *testing.T) {
	defaultTime := test.GetDefaultTime()
	wantErr := false

	var systemAccountIDUUID = uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	aID := account.NewIDFromUUID(systemAccountIDUUID)

	wID := workspace.NewIDFromUUID(uuid.MustParse("018d5e03-498c-7db0-ace0-d915c3449a06"))
	mID := member.NewIDFromUUID(uuid.MustParse("018d5e05-4db2-7665-a362-8db2681f0666"))
	wdID := uuid.MustParse("018e200b-9d01-70ed-8c5a-5a5df2a98f11")

	ws := &models.Workspace{
		WorkspaceID: wID.Value(),
		CreatedAt:   defaultTime,
		Detail: &models.WorkspaceLatestDetail{
			WorkspaceDetailID: wdID,
			WorkspaceID:       wID.Value(),
			WorkspaceDetail: &models.WorkspaceDetail{
				WorkspaceDetailID: wdID,
				WorkspaceID:       wID.Value(),
				Name:              "Example",
				Subdomain:         "driver-example",
				CreatedAt:         defaultTime,
			},
		},
	}

	wsub, _ := workspace.NewSubdomain(ws.Detail.WorkspaceDetail.Subdomain)
	wn, _ := workspace.NewName(ws.Detail.WorkspaceDetail.Name)
	wd := workspace.NewDetail(wn, wsub)
	w := workspace.NewWorkspace(wID, wd)

	email, _ := account.NewEmail("account@example.com")
	name, _ := account.NewName("John Doe")
	u := user.NewUser(aID, email, &name, nil, nil)
	dn := member.NewDisplayName("John Doe")
	pr := member.NewProfile(dn, nil, member.NewAsEmptyBio())
	m := member.NewMember(mID, u, pr, member.RoleAdmin, member.MembershipStatusActive)

	want := &models.Member{
		MemberID:    mID.Value(),
		WorkspaceID: ws.WorkspaceID,
		AccountID:   systemAccountIDUUID,
		CreatedAt:   defaultTime,
	}

	ctx := context.Background()
	t.Run("AddMember", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		prov := core.NewDatabaseProvider(db, db)
		exec := prov.GetExecutor(ctx, false)
		_, err := NewDriver().Create(ctx, exec, w)
		if err != nil {
			t.Errorf("Create() error = %v", err)
			return
		}
		got, err := NewDriver().AddMember(ctx, exec, w, m)
		if (err != nil) != wantErr {
			t.Errorf("AddMember() error = %v, wantErr %v", err, wantErr)
			return
		}
		if got.MemberID != want.MemberID {
			assert.EqualValuesf(t, want, got, "%v failed", "AddMember")
		}
		if got.WorkspaceID != want.WorkspaceID {
			assert.EqualValuesf(t, want, got, "%v failed", "AddMember")
		}
		if got.AccountID != want.AccountID {
			assert.EqualValuesf(t, want, got, "%v failed", "AddMember")
		}
	})
}
