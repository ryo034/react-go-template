package message

import (
	"fmt"

	"golang.org/x/text/language"
)

type Resource interface {
	TypeMessage(key string) Message
	TitleMessage(key string) Message
	DetailMessage(key string) Message
	ErrorMessage(key string) Message
	SuccessMessage(key string) Message
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
	errorMessages    map[string]Message
	successMessages  map[string]Message
	titleMessages    map[string]Message
	detailMessages   map[string]Message
	typeMessages     map[string]Message
	instanceMessages map[string]Message
	fieldNames       map[string]Message
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
	return &resource{
		toMassages(errorMessages),
		toMassages(successMessages),
		toMassages(titleMessages),
		toMassages(detailMessages),
		toMassages(typeMessages),
		toMassages(instanceMessages),
		toMassages(filedNames),
	}
}

func (m *resource) TypeMessage(key string) Message {
	return m.typeMessages[key]
}

func (m *resource) InstanceMessage(key string) Message {
	return m.instanceMessages[key]
}

func (m *resource) TitleMessage(key string) Message {
	return m.titleMessages[key]
}

func (m *resource) DetailMessage(key string) Message {
	return m.detailMessages[key]
}

func (m *resource) ErrorMessage(key string) Message {
	return m.errorMessages[key]
}

func (m *resource) SuccessMessage(key string) Message {
	return m.successMessages[key]
}

func (m *resource) FieldName(key string) Message {
	return m.fieldNames[key]
}

func (m *resource) FieldNameFromTag(tag FieldNameTag) Message {
	return m.fieldNames[string(tag)]
}
