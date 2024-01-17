package middleware

import (
	"context"
	"github.com/ogen-go/ogen/middleware"
	middl "github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
	"golang.org/x/text/language"
)

func Language(defaultLang language.Tag) middleware.Middleware {
	return func(req middleware.Request, next func(req middleware.Request) (middleware.Response, error)) (middleware.Response, error) {
		lang := middl.Language(req.Raw, defaultLang)
		req.SetContext(context.WithValue(req.Context, "lang", lang))
		return next(req)
	}
}
