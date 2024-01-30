//go:build testcontainers

package workspace

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

func Test_driver_FindAll_OK(t *testing.T) {
	var systemAccountIDUUID = uuid.MustParse("394e67b6-2850-4ddf-a4c9-c2a619d5bf70")
	aID := account.NewIDFromUUID(systemAccountIDUUID)
	defaultTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-10 12:00:00")
	if err != nil {
		t.Fatalf("failed to parse defaultTime: %v", err)
	}
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
