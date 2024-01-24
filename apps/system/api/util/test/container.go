package test

import (
	"context"
	"database/sql"
	"fmt"
	rds "github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"path/filepath"
	"testing"
	"time"
)

const testContainerPort = "65432"
const testContainerDBPwd = "password"
const testContainerDBUser = "postgres"

var PsqlTestContainerConnQueryStr = fmt.Sprintf("user=%s&password=%s&sslmode=disable", testContainerDBUser, testContainerDBPwd)
var PsqlTestContainerConnStr = fmt.Sprintf("postgres://%s:%s?%s", "localhost", testContainerPort, PsqlTestContainerConnQueryStr)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

var CreateSystemTablesPath = filepath.Join("../../../../../", "container/database/postgresql/sql", "001_create_tables.sql")
var CreateSystemBaseDataPath = filepath.Join("../../../../../", "container/database/postgresql/sql", "099_initialize_data.sql")
var PsqlTestContainerConfPath = filepath.Join("../../../../../", "container/database/postgresql/primary", "postgresql.conf")

func PSQLTestContainer(ctx context.Context, scripts ...string) (*PostgresContainer, error) {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:latest"),
		postgres.WithInitScripts(scripts...),
		postgres.WithConfigFile(PsqlTestContainerConfPath),
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
	pgContainer, err := PSQLTestContainer(ctx, CreateSystemTablesPath, CreateSystemBaseDataPath)
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
	rc, err := redis.RunContainer(ctx,
		testcontainers.WithImage("docker.io/redis:7"),
		redis.WithSnapshotting(10, 1),
		redis.WithLogLevel(redis.LogLevelVerbose),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = rc.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
	return rc, nil
}

func SetupRedisClient(ctx context.Context) (*rds.Client, error) {
	rc, err := RedisTestContainer(ctx)
	host, err := rc.Host(ctx)
	if err != nil {
		return nil, err
	}
	rdb := rds.NewClient(&rds.Options{
		Addr: fmt.Sprintf("%s:%s", host, "6379"),
		DB:   0,
	})
	_ = rdb.FlushDB(ctx).Err()
	return rdb, nil
}
