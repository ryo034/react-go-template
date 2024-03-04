//go:build testcontainers

package auth

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Test_driver_Create_OK(t *testing.T) {
	var systemAccountIDUUID = uuid.MustParse("1710ba5a-f82e-41f7-a599-71693b99848d")
	st := time.Now()

	email, _ := account.NewEmail("Test_driver_Create_OK@example.com")

	wantErr := false
	ctx := context.Background()
	t.Run("Create", func(t *testing.T) {
		db := bun.NewDB(test.SetupTestDB(t, ctx).DB, pgdialect.New())
		pr := core.NewDatabaseProvider(db, db)
		aID := account.NewIDFromUUID(systemAccountIDUUID)
		apUID, _ := provider.NewUID(aID.ToString())
		ap := provider.NewProvider(
			provider.NewIDFromUUID(uuid.MustParse("018de777-7bb8-7cb7-b705-58c876746288")),
			"email",
			"firebase",
			apUID,
		)
		usr := user.NewUser(aID, email, nil, nil, nil)
		got, err := NewDriver().Create(ctx, pr.GetExecutor(ctx, false), usr, ap)
		if (err != nil) != wantErr {
			t.Errorf("Create() error = %v, wantErr %v", err, wantErr)
			return
		}
		et := time.Now()
		assert.True(t, got.CreatedAt.After(st) && got.CreatedAt.Before(et), "CreatedAt should be within test time range")
		if got.Name != nil {
			assert.True(t, got.Name.AccountName.CreatedAt.After(st) && got.Name.AccountName.CreatedAt.Before(et), "Name.CreatedAt should be within test time range")
		}
		assert.Equal(t, systemAccountIDUUID, got.AccountID)
		assert.Equal(t, "Test_driver_Create_OK@example.com", got.Email.AccountEmail.Email)
		if got.Name != nil {
			t.Errorf("Name should be nil, got: %v", got.Name)
		}
		if got.PhoneNumber != nil {
			t.Errorf("PhoneNumber should be nil, got: %v", got.PhoneNumber)
		}
		if got.AuthProviders == nil {
			t.Errorf("AuthProvider should not be nil")
		}
		assert.Equal(t, "email", got.AuthProviders[0].Provider)
		assert.Equal(t, "1710ba5a-f82e-41f7-a599-71693b99848d", got.AuthProviders[0].ProviderUID)
		assert.Equal(t, "firebase", got.AuthProviders[0].ProvidedBy)
	})
}

func Test_driver_Find_OK(t *testing.T) {
	systemAccountIDUUID := uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	defaultTime := test.GetDefaultTime()
	email, _ := account.NewEmail("account@example.com")
	anID := uuid.MustParse("018e088e-fd36-722d-a927-8cfd34a642bd")
	aeID := uuid.MustParse("018e09c2-9924-7048-9f08-afa2f3ea5b53")

	want := &models.Account{
		AccountID: systemAccountIDUUID,
		CreatedAt: defaultTime,
		//PhoneNumbers: []*models.AccountPhoneNumber{
		//	{
		//		AccountID: systemAccountIDUUID,
		//		PhoneNumber:     "09012345678",
		//		CreatedAt:       defaultTime,
		//	},
		//},
		Name: &models.AccountLatestName{
			AccountNameID: anID,
			AccountID:     systemAccountIDUUID,
			AccountName: &models.AccountName{
				AccountNameID: anID,
				AccountID:     systemAccountIDUUID,
				Name:          "John Doe",
				CreatedAt:     defaultTime,
			},
		},
		Email: &models.AccountLatestEmail{
			AccountEmailID: aeID,
			AccountID:      systemAccountIDUUID,
			AccountEmail: &models.AccountEmail{
				AccountEmailID: aeID,
				AccountID:      systemAccountIDUUID,
				Email:          "account@example.com",
				CreatedAt:      defaultTime,
			},
		},
		AuthProviders: []*models.AuthProvider{
			{
				AuthProviderID: uuid.MustParse("018de2f6-968d-7458-9c67-69ae5698a143"),
				ProviderUID:    "394e67b6-2850-4ddf-a4c9-c2a619d5bf70",
				AccountID:      systemAccountIDUUID,
				Provider:       "email",
				ProvidedBy:     "firebase",
				RegisteredAt:   defaultTime,
			},
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
