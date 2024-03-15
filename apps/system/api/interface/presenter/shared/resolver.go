package shared

import (
	"context"
	"errors"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	domainValidation "github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
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
	msg := r.mr.TitleMessage(string(domainErr.InternalServerErrorMessageKey)).WithLang(tag)

	var t domainValidation.Error
	var badRequest *domainErr.BadRequest
	var unauthenticated *domainErr.Unauthenticated
	var forbidden *domainErr.Forbidden
	var noSuchData *domainErr.NoSuchData
	var conflict *domainErr.Conflicted
	var gone *domainErr.Gone

	var invalidInviteToken *invitation.InvalidInvitationToken
	var expiredInvitationToken *invitation.ExpiredInvitationToken
	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var alreadyAcceptedInvitation *invitation.AlreadyAcceptedInvitation

	switch {
	case errors.As(err, &badRequest):
		msg = r.mr.TitleMessage(string(domainErr.BadRequestMessageKey)).WithLang(tag)
	case errors.As(err, &t):
		msg = r.mr.TitleMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &unauthenticated):
		msg = r.mr.TitleMessage(string(domainErr.UnauthenticatedMessageKey)).WithLang(tag)
	case errors.As(err, &forbidden):
		msg = r.mr.TitleMessage(string(domainErr.ForbiddenMessageKey)).WithLang(tag)
	case errors.As(err, &noSuchData):
		msg = r.mr.TitleMessage(string(domainErr.NoSuchDataMessageKey)).WithLang(tag)
	case errors.As(err, &conflict):
		msg = r.mr.TitleMessage(string(domainErr.ConflictedMessageKey)).WithLang(tag)
	case errors.As(err, &gone):
		msg = r.mr.TitleMessage(string(domainErr.GoneMessageKey)).WithLang(tag)
	case errors.As(err, &invalidInviteToken):
		msg = r.mr.TitleMessage(string(invitation.InvalidInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &expiredInvitationToken):
		msg = r.mr.TitleMessage(string(invitation.ExpiredInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyExpiredInvitation):
		msg = r.mr.TitleMessage(string(invitation.AlreadyExpiredInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyRevokedInvitation):
		msg = r.mr.TitleMessage(string(invitation.AlreadyRevokedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyVerifiedInvitation):
		msg = r.mr.TitleMessage(string(invitation.AlreadyVerifiedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyAcceptedInvitation):
		msg = r.mr.TitleMessage(string(invitation.AlreadyAcceptedInvitationMessageKey)).WithLang(tag)
	}
	return msg
}

func (r *resolver) errDetail(tag language.Tag, err error, msgArgs ...interface{}) string {
	msg := r.mr.DetailMessage(string(domainErr.InternalServerErrorMessageKey)).WithLang(tag)
	var t domainValidation.Error
	var badRequest *domainErr.BadRequest
	var unauthenticated *domainErr.Unauthenticated
	var forbidden *domainErr.Forbidden
	var noSuchData *domainErr.NoSuchData
	var conflict *domainErr.Conflicted
	var gone *domainErr.Gone

	var invalidInviteToken *invitation.InvalidInvitationToken
	var expiredInvitationToken *invitation.ExpiredInvitationToken
	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var alreadyAcceptedInvitation *invitation.AlreadyAcceptedInvitation

	switch {
	case errors.As(err, &badRequest):
		msg = r.mr.DetailMessage(string(domainErr.BadRequestMessageKey)).WithLang(tag)
	case errors.As(err, &t):
		msg = r.mr.DetailMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &unauthenticated):
		msg = r.mr.DetailMessage(string(domainErr.UnauthenticatedMessageKey)).WithLang(tag)
	case errors.As(err, &forbidden):
		msg = r.mr.DetailMessage(string(domainErr.ForbiddenMessageKey)).WithLang(tag)
	case errors.As(err, &noSuchData):
		msg = r.mr.DetailMessage(string(domainErr.NoSuchDataMessageKey)).WithLang(tag)
	case errors.As(err, &conflict):
		msg = r.mr.DetailMessage(string(domainErr.ConflictedMessageKey)).WithLang(tag)
	case errors.As(err, &gone):
		msg = r.mr.DetailMessage(string(domainErr.GoneMessageKey)).WithLang(tag)
	case errors.As(err, &invalidInviteToken):
		msg = r.mr.DetailMessage(string(invitation.InvalidInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &expiredInvitationToken):
		msg = r.mr.DetailMessage(string(invitation.ExpiredInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyExpiredInvitation):
		msg = r.mr.DetailMessage(string(invitation.AlreadyExpiredInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyRevokedInvitation):
		msg = r.mr.DetailMessage(string(invitation.AlreadyRevokedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyVerifiedInvitation):
		msg = r.mr.DetailMessage(string(invitation.AlreadyVerifiedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyAcceptedInvitation):
		msg = r.mr.DetailMessage(string(invitation.AlreadyAcceptedInvitationMessageKey)).WithLang(tag)
	}
	return msg
}

func (r *resolver) errCode(err error) string {
	code := "500-000"
	var t domainValidation.Error
	var badRequest *domainErr.BadRequest
	var unauthenticated *domainErr.Unauthenticated
	var forbidden *domainErr.Forbidden
	var noSuchData *domainErr.NoSuchData
	var conflict *domainErr.Conflicted
	var gone *domainErr.Gone

	var invalidInviteToken *invitation.InvalidInvitationToken
	var expiredInvitationToken *invitation.ExpiredInvitationToken
	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var alreadyAcceptedInvitation *invitation.AlreadyAcceptedInvitation

	switch {
	case errors.As(err, &badRequest):
		code = "400-000"
	case errors.As(err, &t):
		code = "400-000"
	case errors.As(err, &unauthenticated):
		code = "401-000"
	case errors.As(err, &forbidden):
		code = "403-000"
	case errors.As(err, &noSuchData):
		code = "404-000"
	case errors.As(err, &conflict):
		code = "409-000"
	case errors.As(err, &gone):
		code = "410-000"
	case errors.As(err, &expiredInvitationToken):
		code = "400-000"
	case errors.As(err, &invalidInviteToken):
		code = "400-000"
	case errors.As(err, &alreadyVerifiedInvitation):
		code = "409-000"
	case errors.As(err, &alreadyExpiredInvitation):
		code = "410-001"
	case errors.As(err, &alreadyRevokedInvitation):
		code = "410-002"
	case errors.As(err, &alreadyAcceptedInvitation):
		code = "410-003"
	}
	return code
}

func (r *resolver) Error(c context.Context, err error) interface{} {
	er := r.newErrorResponse(r.getLanguage(c), err)

	var t domainValidation.Error
	var badRequest *domainErr.BadRequest
	var unauthenticated *domainErr.Unauthenticated
	var forbidden *domainErr.Forbidden
	var conflict *domainErr.Conflicted
	var gone *domainErr.Gone

	var invalidInviteToken *invitation.InvalidInvitationToken
	var expiredInvitationToken *invitation.ExpiredInvitationToken
	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var alreadyAcceptedInvitation *invitation.AlreadyAcceptedInvitation

	switch {
	case errors.As(err, &badRequest), errors.As(err, &t), errors.As(err, &invalidInviteToken), errors.As(err, &expiredInvitationToken):
		return &openapi.BadRequestError{
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	case errors.As(err, &unauthenticated):
		return &openapi.UnauthorizedError{
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	case errors.As(err, &forbidden):
		return &openapi.ForbiddenError{
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	case errors.As(err, &conflict), errors.As(err, &alreadyVerifiedInvitation):
		return &openapi.ConflictError{
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	case errors.As(err, &gone), errors.As(err, &alreadyExpiredInvitation), errors.As(err, &alreadyRevokedInvitation), errors.As(err, &alreadyAcceptedInvitation):
		return &openapi.GoneError{
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	}
	return &openapi.InternalServerError{
		Title:  openapi.OptString{Value: er.Title, Set: true},
		Detail: openapi.OptString{Value: er.Detail, Set: true},
		Code:   openapi.OptString{Value: er.Code, Set: true},
	}
}

func (r *resolver) getLanguage(c context.Context) language.Tag {
	return r.la.Adapt(c)
}

func (r *resolver) newErrorResponse(tag language.Tag, err error) ErrorResponse {
	//一旦ログは全て英語で出す
	forceTag := language.English
	return ErrorResponse{
		r.errTitle(forceTag, err),
		r.errDetail(forceTag, err),
		r.details(forceTag, err),
		r.errCode(err),
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
	msg := r.mr.ErrorMessage(string(domainErr.InternalServerErrorMessageKey)).WithLang(tag)
	var t domainValidation.Error
	var badRequest *domainErr.BadRequest
	var unauthenticated *domainErr.Unauthenticated
	var forbidden *domainErr.Forbidden
	var noSuchData *domainErr.NoSuchData
	var conflict *domainErr.Conflicted
	var gone *domainErr.Gone

	var invalidInviteToken *invitation.InvalidInvitationToken
	var expiredInvitationToken *invitation.ExpiredInvitationToken
	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var alreadyAcceptedInvitation *invitation.AlreadyAcceptedInvitation

	switch {
	case errors.As(err, &badRequest):
		msg = r.mr.ErrorMessage(string(domainErr.BadRequestMessageKey)).WithLang(tag)
	case errors.As(err, &t):
		msg = r.mr.ErrorMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &unauthenticated):
		msg = r.mr.ErrorMessage(string(domainErr.UnauthenticatedMessageKey)).WithLang(tag)
	case errors.As(err, &forbidden):
		msg = r.mr.ErrorMessage(string(domainErr.ForbiddenMessageKey)).WithLang(tag)
	case errors.As(err, &noSuchData):
		msg = r.mr.ErrorMessage(string(domainErr.NoSuchDataMessageKey)).WithLang(tag)
	case errors.As(err, &conflict):
		msg = r.mr.ErrorMessage(string(domainErr.ConflictedMessageKey)).WithLang(tag)
	case errors.As(err, &gone):
		msg = r.mr.ErrorMessage(string(domainErr.GoneMessageKey)).WithLang(tag)
	case errors.As(err, &invalidInviteToken):
		msg = r.mr.ErrorMessage(string(invitation.InvalidInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &expiredInvitationToken):
		msg = r.mr.ErrorMessage(string(invitation.ExpiredInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyVerifiedInvitation):
		msg = r.mr.ErrorMessage(string(invitation.AlreadyVerifiedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyExpiredInvitation):
		msg = r.mr.ErrorMessage(string(invitation.AlreadyExpiredInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyRevokedInvitation):
		msg = r.mr.ErrorMessage(string(invitation.AlreadyRevokedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyAcceptedInvitation):
		msg = r.mr.ErrorMessage(string(invitation.AlreadyAcceptedInvitationMessageKey)).WithLang(tag)
	}
	return errDetail{msg}
}
