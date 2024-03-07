//go:build testcontainers

package me

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	invitationDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

const systemAccountID = "394e67b6-2850-4ddf-a4c9-c2a619d5bf70"

var accountIDUUID = uuid.MustParse(systemAccountID)

func Test_driver_Find_OK(t *testing.T) {
	defaultTime, _ := time.Parse("2006-01-02 15:04:05", "2024-01-10 12:00:00")
	memberID := uuid.MustParse("377eba35-5560-4f48-a99d-19cbd6a82b0d")

	workspaceID := uuid.MustParse("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
	wID := workspace.NewIDFromUUID(workspaceID)
	anID := uuid.MustParse("018e088e-fd36-722d-a927-8cfd34a642bd")
	aeID := uuid.MustParse("018e09c2-9924-7048-9f08-afa2f3ea5b53")

	mrID := uuid.MustParse("018df76b-260d-759f-9b47-fb5f611f5da6")

	want := &models.Member{
		MemberID:    memberID,
		WorkspaceID: wID.Value(),
		AccountID:   accountIDUUID,
		CreatedAt:   defaultTime,
		Role: &models.MemberLatestRole{
			MemberRoleID: mrID,
			MemberID:     memberID,
			MemberRole: &models.MemberRole{
				MemberRoleID: mrID,
				MemberID:     memberID,
				Role:         "owner",
				AssignedAt:   defaultTime,
				AssignedBy:   memberID,
			},
		},
		Account: &models.Account{
			AccountID: accountIDUUID,
			CreatedAt: defaultTime,
			Name: &models.AccountLatestName{
				AccountNameID: anID,
				AccountID:     accountIDUUID,
				AccountName: &models.AccountName{
					AccountNameID: anID,
					AccountID:     accountIDUUID,
					Name:          "John Doe",
					CreatedAt:     defaultTime,
				},
			},
			Email: &models.AccountLatestEmail{
				AccountEmailID: aeID,
				AccountID:      accountIDUUID,
				AccountEmail: &models.AccountEmail{
					AccountEmailID: aeID,
					AccountID:      accountIDUUID,
					Email:          "account@example.com",
					CreatedAt:      defaultTime,
				},
			},
			AuthProviders: []*models.AuthProvider{
				{
					AuthProviderID: uuid.MustParse("018de2f6-968d-7458-9c67-69ae5698a143"),
					ProviderUID:    "394e67b6-2850-4ddf-a4c9-c2a619d5bf70",
					AccountID:      accountIDUUID,
					Provider:       "email",
					ProvidedBy:     "firebase",
					RegisteredAt:   defaultTime,
				},
			},
			PhoneNumber: nil,
		},
		Profile: &models.MemberProfile{
			MemberID:       memberID,
			MemberIDNumber: "DEV-12345",
			DisplayName:    "John Doe",
			Bio:            "bio",
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
		got, err := NewDriver(invitationDr.NewDriver()).Find(ctx, pr.GetExecutor(ctx, true), member.NewIDFromUUID(memberID))
		if (err != nil) != wantErr {
			t.Errorf("Find() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			assert.EqualValuesf(t, want, got, "%v failed", "Find")
		}
	})
}
