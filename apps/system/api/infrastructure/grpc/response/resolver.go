package response

import (
	"github.com/bufbuild/connect-go"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	domainValidation "github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	languageAdapter "github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/language"
	requestValidation "github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/request/validation"
	errorPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/error/v1"
	"github.com/spf13/cast"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

type Resolver interface {
	Error(ctx context.Context, err error) error
}

type resolver struct {
	mr message.Resource
	la languageAdapter.Adapter
}

func NewResolver(mr message.Resource, la languageAdapter.Adapter) Resolver {
	return &resolver{mr, la}
}

func (r *resolver) newErrorResponse(tag language.Tag, err error) ErrorResponse {
	return ErrorResponse{r.details(tag, err)}
}

func (r *resolver) isNeedCustomDetail(err error) bool {
	switch err.(type) {
	case *domainError.NotBelong:
		return true
	case *domainError.EmailAlreadyInUse:
		return true
	case *domainError.PhoneNumberAlreadyInUse:
		return true
	case *domainError.InvalidAddress:
		return true
	case *domainError.EmailNotVerified:
		return true
	default:
		return false
	}
}

func (r *resolver) basicDetail(tag language.Tag, err error) errDetail {
	switch t := err.(type) {
	case domainValidation.Error:
		return errDetail{
			&errdetails.BadRequest{
				FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{
						Field:       string(t.MessageKey()),
						Description: r.mr.ErrorMessage(string(t.MessageKey())).WithLang(tag, t.Args()...),
					},
				},
			},
		}
	case *domainError.Unauthenticated:
		return errDetail{
			&errdetails.LocalizedMessage{
				Locale:  tag.String(),
				Message: r.mr.ErrorMessage(string(domainError.UnauthenticatedMessageKey)).WithLang(tag),
			},
		}
	case *domainError.Conflicted:
		return errDetail{
			&errdetails.LocalizedMessage{
				Locale:  tag.String(),
				Message: r.mr.ErrorMessage(string(domainError.ConflictedMessageKey)).WithLang(tag),
			},
		}
	case *domainError.ConflictVersion:
		return errDetail{
			&errdetails.LocalizedMessage{
				Locale:  tag.String(),
				Message: r.mr.ErrorMessage(string(domainError.ConflictVersionMessageKey)).WithLang(tag),
			},
		}
	case *domainError.NoSuchData:
		return errDetail{
			&errdetails.LocalizedMessage{
				Locale:  tag.String(),
				Message: r.mr.ErrorMessage(string(domainError.NoSuchDataMessageKey)).WithLang(tag),
			},
		}
	default:
	}
	return errDetail{}
}

func (r *resolver) customCodeDetail(tag language.Tag, err error) errDetail {
	switch err.(type) {
	case *domainError.NotBelong:
		return errDetail{
			&errdetails.ErrorInfo{
				Reason:   r.mr.ErrorMessage(string(domainError.NotBelongMessageKey)).WithLang(tag),
				Domain:   "",
				Metadata: map[string]string{"code": errorPb.ErrorCode_ERROR_CODE_NOT_BELONG_GROUP.String()},
			},
		}
	case *domainError.EmailAlreadyInUse:
		return errDetail{
			&errdetails.ErrorInfo{
				Reason:   r.mr.ErrorMessage(string(domainError.EmailAlreadyInUseMessageKey)).WithLang(tag),
				Domain:   "",
				Metadata: map[string]string{"code": errorPb.ErrorCode_ERROR_CODE_EMAIL_ALREADY_IN_USE.String()},
			},
		}
	case *domainError.InvalidEmail:
		return errDetail{
			&errdetails.ErrorInfo{
				Reason:   r.mr.ErrorMessage(string(domainError.InvalidEmailMessageKey)).WithLang(tag),
				Domain:   "",
				Metadata: map[string]string{"code": errorPb.ErrorCode_ERROR_CODE_INVALID_EMAIL.String()},
			},
		}
	case *domainError.PhoneNumberAlreadyInUse:
		return errDetail{
			&errdetails.ErrorInfo{
				Reason:   r.mr.ErrorMessage(string(domainError.PhoneNumberAlreadyInUseMessageKey)).WithLang(tag),
				Domain:   "",
				Metadata: map[string]string{"code": cast.ToString(errorPb.ErrorCode_value[errorPb.ErrorCode_ERROR_CODE_PHONE_NUMBER_ALREADY_IN_USE.String()])},
			},
		}
	case *domainError.InvalidAddress:
		return errDetail{
			&errdetails.ErrorInfo{
				Reason:   r.mr.ErrorMessage(string(domainError.InvalidAddressMessageKey)).WithLang(tag),
				Domain:   "",
				Metadata: map[string]string{"code": cast.ToString(errorPb.ErrorCode_value[errorPb.ErrorCode_ERROR_CODE_INVALID_ADDRESS.String()])},
			},
		}
	case *domainError.EmailNotVerified:
		return errDetail{
			&errdetails.ErrorInfo{
				Reason:   r.mr.ErrorMessage(string(domainError.EmailNotVerifiedMessageKey)).WithLang(tag),
				Domain:   "",
				Metadata: map[string]string{"code": cast.ToString(errorPb.ErrorCode_value[errorPb.ErrorCode_ERROR_CODE_EMAIL_NOT_VERIFIED.String()])},
			},
		}
	default:
		return errDetail{}
	}
}

func (r *resolver) handleDetail(tag language.Tag, de error) errDetail {
	if r.isNeedCustomDetail(de) {
		return r.customCodeDetail(tag, de)
	} else {
		return r.basicDetail(tag, de)
	}
}

func (r *resolver) details(tag language.Tag, err error) []errDetail {
	if des, ok := err.(domainValidation.Errors); ok {
		result := make([]errDetail, 0, des.Size())
		for _, de := range des.AsSlice() {
			result = append(result, r.handleDetail(tag, de))
		}
		return result
	}
	return []errDetail{r.handleDetail(tag, err)}
}

func (r *resolver) statusCode(err error) codes.Code {
	switch err.(type) {
	case *domainError.Unauthenticated:
		return codes.Unauthenticated
	case *domainError.NotBelong:
		return codes.PermissionDenied
	case requestValidation.FieldError, domainValidation.Errors:
		return codes.InvalidArgument
	case *domainError.EmailAlreadyInUse:
		return codes.InvalidArgument
	case *domainError.PhoneNumberAlreadyInUse:
		return codes.InvalidArgument
	case *domainError.InvalidAddress:
		return codes.InvalidArgument
	case *domainError.EmailNotVerified:
		return codes.InvalidArgument
	case *domainError.Conflicted:
		return codes.AlreadyExists
	case *domainError.ConflictVersion:
		return codes.AlreadyExists
	case *domainError.NoSuchData:
		return codes.NotFound
	}
	return codes.Internal
}

func (r *resolver) statusCodeForConnect(err error) connect.Code {
	switch err.(type) {
	case *domainError.Unauthenticated:
		return connect.CodeUnauthenticated
	case *domainError.NotBelong:
		return connect.CodePermissionDenied
	case requestValidation.FieldError, domainValidation.Errors:
		return connect.CodeInvalidArgument
	case *domainError.EmailAlreadyInUse:
		return connect.CodeInvalidArgument
	case *domainError.PhoneNumberAlreadyInUse:
		return connect.CodeInvalidArgument
	case *domainError.InvalidAddress:
		return connect.CodeInvalidArgument
	case *domainError.EmailNotVerified:
		return connect.CodeInvalidArgument
	case *domainError.Conflicted:
		return connect.CodeAlreadyExists
	case *domainError.ConflictVersion:
		return connect.CodeInvalidArgument
	case *domainError.NoSuchData:
		return connect.CodeNotFound
	}
	return connect.CodeInternal
}

func (r *resolver) mainErrorMessage(tag language.Tag, err error) string {
	switch err.(type) {
	case *domainError.Unauthenticated:
		return r.mr.ErrorMessage(string(domainError.UnauthenticatedMessageKey)).WithLang(tag)
	case *domainError.NotBelong:
		return r.mr.ErrorMessage(string(domainError.NotBelongMessageKey)).WithLang(tag)
	case requestValidation.FieldError, domainValidation.Errors:
		return "invalid arguments"
	case *domainError.PhoneNumberAlreadyInUse:
		return r.mr.ErrorMessage(string(domainError.PhoneNumberAlreadyInUseMessageKey)).WithLang(tag)
	case *domainError.EmailAlreadyInUse:
		return r.mr.ErrorMessage(string(domainError.EmailAlreadyInUseMessageKey)).WithLang(tag)
	case *domainError.InvalidEmail:
		return r.mr.ErrorMessage(string(domainError.InvalidEmailMessageKey)).WithLang(tag)
	case *domainError.InvalidAddress:
		return r.mr.ErrorMessage(string(domainError.InvalidAddressMessageKey)).WithLang(tag)
	case *domainError.EmailNotVerified:
		return r.mr.ErrorMessage(string(domainError.EmailNotVerifiedMessageKey)).WithLang(tag)
	case *domainError.Conflicted:
		return r.mr.ErrorMessage(string(domainError.ConflictedMessageKey)).WithLang(tag)
	case *domainError.ConflictVersion:
		return r.mr.ErrorMessage(string(domainError.ConflictVersionMessageKey)).WithLang(tag)
	case *domainError.NoSuchData:
		return r.mr.ErrorMessage(string(domainError.NoSuchDataMessageKey)).WithLang(tag)
	}
	return "internal server error has occurred"
}

func (r *resolver) getLanguage(ctx context.Context) language.Tag {
	return r.la.Adapt(ctx)
}

func (r *resolver) Error(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	langTag := r.getLanguage(ctx)
	ner := r.newErrorResponse(langTag, err)
	cErr := connect.NewError(
		r.statusCodeForConnect(err),
		err,
	)
	for _, ne := range ner.Errors {
		if detail, detailErr := connect.NewErrorDetail(ne.Message); detailErr == nil {
			cErr.AddDetail(detail)
		}
	}
	return cErr
}
