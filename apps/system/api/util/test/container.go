package test

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"path/filepath"
	"time"
)

const testContainerPort = "5432"
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

//func DriverSetupScriptFilePath(filePath string) string {
//	return filepath.Join("../../", "test/setup/driver", filePath)
//}

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
