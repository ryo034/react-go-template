package middleware

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
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
	co          shared.ContextOperator
}

func (lm *langMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(lm.co.SetLang(r.Context(), Language(r, lm.defaultLang))))
	})
}

func NewLangMiddleware(defaultLang language.Tag, co shared.ContextOperator) LangMiddleware {
	return &langMiddleware{defaultLang, co}
}
