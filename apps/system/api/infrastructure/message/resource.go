package message

import (
	"fmt"

	"golang.org/x/text/language"
)

type Resource interface {
	ErrorMessage(key string) Message
	FieldName(key string) Message
	FieldNameFromTag(tag FieldNameTag) Message
}

type Message interface {
	Default(args ...interface{}) string
	WithLang(tag language.Tag, args ...interface{}) string
}

type message struct {
	defaultLang language.Tag
	mp          map[language.Tag]string
}

func (m message) Default(args ...interface{}) string {
	return fmt.Sprintf(m.mp[m.defaultLang], args...)
}

func (m message) WithLang(tag language.Tag, args ...interface{}) string {
	if j, ok := m.mp[tag]; ok {
		return fmt.Sprintf(j, args...)
	}
	return m.Default(args...)
}

type resource struct {
	errorMessages map[string]Message
	fieldNames    map[string]Message
}

func NewResource(defaultLang language.Tag) Resource {
	domainErrorMessages := DomainErrorMessages
	errorMessages := make(map[string]map[language.Tag]string, len(commonErrorMessages)+len(domainErrorMessages))
	for k, v := range commonErrorMessages {
		errorMessages[k] = v
	}
	for k, v := range domainErrorMessages {
		errorMessages[string(k)] = v
	}
	toMassages := func(m map[string]map[language.Tag]string) map[string]Message {
		result := make(map[string]Message, len(m))
		for k, v := range m {
			result[k] = message{defaultLang, v}
		}
		return result
	}
	return &resource{toMassages(errorMessages), toMassages(filedNames)}
}

func (m *resource) ErrorMessage(key string) Message {
	return m.errorMessages[key]
}

func (m *resource) FieldName(key string) Message {
	return m.fieldNames[key]
}

func (m *resource) FieldNameFromTag(tag FieldNameTag) Message {
	return m.fieldNames[string(tag)]
}
