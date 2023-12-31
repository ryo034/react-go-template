package core

import (
	"database/sql"
	"github.com/ryo034/react-go-template/packages/go/infrastructure/database/datasource"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Closable func()

func Initialize(mds datasource.DataSource, sds datasource.DataSource, isDebug bool, ops ...datasource.Option) Closable {
	mdb := mds.Open(ops...)
	if mdb == nil {
		panic("master db is nil")
	}
	boil.SetDB(mdb)
	// set debug mode
	boil.DebugMode = isDebug

	var sdb *sql.DB
	if reflect.DeepEqual(mds, sds) {
		sdb = mdb
	} else {
		sdb = sds.Open()
	}
	setReadDB(sdb)
	return func() {
		if err := mdb.Close(); err != nil {
			log.Println(err)
		}
		if sdb != mdb {
			if err := sdb.Close(); err != nil {
				log.Println(err)
			}
		}
	}
}
