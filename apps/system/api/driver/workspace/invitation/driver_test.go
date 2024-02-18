package invitation

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"reflect"
	"testing"
	"time"
)

func Test_driver_FindByToken_OK(t *testing.T) {
	defaultTime := test.GetDefaultTime()
	wantErr := false

	wID := uuid.MustParse("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
	invitationID := uuid.MustParse("018d96b8-2211-7862-bcbe-e9f4d002a8fc")
	mID := uuid.MustParse("377eba35-5560-4f48-a99d-19cbd6a82b0d")
	inviteeSystemAccountID := uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	token := uuid.MustParse("018d96b7-df68-792f-97d0-d6a044c2b4a2")

	invitedBy := &models.Member{
		MemberID:        mID,
		WorkspaceID:     wID,
		SystemAccountID: inviteeSystemAccountID,
		CreatedAt:       defaultTime,
		SystemAccount: &models.SystemAccount{
			SystemAccountID: inviteeSystemAccountID,
			CreatedAt:       defaultTime,
			PhoneNumber:     nil,
			Profile: &models.SystemAccountProfile{
				SystemAccountID: inviteeSystemAccountID,
				Name:            "John Doe",
				Email:           "system_account@example.com",
				CreatedAt:       defaultTime,
				UpdatedAt:       defaultTime,
			},
		},
		Profile: &models.MemberProfile{
			MemberID:       mID,
			MemberIDNumber: "DEV-12345",
			DisplayName:    "John Doe",
			CreatedAt:      defaultTime,
			UpdatedAt:      defaultTime,
			Member:         nil,
		},
		Workspace: nil,
	}

	invUnit := &models.InvitationUnit{
		InvitationUnitID: uuid.MustParse("018db4a4-c350-747b-8c4f-bd827e08174b"),
		WorkspaceID:      wID,
		InvitedBy:        invitedBy.MemberID,
		CreatedAt:        defaultTime,
		Workspace: &models.Workspace{
			WorkspaceID: wID,
			CreatedAt:   defaultTime,
			Detail: &models.WorkspaceDetail{
				WorkspaceID: wID,
				Name:        "Example",
				Subdomain:   "example",
				CreatedAt:   defaultTime,
				UpdatedAt:   defaultTime,
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
		//got, err := NewDriver().FindActiveByToken(ctx, pr.GetExecutor(ctx, true), token)
		em, _ := account.NewEmail("invite_test_not_expired@example.com")
		got, err := NewDriver().FindActiveAllByEmail(ctx, pr.GetExecutor(ctx, true), em)
		if (err != nil) != wantErr {
			t.Errorf("FindActiveByToken() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			assert.EqualValuesf(t, want, got, "%v failed", "Find")
		}
	})
}
