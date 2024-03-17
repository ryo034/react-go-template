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

	wID := uuid.MustParse("018e4922-563a-7731-b389-c2a9ac0d97e9")
	invitationID := uuid.MustParse("018e4922-563a-7566-bc5e-65dc2f8faefe")
	mID := uuid.MustParse("018e4922-563a-7807-b01f-2e630e4d22e9")
	inviterAccountID := uuid.MustParse("018e4922-563a-7097-bbdb-ffa9f74da283")
	token := uuid.MustParse("018e4922-563a-735c-a715-2fe940d327cf")

	anID := uuid.MustParse("018e4922-563a-7cff-bc38-66e095586aa0")
	aeID := uuid.MustParse("018e4922-563a-7f98-b31e-e42efc811159")

	mpID := uuid.MustParse("018e4922-563a-71f9-a7a1-25455276884f")
	mrID := uuid.MustParse("018e4922-563a-7abc-a66a-24ea32938e9c")
	wdID := uuid.MustParse("018e4922-563a-736e-b632-b32b7df08c67")

	msheID := uuid.MustParse("018e4922-563a-7b08-8fb3-955ae030f5d1")

	invitedBy := &models.Member{
		MemberID:    mID,
		WorkspaceID: wID,
		AccountID:   inviterAccountID,
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
			AccountID: inviterAccountID,
			CreatedAt: defaultTime,
			Name: &models.AccountLatestName{
				AccountNameID: anID,
				AccountID:     inviterAccountID,
				AccountName: &models.AccountName{
					AccountNameID: anID,
					AccountID:     inviterAccountID,
					Name:          "Invite TestHasEvent",
					CreatedAt:     defaultTime,
				},
			},
			Email: &models.AccountLatestEmail{
				AccountEmailID: aeID,
				AccountID:      inviterAccountID,
				AccountEmail: &models.AccountEmail{
					AccountEmailID: aeID,
					AccountID:      inviterAccountID,
					Email:          "invite_test_has_event_inviter@example.com",
					CreatedAt:      defaultTime,
				},
			},
			AuthProviders: []*models.AuthProvider{
				{
					AuthProviderID: uuid.MustParse("018e4922-563a-71b5-bdb4-ba06511f2590"),
					ProviderUID:    "018e4922-563a-7097-bbdb-ffa9f74da283",
					AccountID:      inviterAccountID,
					Provider:       "email",
					ProvidedBy:     "firebase",
					RegisteredAt:   defaultTime,
				},
			},
		},
		Profile: &models.MemberLatestProfile{
			MemberID:        mID,
			MemberProfileID: mpID,
			MemberProfile: &models.MemberProfile{
				MemberProfileID: mpID,
				MemberID:        mID,
				MemberIDNumber:  "DEV-12345",
				DisplayName:     "Invite TestHasEvent",
				Bio:             "bio",
				CreatedAt:       defaultTime,
			},
		},
		MembershipEvent: &models.LatestMembershipEvent{
			MembershipEventID: msheID,
			MemberID:          mID,
			MembershipEvent: &models.MembershipEvent{
				MembershipEventID: msheID,
				MemberID:          mID,
				EventType:         "join",
				CreatedBy:         mID,
				EventAt:           defaultTime,
			},
		},
		Workspace: nil,
	}

	invUnit := &models.InvitationUnit{
		InvitationUnitID: uuid.MustParse("018e4922-563a-7c76-8ea0-815441c038cb"),
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
					Name:              "Invite TestHasEvent",
					Subdomain:         "invite-test-has-event",
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

	invtID := uuid.MustParse("018e493a-1b7f-79f3-83f9-c2ee307ce23d")

	want := &models.Invitation{
		InvitationID:     invitationID,
		InvitationUnitID: invUnit.InvitationUnitID,
		InvitationUnit:   invUnit,
		InviteeName:      nil,
		Invitee:          invitee,
		Event:            nil,
		Token: &models.LatestInvitationToken{
			InvitationTokenID: invtID,
			InvitationID:      invitationID,
			InvitationToken: &models.InvitationToken{
				InvitationTokenID: invtID,
				InvitationID:      invitationID,
				Token:             token,
				ExpiredAt:         expTime,
				CreatedAt:         defaultTime,
			},
		},
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
