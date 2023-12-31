package error

import (
	"database/sql"

	"github.com/friendsofgo/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	DuplicateErrorCode            = 1022
	UniqueConstraintErrorCode     = 1062
	ForeignKeyConstraintErrorCode = 1452
)

func IsNoSuchDataError(err error) bool {
	return err == sql.ErrNoRows
}

func IsDuplicateError(err error) bool {
	return isMysqlError(err, DuplicateErrorCode)
}

func IsUniqueConstraintError(err error) bool {
	return isMysqlError(err, UniqueConstraintErrorCode)
}

func IsForeignKeyConstraintError(err error) bool {
	return isMysqlError(err, ForeignKeyConstraintErrorCode)
}

func isMysqlError(err error, errorCode uint16) bool {
	if err == nil {
		return false
	}
	num, ok := extractMysqlErrorCode(err)
	return ok && (num == errorCode)
}

func extractMysqlErrorCode(err error) (uint16, bool) {
	cause := errors.Cause(err)
	mErr, ok := cause.(*mysql.MySQLError)
	if ok {
		return mErr.Number, ok
	}
	return 0, ok
}
