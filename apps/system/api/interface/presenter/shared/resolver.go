package shared

import (
	"context"
	"errors"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	domainValidation "github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/request/validation"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"golang.org/x/text/language"
)

type Resolver interface {
	SuccessMessage(c context.Context, msgKey string, msgArgs ...interface{}) string
	Error(c context.Context, err error) interface{}
}

type resolver struct {
	mr message.Resource
	//th translation.TranslatorHolder
	la LanguageAdapter
}

func NewResolver(
	mr message.Resource,
	//th translation.TranslatorHolder,
	la LanguageAdapter,
) Resolver {
	return &resolver{mr, la}
}

func (r *resolver) SuccessMessage(c context.Context, msgKey string, msgArgs ...interface{}) string {
	return r.mr.SuccessMessage(msgKey).WithLang(r.getLanguage(c), msgArgs...)
}

func (r *resolver) errTitle(tag language.Tag, err error, msgArgs ...interface{}) string {
	msg := ""
	var t domainValidation.Error
	var conflictErr *domainError.Conflicted
	var noSuchData *domainError.NoSuchData
	switch {
	case errors.As(err, &t):
		msg = r.mr.TitleMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &conflictErr):
		msg = r.mr.TitleMessage("conflicted").WithLang(tag)
	case errors.As(err, &noSuchData):
		msg = r.mr.TitleMessage("noSuchData").WithLang(tag)
	}
	return msg
}

func (r *resolver) errDetail(tag language.Tag, err error, msgArgs ...interface{}) string {
	msg := ""
	var t domainValidation.Error
	var conflictErr *domainError.Conflicted
	var noSuchData *domainError.NoSuchData
	switch {
	case errors.As(err, &t):
		msg = r.mr.DetailMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &conflictErr):
		msg = r.mr.DetailMessage("conflicted").WithLang(tag)
	case errors.As(err, &noSuchData):
		msg = r.mr.DetailMessage("noSuchData").WithLang(tag)
	}
	return msg
}

func (r *resolver) errType(tag language.Tag, err error, msgArgs ...interface{}) string {
	msg := ""
	var t domainValidation.Error
	var conflictErr *domainError.Conflicted
	var noSuchData *domainError.NoSuchData
	switch {
	case errors.As(err, &t):
		msg = r.mr.TypeMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &conflictErr):
		msg = r.mr.TypeMessage("conflicted").WithLang(tag)
	case errors.As(err, &noSuchData):
		msg = r.mr.TypeMessage("noSuchData").WithLang(tag)
	}
	return msg
}

func (r *resolver) Error(c context.Context, err error) interface{} {
	er := r.newErrorResponse(r.getLanguage(c), err)

	var t domainValidation.Error
	var conflictErr *domainError.Conflicted
	var noSuchData *domainError.NoSuchData

	switch {
	case errors.As(err, &t):
		return &openapi.BadRequestError{
			Type:   openapi.OptString{Value: er.Type, Set: true},
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
		}
	case errors.As(err, &conflictErr):
	case errors.As(err, &noSuchData):
	}
	return &openapi.InternalServerError{
		Type:   openapi.OptString{Value: er.Type, Set: true},
		Title:  openapi.OptString{Value: er.Title, Set: true},
		Detail: openapi.OptString{Value: er.Detail, Set: true},
	}
}

func (r *resolver) getLanguage(c context.Context) language.Tag {
	return r.la.Adapt(c)
}

func (r *resolver) newErrorResponse(tag language.Tag, err error) ErrorResponse {
	return ErrorResponse{
		r.errTitle(tag, err),
		r.errDetail(tag, err),
		r.errType(tag, err),
		r.details(tag, err),
	}
}

func (r *resolver) details(tag language.Tag, err error) []errDetail {
	var fes validation.FieldErrors
	if errors.As(err, &fes) {
		result := make([]errDetail, 0, len(fes))
		for _, fe := range fes {
			result = append(result, r.detail(tag, fe))
		}
		return result
	}
	var des domainValidation.Errors
	if errors.As(err, &des) {
		result := make([]errDetail, 0, des.Size())
		for _, de := range des.AsSlice() {
			result = append(result, r.detail(tag, de))
		}
		return result
	}
	return []errDetail{r.detail(tag, err)}
}

func (r *resolver) detail(tag language.Tag, err error) errDetail {
	msg := ""
	var t domainValidation.Error
	var conflictErr *domainError.Conflicted
	var noSuchData *domainError.NoSuchData
	switch {
	case errors.As(err, &t):
		msg = r.mr.ErrorMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &conflictErr):
		msg = r.mr.ErrorMessage("conflicted").WithLang(tag)
	case errors.As(err, &noSuchData):
		msg = r.mr.ErrorMessage("noSuchData").WithLang(tag)
	}
	return errDetail{msg}
}
