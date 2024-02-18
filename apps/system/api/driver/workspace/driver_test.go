//go:build testcontainers

package workspace

import (
	"context"
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
	"reflect"
	"testing"
)

func Test_driver_FindAll_OK(t *testing.T) {
	var systemAccountIDUUID = uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	aID := account.NewIDFromUUID(systemAccountIDUUID)
	defaultTime := test.GetDefaultTime()
	wantErr := false

	wID := uuid.MustParse("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")

	want := models.Workspaces{
		{
			WorkspaceID: wID,
			CreatedAt:   defaultTime,
			Detail: &models.WorkspaceDetail{
				WorkspaceID: wID,
				Name:        "Example",
				CreatedAt:   defaultTime,
				UpdatedAt:   defaultTime,
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
		a.Detail.Name == b.Detail.Name &&
		a.Detail.WorkspaceID == b.Detail.WorkspaceID &&
		a.Detail.Subdomain == b.Detail.Subdomain
}

func Test_driver_Create_OK(t *testing.T) {
	defaultTime := test.GetDefaultTime()
	wantErr := false

	wID := workspace.NewIDFromUUID(uuid.MustParse("018d5d05-6061-7f67-ba81-1626c15622c7"))

	want := &models.Workspace{
		WorkspaceID: wID.Value(),
		CreatedAt:   defaultTime,
		Detail: &models.WorkspaceDetail{
			WorkspaceID: wID.Value(),
			Name:        "Example",
			Subdomain:   "example-test",
			CreatedAt:   defaultTime,
			UpdatedAt:   defaultTime,
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

	ws := &models.Workspace{
		WorkspaceID: wID.Value(),
		CreatedAt:   defaultTime,
		Detail: &models.WorkspaceDetail{
			WorkspaceID: wID.Value(),
			Name:        "Example",
			Subdomain:   "example-test",
			CreatedAt:   defaultTime,
			UpdatedAt:   defaultTime,
		},
	}

	wsub, _ := workspace.NewSubdomain(ws.Detail.Subdomain)
	wn, _ := workspace.NewName(ws.Detail.Name)
	wd := workspace.NewDetail(wn, wsub)
	w := workspace.NewWorkspace(wID, wd)

	email, _ := account.NewEmail("system_account@example.com")
	name, _ := account.NewName("John Doe")
	u := user.NewUser(aID, email, &name, nil)
	dn := member.NewDisplayName("John Doe")
	m := member.NewMember(mID, u, dn, nil)

	want := &models.Member{
		MemberID:        mID.Value(),
		WorkspaceID:     ws.WorkspaceID,
		SystemAccountID: systemAccountIDUUID,
		CreatedAt:       defaultTime,
	}

	ctx := context.Background()
	t.Run("AddMember", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		pr := core.NewDatabaseProvider(db, db)
		ws, _ = NewDriver().Create(ctx, pr.GetExecutor(ctx, false), w)
		got, err := NewDriver().AddMember(ctx, pr.GetExecutor(ctx, false), w, m)
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
		if got.SystemAccountID != want.SystemAccountID {
			assert.EqualValuesf(t, want, got, "%v failed", "AddMember")
		}
	})
}
