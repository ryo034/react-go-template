//go:build testcontainers

package invitation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func Test_driver_FindByToken_OK(t *testing.T) {
	defaultTime := test.GetDefaultTime()
	wantErr := false

	wID := uuid.MustParse("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
	invitationID := uuid.MustParse("018d96b8-2211-7862-bcbe-e9f4d002a8fc")
	mID := uuid.MustParse("377eba35-5560-4f48-a99d-19cbd6a82b0d")
	inviteeAccountID := uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	token := uuid.MustParse("018d96b7-df68-792f-97d0-d6a044c2b4a2")

	anID := uuid.MustParse("018e088e-fd36-722d-a927-8cfd34a642bd")
	aeID := uuid.MustParse("018e09c2-9924-7048-9f08-afa2f3ea5b53")

	mrID := uuid.MustParse("018df76b-260d-759f-9b47-fb5f611f5da6")

	invitedBy := &models.Member{
		MemberID:    mID,
		WorkspaceID: wID,
		AccountID:   inviteeAccountID,
		CreatedAt:   defaultTime,
		Role: &models.MemberLatestRole{
			MemberRoleID: mrID,
			MemberID:     mID,
			MemberRole: &models.MemberRole{
				MemberRoleID: mrID,
				MemberID:     mID,
				Role:         "owner",
				AssignedAt:   defaultTime,
				AssignedBy:   mID,
			},
		},
		Account: &models.Account{
			AccountID: inviteeAccountID,
			CreatedAt: defaultTime,
			Name: &models.AccountLatestName{
				AccountNameID: anID,
				AccountID:     inviteeAccountID,
				AccountName: &models.AccountName{
					AccountNameID: anID,
					AccountID:     inviteeAccountID,
					Name:          "John Doe",
					CreatedAt:     defaultTime,
				},
			},
			Email: &models.AccountLatestEmail{
				AccountEmailID: aeID,
				AccountID:      inviteeAccountID,
				AccountEmail: &models.AccountEmail{
					AccountEmailID: aeID,
					AccountID:      inviteeAccountID,
					Email:          "account@example.com",
					CreatedAt:      defaultTime,
				},
			},
			AuthProviders: []*models.AuthProvider{
				{
					AuthProviderID: uuid.MustParse("018de2f6-968d-7458-9c67-69ae5698a143"),
					ProviderUID:    "394e67b6-2850-4ddf-a4c9-c2a619d5bf70",
					AccountID:      inviteeAccountID,
					Provider:       "email",
					ProvidedBy:     "firebase",
					RegisteredAt:   defaultTime,
				},
			},
		},
		Profile: &models.MemberProfile{
			MemberID:       mID,
			MemberIDNumber: "DEV-12345",
			DisplayName:    "John Doe",
			Bio:            "bio",
			CreatedAt:      defaultTime,
			UpdatedAt:      defaultTime,
			Member:         nil,
		},
		Workspace: nil,
	}

	wdID := uuid.MustParse("018e200b-9d01-70ed-8c5a-5a5df2a98f11")

	invUnit := &models.InvitationUnit{
		InvitationUnitID: uuid.MustParse("018db4a4-c350-747b-8c4f-bd827e08174b"),
		WorkspaceID:      wID,
		InvitedBy:        invitedBy.MemberID,
		CreatedAt:        defaultTime,
		Workspace: &models.Workspace{
			WorkspaceID: wID,
			CreatedAt:   defaultTime,
			Detail: &models.WorkspaceLatestDetail{
				WorkspaceDetailID: wdID,
				WorkspaceID:       wID,
				WorkspaceDetail: &models.WorkspaceDetail{
					WorkspaceID:       wID,
					WorkspaceDetailID: wdID,
					Name:              "Example",
					Subdomain:         "example",
					CreatedAt:         defaultTime,
				},
			},
			Members: nil,
		},
		Member:      invitedBy,
		Invitations: nil,
	}

	expTime, _ := time.Parse("2006-01-02 15:04:05", "2200-01-10 12:00:00")

	tokens := make([]*models.InvitationToken, 0, 1)
	tokens = append(tokens, &models.InvitationToken{
		InvitationID: invitationID,
		Token:        token,
		ExpiredAt:    expTime,
		CreatedAt:    defaultTime,
	})

	invitee := &models.Invitee{
		InvitationID: invitationID,
		Email:        "invite_test_not_expired@example.com",
		Invitation:   nil,
	}

	want := &models.Invitation{
		InvitationID:     invitationID,
		InvitationUnitID: invUnit.InvitationUnitID,
		InvitationUnit:   invUnit,
		InviteeName:      nil,
		Invitee:          invitee,
		Events:           nil,
		Tokens:           tokens,
	}

	ctx := context.Background()
	t.Run("FindActiveByToken", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
		pr := core.NewDatabaseProvider(db, db)
		got, err := NewDriver().FindActiveByToken(ctx, pr.GetExecutor(ctx, true), invitation.NewToken(token))
		if (err != nil) != wantErr {
			t.Errorf("FindActiveByToken() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			assert.EqualValuesf(t, want, got, "%v failed", "Find")
		}
	})
}
