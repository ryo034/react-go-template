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
			Profile: &models.SystemAccountProfile{
				SystemAccountID: systemAccountIDUUID,
				Name:            "John Doe",
				CreatedAt:       defaultTime,
				UpdatedAt:       defaultTime,
			},
			Emails: []*models.SystemAccountEmail{
				{
					SystemAccountID: systemAccountIDUUID,
					Email:           "system_account@example.com",
					CreatedAt:       defaultTime,
				},
			},
			AuthProviders: []*models.AuthProvider{
				{
					AuthProviderID:  uuid.MustParse("018de2f6-968d-7458-9c67-69ae5698a143"),
					SystemAccountID: systemAccountIDUUID,
					ProviderUID:     "394e67b6-2850-4ddf-a4c9-c2a619d5bf70",
					Provider:        "email",
					ProvidedBy:      "firebase",
					CreatedAt:       defaultTime,
				},
			},
			PhoneNumbers: nil,
		},
		Profile: &models.MemberProfile{
			MemberID:       memberID,
			MemberIDNumber: "DEV-12345",
			DisplayName:    "John Doe",
			Bio:            "John Doe is a passionate software engineer with 8 years of experience specializing in web development, particularly with React and Node.js. A graduate from MIT with a strong focus on clean architecture and Agile methodologies, John has successfully led multiple projects, from innovative startups to established tech giants. He's a firm believer in continuous learning, contributing regularly to open-source projects, and sharing insights through tech blogs and meetups. Outside of work, John enjoys hiking 🚶‍♂️, drone photography 📸, and playing the guitar 🎸. He's committed to using technology to drive positive social change.",
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
