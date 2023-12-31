package config

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/datasource"
)

const (
	dbUser        Key = "DB_USER"
	dbPass        Key = "DB_PASS"
	dbName        Key = "DB_NAME"
	dbSourceHost  Key = "DB_SOURCE_HOST"
	dbSourcePort  Key = "DB_SOURCE_PORT"
	dbReplicaHost Key = "DB_REPLICA_HOST"
	dbReplicaPort Key = "DB_REPLICA_PORT"
)

func (r *reader) SourceDataSource() datasource.DataSource {
	return datasource.NewPostgresqlDataSource(
		r.fromEnv(dbSourceHost),
		r.fromEnvUint(dbSourcePort),
		r.fromEnv(dbName),
		r.fromEnv(dbUser),
		r.fromEnv(dbPass),
	)
}

func (r *reader) ReplicaDataSource() datasource.DataSource {
	return datasource.NewPostgresqlDataSource(
		r.fromEnv(dbReplicaHost),
		r.fromEnvUint(dbReplicaPort),
		r.fromEnv(dbName),
		r.fromEnv(dbUser),
		r.fromEnv(dbPass),
	)
}
