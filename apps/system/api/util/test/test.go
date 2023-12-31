package test

import (
	"reflect"
	"sort"
	"testing"

	domainError "github.com/ryo034/react-go-template/packages/go/domain/shared/error"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/validation"
)

func PanicToFatal(t *testing.T) {
	if p := recover(); p != nil {
		t.Fatal(p)
	}
}

type ValidationErrorBuilder struct {
	Errs validation.Errors
}

func (v *ValidationErrorBuilder) Append(messageKey domainError.MessageKey, args ...interface{}) *ValidationErrorBuilder {
	v.Errs.Append(messageKey, args...)
	return v
}

func NewValidationErrors(messageKey domainError.MessageKey, args ...interface{}) *ValidationErrorBuilder {
	errs := validation.NewErrors()
	errs.Append(messageKey, args...)
	return &ValidationErrorBuilder{errs}
}

func ValidationErrorEquals(v1, v2 validation.Errors) bool {
	if v1 == v2 {
		return true
	}
	if v1.Size() != v2.Size() {
		return false
	}
	es1 := v1.AsSlice()
	es2 := v2.AsSlice()
	// ignore order
	sort.Slice(es1, func(i, j int) bool {
		return es1[i].Error() < es1[j].Error()
	})
	sort.Slice(es2, func(i, j int) bool {
		return es2[i].Error() < es2[j].Error()
	})
	for i := 0; i < v1.Size(); i++ {
		if !reflect.DeepEqual(es1[i], es2[i]) {
			return false
		}
	}
	return true
}
