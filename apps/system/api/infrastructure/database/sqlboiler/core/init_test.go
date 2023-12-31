package core

import (
	"database/sql"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/datasource"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	db1 *sql.DB = &sql.DB{}
	db2 *sql.DB = &sql.DB{}
)

type mockDataSource1 struct {
	db *sql.DB
}

func (m mockDataSource1) Open(_ ...datasource.Option) *sql.DB {
	return m.db
}

type mockDataSource2 struct {
	db *sql.DB
}

func (m mockDataSource2) Open(_ ...datasource.Option) *sql.DB {
	return m.db
}

func TestInitialize(t *testing.T) {
	tests := []struct {
		name    string
		master  datasource.DataSource
		slave   datasource.DataSource
		isDebug bool
		wantWdb boil.ContextExecutor
		wantRdb boil.ContextExecutor
	}{
		{"no slave without debug", mockDataSource1{db1}, mockDataSource1{db1}, false, db1, db1},
		{"has slave with debug", mockDataSource1{db1}, mockDataSource2{db2}, true, db1, db2},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Initialize(test.master, test.slave, test.isDebug)
			if got := readDB(); got != test.wantRdb {
				t.Errorf("readDB() = %v, want %v", got, test.wantRdb)
			}
			if got := boil.GetContextDB(); got != test.wantWdb {
				t.Errorf("boil.GetContextDB() = %v, want %v", got, test.wantWdb)
			}
			if got := boil.DebugMode; got != test.isDebug {
				t.Errorf("boil.DebugMode = %v, want %v", got, test.isDebug)
			}
		})
	}
}
