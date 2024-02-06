//go:build testcontainers

package me

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
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
	"time"
)

const systemAccountID = "394e67b6-2850-4ddf-a4c9-c2a619d5bf70"

var systemAccountIDUUID = uuid.MustParse(systemAccountID)

func Test_driver_Find_OK(t *testing.T) {
	defaultTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-10 12:00:00")
	if err != nil {
		t.Fatalf("failed to parse defaultTime: %v", err)
	}
	memberID := uuid.MustParse("377eba35-5560-4f48-a99d-19cbd6a82b0d")

	workspaceID := uuid.MustParse("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
	wID := workspace.NewIDFromUUID(workspaceID)

	want := &models.Member{
		MemberID:        memberID,
		WorkspaceID:     wID.Value(),
		SystemAccountID: systemAccountIDUUID,
		CreatedAt:       defaultTime,
		SystemAccount: &models.SystemAccount{
			SystemAccountID: systemAccountIDUUID,
			CreatedAt:       defaultTime,
			PhoneNumber:     nil,
			Profile: &models.SystemAccountProfile{
				SystemAccountID: systemAccountIDUUID,
				Name:            "John Doe",
				Email:           "system_account@example.com",
				CreatedAt:       defaultTime,
				UpdatedAt:       defaultTime,
			},
		},
		Profile: &models.MemberProfile{
			MemberID:       memberID,
			MemberIDNumber: "DEV-12345",
			DisplayName:    "John Doe",
			CreatedAt:      defaultTime,
			UpdatedAt:      defaultTime,
		},
		Workspace: &models.Workspace{
			WorkspaceID: wID.Value(),
			CreatedAt:   defaultTime,
			Detail: &models.WorkspaceDetail{
				WorkspaceID: wID.Value(),
				Name:        "Example",
				Subdomain:   "example",
				CreatedAt:   defaultTime,
				UpdatedAt:   defaultTime,
			},
			Members: nil,
		},
	}
	wantErr := false
	ctx := context.Background()
	t.Run("Find", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		pr := core.NewDatabaseProvider(db, db)
		got, err := NewDriver().Find(ctx, pr.GetExecutor(ctx, true), member.NewIDFromUUID(memberID))
		if (err != nil) != wantErr {
			t.Errorf("Find() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			assert.EqualValuesf(t, want, got, "%v failed", "Find")
		}
	})
}

func Test_driver_UpdateName_OK(t *testing.T) {
	defaultTime := test.GetDefaultTime()
	wantErr := false

	var systemAccountIDUUID = uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	aID := account.NewIDFromUUID(systemAccountIDUUID)

	name, _ := account.NewName("John Doe 2")

	want := &models.SystemAccount{
		SystemAccountID: systemAccountIDUUID,
		CreatedAt:       defaultTime,
		PhoneNumber: &models.SystemAccountPhoneNumber{
			SystemAccountID: systemAccountIDUUID,
			PhoneNumber:     "",
			CreatedAt:       defaultTime,
			UpdatedAt:       defaultTime,
		},
		Profile: &models.SystemAccountProfile{
			SystemAccountID: systemAccountIDUUID,
			Name:            "John Doe 2",
			Email:           "",
			CreatedAt:       defaultTime,
			UpdatedAt:       defaultTime,
		},
	}

	ctx := context.Background()
	t.Run("UpdateName", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		pr := core.NewDatabaseProvider(db, db)
		res, err := NewDriver().UpdateName(ctx, pr.GetExecutor(ctx, false), aID, name)
		if (err != nil) != wantErr {
			t.Errorf("UpdateName() error = %v, wantErr %v", err, wantErr)
			return
		}
		if res.Profile.Name != want.Profile.Name {
			assert.EqualValuesf(t, want.Profile.Name, res.Profile.Name, "%v failed", "UpdateName")
		}
		if res.SystemAccountID != want.SystemAccountID {
			assert.EqualValuesf(t, want.SystemAccountID, res.SystemAccountID, "%v failed", "UpdateName")
		}
	})
}
