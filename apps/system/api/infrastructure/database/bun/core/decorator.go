package core

import (
	"database/sql/driver"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryo034/react-go-template/apps/system/api/util/reflect/function"
)

func Decorate(fn interface{}, tx driver.Tx) function.AnyFunc {
	parsed := function.Parse(fn)
	abort := func(err error) function.Returns {
		if txErr := tx.Rollback(); txErr != nil {
			err = errors.Wrap(txErr, err.Error())
		}
		return function.ErrReturns(err)
	}
	return func() function.Returns {
		defer func() {
			if p := recover(); p != nil {
				if txErr := tx.Rollback(); txErr != nil {
					p = errors.Wrap(txErr, fmt.Sprint(p))
				}
				panic(p)
			}
		}()
		result := parsed()
		if err := result.Error(); err != nil {
			return abort(err)
		}
		if err := tx.Commit(); err != nil {
			return abort(err)
		}
		return result
	}
}
