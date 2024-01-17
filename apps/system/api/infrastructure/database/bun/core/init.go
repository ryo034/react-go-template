package core

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/datasource"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"reflect"
)

type Closable func()

func Initialize(pds datasource.DataSource, rds datasource.DataSource, isDebug bool, ops ...datasource.Option) (Closable, Provider, TransactionProvider) {
	pdb := pds.Open(ops...)
	if pdb == nil {
		panic("master db is nil")
	}

	hook := bundebug.NewQueryHook(bundebug.WithVerbose(isDebug))

	priDB := bun.NewDB(pdb, pgdialect.New())
	priDB.AddQueryHook(hook)

	var sdb *sql.DB
	if reflect.DeepEqual(pds, rds) {
		sdb = pdb
	}
	sdb = rds.Open()

	repDB := bun.NewDB(sdb, pgdialect.New())
	repDB.AddQueryHook(hook)

	pr := NewDatabaseProvider(priDB, repDB)
	txp := NewTransactionProvider(priDB)

	return func() {
		if err := pdb.Close(); err != nil {
			log.Println(err)
		}
		if sdb != pdb {
			if err := sdb.Close(); err != nil {
				log.Println(err)
			}
		}
	}, pr, txp
}
