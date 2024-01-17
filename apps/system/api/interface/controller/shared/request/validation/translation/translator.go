package translation

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"golang.org/x/text/language"
)

type TranslatorHolder interface {
	Translator(tag language.Tag) Translator
}

type translatorHolder map[language.Tag]Translator

func (t translatorHolder) Translator(tag language.Tag) Translator {
	return t[tag]
}

type Translator interface {
	Translate(fe validator.FieldError) string
}

type translator map[string]func(fe validator.FieldError) string

func (t translator) Translate(fe validator.FieldError) string {
	return t[fe.Tag()](fe)
}

func NewTranslatorHolder(mr message.Resource) (TranslatorHolder, error) {
	uni := ut.New(ja.New(), en.New())
	jaHelper, err := defaultHelper(language.Japanese, uni, mr)
	if err != nil {
		return nil, err
	}
	enHelper, err := enHelper(uni, mr)
	if err != nil {
		return nil, err
	}
	return newTranslatorHolder(jaHelper, enHelper)
}

func newTranslatorHolder(helpers ...*helper) (TranslatorHolder, error) {
	var th translatorHolder = make(map[language.Tag]Translator, len(helpers))
	for _, helper := range helpers {
		if err := registerMessages(helper); err != nil {
			return nil, err
		}
		th[helper.lTag] = newTranslator(helper)
	}
	return th, nil
}

func defaultHelper(lTag language.Tag, uni *ut.UniversalTranslator, mr message.Resource) (*helper, error) {
	trans, _ := uni.GetTranslator(lTag.String())
	return newHelper(trans, mr)
}

func enHelper(uni *ut.UniversalTranslator, mr message.Resource) (*helper, error) {
	helper, err := defaultHelper(language.English, uni, mr)
	if err != nil {
		return nil, err
	}
	rHook := func() error {
		return helper.addCardinalAll(locales.PluralRuleOne, "character-count", "item-count")
	}
	helper.setRegisterMessagesHook(rHook)
	return helper, nil
}

func registerMessages(helper *helper) error {
	helper.addAll(simpleTags...)
	helper.addAll(withParamTags...)
	helper.addAll(fieldCompareTags...)
	if err := helper.addCardinalAll(locales.PluralRuleOther, "character-count", "item-count"); err != nil {
		return err
	}
	for _, prefix := range []string{"len", "min", "max", "lt", "lte", "gt", "gte"} {
		helper.addAll(prefix+"-string", prefix+"-number", prefix+"-items")
	}
	return helper.registerMessagesHook()
}

func newTranslator(helper *helper) translator {
	var translator translator = make(map[string]func(fe validator.FieldError) string, 100)
	for _, tag := range simpleTags {
		translator[tag] = helper.translate
	}
	for _, tag := range withParamTags {
		translator[tag] = helper.translateWithParam
	}
	for _, tag := range fieldCompareTags {
		translator[tag] = helper.translateWithFieldNameParam
	}
	for _, tag := range []string{"len", "min", "max", "lt", "lte", "gt", "gte"} {
		translator[tag] = helper.translateCompareError
	}
	return translator
}
