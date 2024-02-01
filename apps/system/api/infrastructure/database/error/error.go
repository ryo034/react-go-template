package error

import (
	"github.com/go-faster/errors"
	"github.com/jackc/pgconn"
)

func IsDuplicateError(e error) bool {
	var err *pgconn.PgError
	if errors.As(e, &err) {
		return err.Code == "23505"
	}
	return false
}
