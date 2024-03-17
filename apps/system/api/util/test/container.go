package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	rds "github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

const testContainerPort = "65432"
const testContainerDBUser = "postgres"
const testContainerDBPwd = "password"

var PsqlTestContainerConnQueryStr = fmt.Sprintf("user=%s&password=%s&sslmode=disable", testContainerDBUser, testContainerDBPwd)
var PsqlTestContainerConnStr = fmt.Sprintf("postgres://%s:%s?%s", "localhost", testContainerPort, PsqlTestContainerConnQueryStr)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func getProjectRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("runtime.Caller(0) failed to get current file path")
	}
	projectRoot := filepath.Dir(filename)
	for projectRoot != "" && filepath.Base(projectRoot) != "test" {
		projectRoot = filepath.Dir(projectRoot)
	}
	if projectRoot == "" {
		log.Fatal("Unable to find the project root directory")
	}
	return projectRoot
}

func createSystemTablesPath() string {
	return filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/001_create_tables.sql")
}

func createSystemBaseDataPath() []string {
	return []string{
		createSystemTablesPath(),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/099_setup_base_data.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/100_setup_update_workspace_detail.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/101_setup_update_role.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/102_setup_update_account_profile.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/103_setup_update_account_photo.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/104_setup_onboarding.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/105_setup_google_auth_name.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/106_setup_long_bio.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/107_setup_login_logout_retry.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/108_setup_invite.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/109_setup_invitation_has_event.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/110_setup_leave_workspace.sql"),
		filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/sql/111_setup_once_leave_workspace.sql"),
	}
}

func psqlTestContainerConfPath() string {
	return filepath.Join(getProjectRoot(), "../../../../../container/database/postgresql/primary/postgresql.conf")
}

func PSQLTestContainer(ctx context.Context, scripts ...string) (*PostgresContainer, error) {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:latest"),
		postgres.WithInitScripts(scripts...),
		postgres.WithConfigFile(psqlTestContainerConfPath()),
		postgres.WithDatabase("main"),
		postgres.WithUsername(testContainerDBUser),
		postgres.WithPassword(testContainerDBPwd),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}

func SetupTestDB(t *testing.T, ctx context.Context) *bun.DB {
	pgContainer, err := PSQLTestContainer(ctx, createSystemBaseDataPath()...)
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
	return db
}

func RedisTestContainer(ctx context.Context) (*redis.RedisContainer, error) {
	return redis.RunContainer(ctx,
		testcontainers.WithImage("redis:latest"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("Ready to accept connections").
				WithStartupTimeout(5*time.Second)),
		redis.WithSnapshotting(10, 1),
		redis.WithLogLevel(redis.LogLevelVerbose),
	)
}

func SetupRedisClient(t *testing.T, ctx context.Context) (*rds.Client, error) {
	rc, err := RedisTestContainer(ctx)
	if err != nil {
		return nil, err
	}

	endpoint, err := rc.Endpoint(ctx, "")
	if err != nil {
		return nil, err
	}

	rdb := rds.NewClient(&rds.Options{
		Addr: endpoint,
		DB:   0,
	})
	if err = rdb.FlushDB(ctx).Err(); err != nil {
		return nil, err
	}

	t.Cleanup(func() {
		if err = rdb.FlushDB(ctx).Err(); err != nil {
			t.Fatalf("failed to flush redis: %s", err)
		}
		if err = rc.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate redisContainer: %s", err)
		}
	})

	return rdb, nil
}
