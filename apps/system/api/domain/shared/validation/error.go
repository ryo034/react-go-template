package validation

import (
	"fmt"

	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
)

type Error interface {
	error
	MessageKey() domainError.MessageKey
	Code() string
	Args() []interface{}
}

type err struct {
	messageKey domainError.MessageKey
	code       domainError.Code
	args       []interface{}
}

func (e *err) Error() string {
	return fmt.Sprintf("Validation failed (key:%v, args: %v)", e.messageKey, e.args)
}

func (e *err) MessageKey() domainError.MessageKey {
	return e.messageKey
}

func (e *err) Code() string {
	return string(e.code)
}

func (e *err) Args() []interface{} {
	return e.args
}

type Errors interface {
	Append(messageKey domainError.MessageKey, code *domainError.Code, args ...interface{})
	AppendAll(errs Errors)
	NilIfEmpty() error
	AsSlice() []error
	Size() int
	IsEmpty() bool
	IsNotEmpty() bool
	error
}

type errors struct {
	errors []error
}

func NewErrors() Errors {
	return &errors{
		errors: make([]error, 0, 10),
	}
}

func (v *errors) NilIfEmpty() error {
	if v.IsEmpty() {
		return nil
	}
	return v
}

func (v *errors) AppendAll(errs Errors) {
	v.errors = append(v.errors, errs.AsSlice()...)
}

func (v *errors) Append(messageKey domainError.MessageKey, code *domainError.Code, args ...interface{}) {
	c := domainError.BasicCodeKey
	if code != nil {
		c = *code
	}
	v.errors = append(v.errors, &err{messageKey, c, args})
}

func (v *errors) AsSlice() []error {
	return v.errors
}

func (v *errors) Size() int {
	return len(v.errors)
}

func (v *errors) IsEmpty() bool {
	return v.Size() == 0
}

func (v *errors) IsNotEmpty() bool {
	return !v.IsEmpty()
}

func (v *errors) Error() string {
	return fmt.Sprint(v.errors)
}
