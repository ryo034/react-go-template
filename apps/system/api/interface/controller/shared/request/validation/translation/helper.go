package translation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"golang.org/x/text/language"
)

type helper struct {
	ut.Translator
	message.Resource
	lTag                 language.Tag
	registerMessagesHook func() error
}

func newHelper(trans ut.Translator, mr message.Resource) (*helper, error) {
	lTag, err := language.Parse(trans.Locale())
	rHook := func() error {
		return nil
	}
	if err != nil {
		return nil, err
	}
	return &helper{trans, mr, lTag, rHook}, nil
}

func (th *helper) setRegisterMessagesHook(rHook func() error) {
	th.registerMessagesHook = rHook
}

func (th *helper) template(tag string) string {
	return th.ErrorMessage(tag).WithLang(th.lTag)
}

func (th *helper) fieldName(tag string) string {
	return th.FieldName(tag).WithLang(th.lTag)
}

func (th *helper) addAll(tags ...string) {
	for _, tag := range tags {
		_ = th.Add(tag, th.template(tag), false) // override false will never cause an error
	}
}

func (th *helper) addCardinalAll(rule locales.PluralRule, tags ...string) error {
	for _, tag := range tags {
		msgKey := tag + "-" + strings.ToLower(rule.String())
		if err := th.AddCardinal(tag, th.template(msgKey), rule, false); err != nil {
			return err
		}
	}
	return nil
}

func (th *helper) translate(fe validator.FieldError) string {
	t, err := th.T(fe.Tag(), th.fieldName(fe.Field()))
	if err != nil {
		return errTranslate(err, fe)
	}
	return t
}

func (th *helper) translateWithParam(fe validator.FieldError) string {
	return th.customTranslateWithParam(fe, fe.Tag(), fe.Param())
}

func (th *helper) translateWithFieldNameParam(fe validator.FieldError) string {
	return th.customTranslateWithParam(fe, fe.Tag(), th.fieldName(fe.Param()))
}

func (th *helper) translateCompareError(fe validator.FieldError) string {
	switch kindFrom(fe) {
	case reflect.String:
		return th.translateWithC(fe, fe.Tag()+"-string", "character-count")
	case reflect.Slice, reflect.Map, reflect.Array:
		return th.translateWithC(fe, fe.Tag()+"-items", "item-count")
	default:
		return th.translateNum(fe)
	}
}

func (th *helper) translateWithC(fe validator.FieldError, key string, countKey string) string {
	f64, digits, err := th.toNumber(fe.Param())
	if err != nil {
		return errTranslate(err, fe)
	}
	c, err := th.C(countKey, f64, digits, th.FmtNumber(f64, digits))
	if err != nil {
		return errTranslate(err, fe)
	}
	return th.customTranslateWithParam(fe, key, c)
}

func (th *helper) translateNum(fe validator.FieldError) string {
	f64, digits, err := th.toNumber(fe.Param())
	if err != nil {
		return errTranslate(err, fe)
	}
	return th.customTranslateWithParam(fe, fe.Tag()+"-number", th.FmtNumber(f64, digits))
}

func (th *helper) toNumber(param string) (float64, uint64, error) {
	digitsFrom := func(param string) uint64 {
		if idx := strings.Index(param, "."); idx != -1 {
			return uint64(len(param[idx+1:]))
		}
		return 0
	}
	f64, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, 0, err
	}
	return f64, digitsFrom(param), nil
}

func (th *helper) customTranslateWithParam(fe validator.FieldError, key string, param string) string {
	t, err := th.T(key, th.fieldName(fe.Field()), param)
	if err != nil {
		return errTranslate(err, fe)
	}
	return t
}

func errTranslate(err error, fe validator.FieldError) string {
	fmt.Printf("warning: error translating FieldError: %s", err)
	return fe.(error).Error()
}

func kindFrom(fe validator.FieldError) reflect.Kind {
	kind := fe.Kind()
	if kind == reflect.Ptr {
		return fe.Type().Elem().Kind()
	}
	return kind
}
