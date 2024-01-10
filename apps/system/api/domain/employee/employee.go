package employee

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
)

type Employee struct {
	u        *user.User
	idNumber string
}

func NewEmployee(u *user.User, idNumber string) *Employee {
	return &Employee{u, idNumber}
}

func (e *Employee) User() *user.User {
	return e.u
}

func (e *Employee) IDNumber() string {
	return e.idNumber
}
