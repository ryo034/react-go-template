package middleware

import (
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
