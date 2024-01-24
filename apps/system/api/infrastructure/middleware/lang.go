package middleware

import (
	"context"
	"golang.org/x/text/language"
	"net/http"
)

var supportedLanguages = []language.Tag{language.English, language.Japanese}

func Language(req *http.Request, defaultLang language.Tag) language.Tag {
	lang := defaultLang
	al := req.Header.Get("Accept-Language")
	if al != "" {
		if tags, _, err := language.ParseAcceptLanguage(al); err == nil {
			switch tags[0] {
			case supportedLanguages[0], supportedLanguages[1]:
				lang = tags[0]
			}
		}
	}
	return lang
}

type LangMiddleware interface {
	Handler(h http.Handler) http.Handler
}

type langMiddleware struct {
	defaultLang language.Tag
}

func (lm *langMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := Language(r, lm.defaultLang)
		ctx := context.WithValue(r.Context(), "lang", lang)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewLangMiddleware(defaultLang language.Tag) LangMiddleware {
	return &langMiddleware{defaultLang}
}
