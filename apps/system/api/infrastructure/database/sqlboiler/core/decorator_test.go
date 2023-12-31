package core

import (
	"errors"
	"fmt"
	"testing"
)

type mockTx struct {
	rollbacked         bool
	committed          bool
	commitReturnsError bool
}

func (m *mockTx) Rollback() error {
	m.rollbacked = true
	return nil
}

var commitError = errors.New("commit returns error")

func (m *mockTx) Commit() error {
	m.committed = true
	if m.commitReturnsError {
		return commitError
	}
	return nil
}

var funcError = errors.New("commit returns error")
var funcReturn = "success"

func mockFunc(returnsError bool) func() (string, error) {
	return func() (string, error) {
		if returnsError {
			return "", funcError
		}
		return funcReturn, nil
	}
}

func TestDecorate(t *testing.T) {
	testCases := map[string]struct {
		funcReturnsError    bool
		commitReturnsError  bool
		expectedError       error
		expectedRollebacked bool
		expectedCommited    bool
	}{"all fine": {false, false, nil, false, true},
		"commit fails":   {false, true, commitError, true, true},
		"function fails": {true, false, funcError, true, false},
		"both fails":     {true, true, funcError, true, false}}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mtx := mockTx{
				rollbacked:         false,
				committed:          false,
				commitReturnsError: tc.commitReturnsError,
			}
			fn := mockFunc(tc.funcReturnsError)
			returns := Decorate(fn, &mtx)()
			err := returns.Error()
			// if error is nil
			if tc.expectedError == nil {
				// then value should be funcReturn.
				actual := returns.Value(0)
				if actual != funcReturn {
					t.Errorf("actual return %s should be %s", actual, funcError)
				}
				// if error is not nil
			} else {
				// then error should be expectedError
				if err != tc.expectedError {
					fmt.Print(mtx.Commit().Error())
					t.Errorf("actual error %s should be %s", err, funcError)
				}
			}

			if mtx.rollbacked != tc.expectedRollebacked {
				t.Errorf("rollebacked should be %t", tc.expectedRollebacked)
			}
			if mtx.committed != tc.expectedCommited {
				t.Errorf("committed should be %t", tc.expectedCommited)
			}
		})
	}
}
