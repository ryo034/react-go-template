package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type FieldErrors []FieldError

func (fes FieldErrors) Error() string {
	errors := make([]string, 0, len(fes))
	for _, fe := range fes {
		errors = append(errors, fe.Error())
	}
	return fmt.Sprint(errors)
}

func wrap(ves validator.ValidationErrors) FieldErrors {
	result := make([]FieldError, 0, len(ves))
	for _, fe := range ves {
		result = append(result, &fieldError{fe})
	}
	return result
}

type FieldError interface {
	validator.FieldError
	error
}

type fieldError struct {
	validator.FieldError
}

func (fe *fieldError) Error() string {
	return fmt.Sprint(fe.FieldError.(error).Error())
}
