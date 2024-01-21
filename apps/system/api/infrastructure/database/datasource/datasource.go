package datasource

import (
	"database/sql"
	"time"
)

type DataSource interface {
	Open(ops ...Option) *sql.DB
}

type Option interface {
	apply(db *sql.DB)
}

func MaxOpenConns(val int) Option {
	return maxOpenConns{val}
}

type maxOpenConns struct {
	val int
}

func (o maxOpenConns) apply(db *sql.DB) {
	db.SetMaxOpenConns(o.val)
}

type maxIdleConns struct {
	val int
}

func (o maxIdleConns) apply(db *sql.DB) {
	db.SetMaxIdleConns(o.val)
}

func ConnMaxLifetime(val time.Duration) Option {
	return connMaxLifetime{val}
}

type connMaxLifetime struct {
	val time.Duration
}

func (o connMaxLifetime) apply(db *sql.DB) {
	db.SetConnMaxLifetime(o.val)
}
