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
	var noSuchData *domainErr.NoSuchData
	var unauthenticatedErr *domainErr.Unauthenticated
	var conflictErr *domainErr.Conflicted

	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var invalidInviteToken *invitation.InvalidInviteToken
	var inviteTokenExpired *invitation.ExpiredInviteToken

	switch {
	case errors.As(err, &t):
		msg = r.mr.TitleMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &noSuchData):
		msg = r.mr.TitleMessage(string(domainErr.NoSuchDataMessageKey)).WithLang(tag)
	case errors.As(err, &unauthenticatedErr):
		msg = r.mr.TitleMessage(string(domainErr.UnauthenticatedMessageKey)).WithLang(tag)
	case errors.As(err, &conflictErr):
		msg = r.mr.TitleMessage(string(domainErr.ConflictedMessageKey)).WithLang(tag)
	case errors.As(err, &inviteTokenExpired):
		msg = r.mr.TitleMessage(string(invitation.ExpiredInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &invalidInviteToken):
		msg = r.mr.TitleMessage(string(invitation.InvalidInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyExpiredInvitation):
		msg = r.mr.TitleMessage(string(invitation.AlreadyExpiredInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyRevokedInvitation):
		msg = r.mr.TitleMessage(string(invitation.AlreadyRevokedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyVerifiedInvitation):
		msg = r.mr.TitleMessage(string(invitation.AlreadyVerifiedInvitationMessageKey)).WithLang(tag)
	}
	return msg
}

func (r *resolver) errDetail(tag language.Tag, err error, msgArgs ...interface{}) string {
	msg := r.mr.DetailMessage(string(domainErr.InternalServerErrorMessageKey)).WithLang(tag)
	var t domainValidation.Error
	var invalidInviteToken *invitation.InvalidInviteToken
	var inviteTokenExpired *invitation.ExpiredInvitation
	var unauthenticatedErr *domainErr.Unauthenticated
	var noSuchData *domainErr.NoSuchData
	var conflictErr *domainErr.Conflicted

	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation

	switch {
	case errors.As(err, &t):
		msg = r.mr.DetailMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &noSuchData):
		msg = r.mr.DetailMessage(string(domainErr.NoSuchDataMessageKey)).WithLang(tag)
	case errors.As(err, &unauthenticatedErr):
		msg = r.mr.DetailMessage(string(domainErr.UnauthenticatedMessageKey)).WithLang(tag)
	case errors.As(err, &conflictErr):
		msg = r.mr.DetailMessage(string(domainErr.ConflictedMessageKey)).WithLang(tag)
	case errors.As(err, &inviteTokenExpired):
		msg = r.mr.DetailMessage(string(invitation.ExpiredInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &invalidInviteToken):
		msg = r.mr.DetailMessage(string(invitation.InvalidInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyExpiredInvitation):
		msg = r.mr.DetailMessage(string(invitation.AlreadyExpiredInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyRevokedInvitation):
		msg = r.mr.DetailMessage(string(invitation.AlreadyRevokedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyVerifiedInvitation):
		msg = r.mr.DetailMessage(string(invitation.AlreadyVerifiedInvitationMessageKey)).WithLang(tag)
	}
	return msg
}

func (r *resolver) errType(tag language.Tag, err error, msgArgs ...interface{}) string {
	msg := r.mr.TypeMessage(string(domainErr.InternalServerErrorMessageKey)).WithLang(tag)
	var t domainValidation.Error
	var noSuchData *domainErr.NoSuchData
	var unauthenticatedErr *domainErr.Unauthenticated
	var conflictErr *domainErr.Conflicted

	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var invalidInviteToken *invitation.InvalidInviteToken
	var inviteTokenExpired *invitation.ExpiredInviteToken

	switch {
	case errors.As(err, &t):
		msg = r.mr.TypeMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &noSuchData):
		msg = r.mr.TypeMessage(string(domainErr.NoSuchDataMessageKey)).WithLang(tag)
	case errors.As(err, &unauthenticatedErr):
		msg = r.mr.TypeMessage(string(domainErr.UnauthenticatedMessageKey)).WithLang(tag)
	case errors.As(err, &conflictErr):
		msg = r.mr.TypeMessage(string(domainErr.ConflictedMessageKey)).WithLang(tag)
	case errors.As(err, &inviteTokenExpired):
		msg = r.mr.TypeMessage(string(invitation.ExpiredInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &invalidInviteToken):
		msg = r.mr.TypeMessage(string(invitation.InvalidInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyExpiredInvitation):
		msg = r.mr.TypeMessage(string(invitation.AlreadyExpiredInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyRevokedInvitation):
		msg = r.mr.TypeMessage(string(invitation.AlreadyRevokedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyVerifiedInvitation):
		msg = r.mr.TypeMessage(string(invitation.AlreadyVerifiedInvitationMessageKey)).WithLang(tag)
	}
	return msg
}

func (r *resolver) errCode(err error) string {
	code := "500-000"
	var t domainValidation.Error
	var noSuchData *domainErr.NoSuchData
	var unauthenticatedErr *domainErr.Unauthenticated
	var conflictErr *domainErr.Conflicted

	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var invalidInviteToken *invitation.InvalidInviteToken
	var inviteTokenExpired *invitation.ExpiredInviteToken

	switch {
	case errors.As(err, &t):
		code = "400-000"
	case errors.As(err, &unauthenticatedErr):
		code = "401-000"
	case errors.As(err, &noSuchData):
		code = "404-000"
	case errors.As(err, &conflictErr):
		code = "409-000"
	case errors.As(err, &inviteTokenExpired):
		code = "400-001"
	case errors.As(err, &invalidInviteToken):
		code = "400-000"
	case errors.As(err, &alreadyExpiredInvitation):
		code = "410-000"
	case errors.As(err, &alreadyRevokedInvitation):
		code = "410-001"
	case errors.As(err, &alreadyVerifiedInvitation):
		code = "409-000"
	}
	return code
}

func (r *resolver) Error(c context.Context, err error) interface{} {
	er := r.newErrorResponse(r.getLanguage(c), err)

	var t domainValidation.Error
	var noSuchData *domainErr.NoSuchData
	var unauthenticatedErr *domainErr.Unauthenticated
	var conflictErr *domainErr.Conflicted

	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var invalidInviteToken *invitation.InvalidInviteToken
	var inviteTokenExpired *invitation.ExpiredInviteToken

	switch {
	case errors.As(err, &t), errors.As(err, &inviteTokenExpired), errors.As(err, &invalidInviteToken):
		return &openapi.BadRequestError{
			Type:   openapi.OptString{Value: er.Type, Set: true},
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	case errors.As(err, &noSuchData):
		return &openapi.NotFoundError{
			Type:   openapi.OptString{Value: er.Type, Set: true},
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	case errors.As(err, &unauthenticatedErr):
		return &openapi.UnauthorizedError{
			Type:   openapi.OptString{Value: er.Type, Set: true},
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	case errors.As(err, &conflictErr), errors.As(err, &alreadyVerifiedInvitation):
		return &openapi.ConflictError{
			Type:   openapi.OptString{Value: er.Type, Set: true},
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	case errors.As(err, &alreadyExpiredInvitation), errors.As(err, &alreadyRevokedInvitation):
		return &openapi.GoneError{
			Type:   openapi.OptString{Value: er.Type, Set: true},
			Title:  openapi.OptString{Value: er.Title, Set: true},
			Detail: openapi.OptString{Value: er.Detail, Set: true},
			Code:   openapi.OptString{Value: er.Code, Set: true},
		}
	}

	return &openapi.InternalServerError{
		Type:   openapi.OptString{Value: er.Type, Set: true},
		Title:  openapi.OptString{Value: er.Title, Set: true},
		Detail: openapi.OptString{Value: er.Detail, Set: true},
		Code:   openapi.OptString{Value: er.Code, Set: true},
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
	var noSuchData *domainErr.NoSuchData
	var unauthenticatedErr *domainErr.Unauthenticated
	var conflictErr *domainErr.Conflicted

	var alreadyExpiredInvitation *invitation.AlreadyExpiredInvitation
	var alreadyRevokedInvitation *invitation.AlreadyRevokedInvitation
	var alreadyVerifiedInvitation *invitation.AlreadyVerifiedInvitation
	var invalidInviteToken *invitation.InvalidInviteToken
	var inviteTokenExpired *invitation.ExpiredInviteToken

	switch {
	case errors.As(err, &t):
		msg = r.mr.ErrorMessage(string(t.MessageKey())).WithLang(tag, t.Args()...)
	case errors.As(err, &noSuchData):
		msg = r.mr.ErrorMessage(string(domainErr.NoSuchDataMessageKey)).WithLang(tag)
	case errors.As(err, &unauthenticatedErr):
		msg = r.mr.ErrorMessage(string(domainErr.UnauthenticatedMessageKey)).WithLang(tag)
	case errors.As(err, &conflictErr):
		msg = r.mr.ErrorMessage(string(domainErr.ConflictedMessageKey)).WithLang(tag)
	case errors.As(err, &inviteTokenExpired):
		msg = r.mr.ErrorMessage(string(invitation.ExpiredInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &invalidInviteToken):
		msg = r.mr.ErrorMessage(string(invitation.InvalidInviteTokenMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyVerifiedInvitation):
		msg = r.mr.ErrorMessage(string(invitation.AlreadyVerifiedInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyExpiredInvitation):
		msg = r.mr.ErrorMessage(string(invitation.AlreadyExpiredInvitationMessageKey)).WithLang(tag)
	case errors.As(err, &alreadyRevokedInvitation):
		msg = r.mr.ErrorMessage(string(invitation.AlreadyRevokedInvitationMessageKey)).WithLang(tag)
	}
	return errDetail{msg}
}
