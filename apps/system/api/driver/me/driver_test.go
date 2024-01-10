package me

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/models"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"reflect"
	"testing"
	"time"
)

// Check DB
// psql -h localhost -p 5432 -U postgres -d main
// \dt

const systemAccountID = "394e67b6-2850-4ddf-a4c9-c2a619d5bf70"

var systemAccountIDUUID = uuid.MustParse(systemAccountID)

func Test_driver_Find_OK(t *testing.T) {
	defaultTime, err := time.Parse("2006-01-02 15:04:05 -0700", "2024-01-10 12:00:00 +0000")
	if err != nil {
		t.Fatalf("failed to parse defaultTime: %v", err)
	}
	accountID, _ := account.NewID(systemAccountID)
	employeeID := uuid.MustParse("377eba35-5560-4f48-a99d-19cbd6a82b0d")

	want := &models.Employee{
		EmployeeID:      employeeID,
		SystemAccountID: systemAccountIDUUID,
		OrganizationID:  uuid.MustParse("c1bd2603-b9cd-4f84-8b83-3548f6ae150b"),
		CreatedAt:       defaultTime,
		SystemAccount: &models.SystemAccount{
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
				Email:           "system_account@example.com",
				EmailVerified:   true,
				CreatedAt:       defaultTime,
				UpdatedAt:       defaultTime,
			},
		},
		Profile: &models.EmployeeProfile{
			EmployeeID:       employeeID,
			EmployeeIDNumber: "EMP-12345",
			Name:             "John Doe",
			CreatedAt:        defaultTime,
			UpdatedAt:        defaultTime,
		},
		Hires:       nil,
		Separations: nil,
	}
	wantErr := false
	ctx := context.Background()
	t.Run("Find", func(t *testing.T) {
		pgContainer, err := test.PSQLTestContainer(ctx, test.CreateSystemTablesPath, test.CreateSystemBaseDataPath)
		if err != nil {
			t.Fatalf("failed to PSQLContainer creation: %v", err)
		}

		sqlDB, err := sql.Open("postgres", pgContainer.ConnectionString)
		if err != nil {
			t.Fatalf("failed to sql.Open: %v", err)
		}

		db := bun.NewDB(sqlDB, pgdialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

		t.Cleanup(func() {
			if err = db.Close(); err != nil {
				t.Fatalf("failed to close db: %s", err)
			}
			if err = pgContainer.Terminate(ctx); err != nil {
				t.Fatalf("failed to terminate pgContainer: %s", err)
			}
		})

		got, err := NewDriver(db).Find(ctx, accountID)
		if (err != nil) != wantErr {
			t.Errorf("Find() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			assert.EqualValuesf(t, want, got, "%v failed", "Find")
		}
	})
}
