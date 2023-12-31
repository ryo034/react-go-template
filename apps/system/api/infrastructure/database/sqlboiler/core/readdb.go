package core

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	// currentDB is a global database handle for the package
	currentDB boil.ContextExecutor
)

// initializes the database handle for all template db interactions
func setReadDB(db boil.ContextExecutor) {
	currentDB = db
}

// retrieves the global state database handle
func readDB() boil.ContextExecutor {
	return currentDB
}
