//go:build testcontainers

package auth

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
	"reflect"
	"testing"
	"time"
)

func Test_driver_Create_OK(t *testing.T) {
	var systemAccountIDUUID = uuid.MustParse("1710ba5a-f82e-41f7-a599-71693b99848d")
	st := time.Now()

	email, _ := account.NewEmail("system_account@example.com")

	wantErr := false
	ctx := context.Background()
	t.Run("Find", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		pr := core.NewDatabaseProvider(db, db)
		aID := account.NewIDFromUUID(systemAccountIDUUID)
		got, err := NewDriver().Create(ctx, pr.GetExecutor(ctx, false), aID, email)
		if (err != nil) != wantErr {
			t.Errorf("Create() error = %v, wantErr %v", err, wantErr)
			return
		}
		et := time.Now()
		assert.True(t, got.CreatedAt.After(st) && got.CreatedAt.Before(et), "CreatedAt should be within test time range")
		if got.Profile != nil {
			assert.True(t, got.Profile.CreatedAt.After(st) && got.Profile.CreatedAt.Before(et), "Profile.CreatedAt should be within test time range")
			assert.True(t, got.Profile.UpdatedAt.After(st) && got.Profile.UpdatedAt.Before(et), "Profile.UpdatedAt should be within test time range")
		}

		assert.Equal(t, systemAccountIDUUID, got.SystemAccountID)
		assert.Equal(t, "system_account@example.com", got.Profile.Email)
		assert.Equal(t, "", got.Profile.Name)
		if got.PhoneNumber != nil {
			t.Errorf("PhoneNumber should be nil, got: %v", got.PhoneNumber)
		}
	})
}

func Test_driver_Find_OK(t *testing.T) {
	var systemAccountIDUUID = uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	defaultTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-10 12:00:00")
	if err != nil {
		t.Fatalf("failed to parse defaultTime: %v", err)
	}
	email, _ := account.NewEmail("system_account@example.com")

	want := &models.SystemAccount{
		SystemAccountID: systemAccountIDUUID,
		CreatedAt:       defaultTime,
		PhoneNumber: &models.SystemAccountPhoneNumber{
			SystemAccountID: systemAccountIDUUID,
			PhoneNumber:     "09012345678",
			CreatedAt:       defaultTime,
			UpdatedAt:       defaultTime,
		},
		Profile: &models.SystemAccountProfile{
			SystemAccountID: systemAccountIDUUID,
			Name:            "John Doe",
			Email:           "system_account@example.com",
			CreatedAt:       defaultTime,
			UpdatedAt:       defaultTime,
		},
	}
	wantErr := false
	ctx := context.Background()
	t.Run("Find", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		pr := core.NewDatabaseProvider(db, db)
		got, err := NewDriver().Find(ctx, pr.GetExecutor(ctx, true), email)
		if (err != nil) != wantErr {
			t.Errorf("Find() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			assert.EqualValuesf(t, want, got, "%v failed", "Find")
		}
	})
}
