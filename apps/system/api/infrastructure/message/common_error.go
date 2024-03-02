package message

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"golang.org/x/text/language"
)

const (
	Required            domainError.MessageKey = "required"
	Alpha               domainError.MessageKey = "alpha"
	Alphanum            domainError.MessageKey = "alphanum"
	Numeric             domainError.MessageKey = "numeric"
	Number              domainError.MessageKey = "number"
	Email               domainError.MessageKey = "email"
	Latitude            domainError.MessageKey = "latitude"
	Longitude           domainError.MessageKey = "longitude"
	Eqfield             domainError.MessageKey = "eqfield"
	Eqcsfield           domainError.MessageKey = "eqcsfield"
	Necsfield           domainError.MessageKey = "necsfield"
	Gtcsfield           domainError.MessageKey = "gtcsfield"
	Gtecsfield          domainError.MessageKey = "gtecsfield"
	Ltcsfield           domainError.MessageKey = "ltcsfield"
	Ltecsfield          domainError.MessageKey = "ltecsfield"
	Nefield             domainError.MessageKey = "nefield"
	Gtfield             domainError.MessageKey = "gtfield"
	Gtefield            domainError.MessageKey = "gtefield"
	Ltfield             domainError.MessageKey = "ltfield"
	Ltefield            domainError.MessageKey = "ltefield"
	Contains            domainError.MessageKey = "contains"
	Containsany         domainError.MessageKey = "containsany"
	Excludes            domainError.MessageKey = "excludes"
	Excludesall         domainError.MessageKey = "excludesall"
	Excludesrune        domainError.MessageKey = "excludesrune"
	Oneof               domainError.MessageKey = "oneof"
	Eq                  domainError.MessageKey = "eq"
	Ne                  domainError.MessageKey = "ne"
	CharacterCountOne   domainError.MessageKey = "character-count-one"
	CharacterCountOther domainError.MessageKey = "character-count-other"
	ItemCountOne        domainError.MessageKey = "item-count-one"
	ItemCountOther      domainError.MessageKey = "item-count-other"
	LenString           domainError.MessageKey = "len-string"
	LenNumber           domainError.MessageKey = "len-number"
	LenItems            domainError.MessageKey = "len-items"
	MinString           domainError.MessageKey = "min-string"
	MinNumber           domainError.MessageKey = "min-number"
	MinItems            domainError.MessageKey = "min-items"
	MaxString           domainError.MessageKey = "max-string"
	MaxNumber           domainError.MessageKey = "max-number"
	MaxItems            domainError.MessageKey = "max-items"
	LtString            domainError.MessageKey = "lt-string"
	LtNumber            domainError.MessageKey = "lt-number"
	LtItems             domainError.MessageKey = "lt-items"
	LteString           domainError.MessageKey = "lte-string"
	LteNumber           domainError.MessageKey = "lte-number"
	LteItems            domainError.MessageKey = "lte-items"
	GtString            domainError.MessageKey = "gt-string"
	GtNumber            domainError.MessageKey = "gt-number"
	GtItems             domainError.MessageKey = "gt-items"
	GteString           domainError.MessageKey = "gte-string"
	GteNumber           domainError.MessageKey = "gte-number"
	GteItems            domainError.MessageKey = "gte-items"
	Time                domainError.MessageKey = "time"
	Date                domainError.MessageKey = "date"
	FacilityAuth        domainError.MessageKey = "facility-authentication"
)

var commonErrorMessages = map[string]map[language.Tag]string{
	// messages for github.com/go-playground/validator/v10
	string(Required): {
		language.Japanese: "{0}は必須です",
		language.English:  "{0} is a required field",
	},
	string(Alpha): {
		language.Japanese: "{0}はアルファベットのみを含むことができます",
		language.English:  "{0} can only contain alphabetic characters",
	},
	string(Alphanum): {
		language.Japanese: "{0}はアルファベットと数字のみを含むことができます",
		language.English:  "{0} can only contain alphanumeric characters",
	},
	string(Numeric): {
		language.Japanese: "{0}は正しい数字でなければなりません",
		language.English:  "{0} must be a valid numeric value",
	},
	string(Number): {
		language.Japanese: "{0}は正しい数でなければなりません",
		language.English:  "{0} must be a valid number",
	},
	string(Email): {
		language.Japanese: "{0}は正しいメールアドレスでなければなりません",
		language.English:  "{0} must be a valid email address",
	},
	string(Latitude): {
		language.Japanese: "{0}は正しい緯度の座標を含まなければなりません",
		language.English:  "{0} must contain valid latitude coordinates",
	},
	string(Longitude): {
		language.Japanese: "{0}は正しい経度の座標を含まなければなりません",
		language.English:  "{0} must contain a valid longitude coordinates",
	},
	string(Eqfield): {
		language.Japanese: "{0}は{1}と等しくなければなりません",
		language.English:  "{0} must be equal to {1}",
	},
	string(Eqcsfield): {
		language.Japanese: "{0}は{1}と等しくなければなりません",
		language.English:  "{0} must be equal to {1}",
	},
	string(Necsfield): {
		language.Japanese: "{0}は{1}とは異ならなければなりません",
		language.English:  "{0} cannot be equal to {1}",
	},
	string(Gtcsfield): {
		language.Japanese: "{0}は{1}よりも大きくなければなりません",
		language.English:  "{0} must be greater than {1}",
	},
	string(Gtecsfield): {
		language.Japanese: "{0}は{1}以上でなければなりません",
		language.English:  "{0} must be greater than or equal to {1}",
	},
	string(Ltcsfield): {
		language.Japanese: "{0}は{1}よりも小さくなければなりません",
		language.English:  "{0} must be less than {1}",
	},
	string(Ltecsfield): {
		language.Japanese: "{0}は{1}以下でなければなりません",
		language.English:  "{0} must be less than or equal to {1}",
	},
	string(Nefield): {
		language.Japanese: "{0}は{1}とは異ならなければなりません",
		language.English:  "{0} cannot be equal to {1}",
	},
	string(Gtfield): {
		language.Japanese: "{0}は{1}よりも大きくなければなりません",
		language.English:  "{0} must be greater than {1}",
	},
	string(Gtefield): {
		language.Japanese: "{0}は{1}以上でなければなりません",
		language.English:  "{0} must be greater than or equal to {1}",
	},
	string(Ltfield): {
		language.Japanese: "{0}は{1}よりも小さくなければなりません",
		language.English:  "{0} must be less than {1}",
	},
	string(Ltefield): {
		language.Japanese: "{0}は{1}以下でなければなりません",
		language.English:  "{0} must be less than or equal to {1}",
	},
	string(Contains): {
		language.Japanese: "{0}は'{1}'を含まなければなりません",
		language.English:  "{0} must contain the text '{1}'",
	},
	string(Containsany): {
		language.Japanese: "{0}は'{1}'の少なくとも1つを含まなければなりません",
		language.English:  "{0} must contain at least one of the following characters '{1}'",
	},
	string(Excludes): {
		language.Japanese: "{0}には'{1}'というテキストを含むことはできません",
		language.English:  "{0} cannot contain the text '{1}'",
	},
	string(Excludesall): {
		language.Japanese: "{0}には'{1}'のどれも含めることはできません",
		language.English:  "{0} cannot contain any of the following characters '{1}'",
	},
	string(Excludesrune): {
		language.Japanese: "{0}には'{1}'を含めることはできません",
		language.English:  "{0} cannot contain the following '{1}'",
	},
	string(Oneof): {
		language.Japanese: "{0}は[{1}]のうちのいずれかでなければなりません",
		language.English:  "{0} must be one of [{1}]",
	},
	string(Eq): {
		language.Japanese: "{0}は{1}と等しくありません",
		language.English:  "{0} is not equal to {1}",
	},
	string(Ne): {
		language.Japanese: "{0}は{1}と異ならなければなりません",
		language.English:  "{0} should not be equal to {1}",
	},
	string(CharacterCountOne): {
		language.English: "{0} character",
	},
	string(CharacterCountOther): {
		language.Japanese: "{0}文字",
		language.English:  "{0} characters",
	},
	string(ItemCountOne): {
		language.English: "{0} item",
	},
	string(ItemCountOther): {
		language.Japanese: "{0}つの項目",
		language.English:  "{0} items",
	},
	string(LenString): {
		language.Japanese: "{0}の長さは{1}でなければなりません",
		language.English:  "{0} must be {1} in length",
	},
	string(LenNumber): {
		language.Japanese: "{0}は{1}と等しくなければなりません",
		language.English:  "{0} must be equal to {1}",
	},
	string(LenItems): {
		language.Japanese: "{0}は{1}を含まなければなりません",
		language.English:  "{0} must contain {1}",
	},
	string(MinString): {
		language.Japanese: "{0}の長さは少なくとも{1}はなければなりません",
		language.English:  "{0} must be at least {1} in length",
	},
	string(MinNumber): {
		language.Japanese: "{0}は{1}かより大きくなければなりません",
		language.English:  "{0} must be {1} or greater",
	},
	string(MinItems): {
		language.Japanese: "{0}は少なくとも{1}を含まなければなりません",
		language.English:  "{0} must contain at least {1}",
	},
	string(MaxString): {
		language.Japanese: "{0}の長さは最大でも{1}でなければなりません",
		language.English:  "{0} must be a maximum of {1} in length",
	},
	string(MaxNumber): {
		language.Japanese: "{0}は{1}かより小さくなければなりません",
		language.English:  "{0} must be {1} or less",
	},
	string(MaxItems): {
		language.Japanese: "{0}は最大でも{1}を含まなければなりません",
		language.English:  "{0} must contain at maximum {1}",
	},
	string(LtString): {
		language.Japanese: "{0}の長さは{1}よりも少なくなければなりません",
		language.English:  "{0} must be less than {1} in length",
	},
	string(LtNumber): {
		language.Japanese: "{0}は{1}よりも小さくなければなりません",
		language.English:  "{0} must be less than {1}",
	},
	string(LtItems): {
		language.Japanese: "{0}は{1}よりも少ない項目を含まなければなりません",
		language.English:  "{0} must contain less than {1}",
	},
	string(LteString): {
		language.Japanese: "{0}の長さは最大でも{1}でなければなりません",
		language.English:  "{0} must be at maximum {1} in length",
	},
	string(LteNumber): {
		language.Japanese: "{0}は{1}かより小さくなければなりません",
		language.English:  "{0} must be {1} or less",
	},
	string(LteItems): {
		language.Japanese: "{0}は最大でも{1}を含まなければなりません",
		language.English:  "{0} must contain at maximum {1}",
	},
	string(GtString): {
		language.Japanese: "{0}の長さは{1}よりも多くなければなりません",
		language.English:  "{0} must be greater than {1} in length",
	},
	string(GtNumber): {
		language.Japanese: "{0}は{1}よりも大きくなければなりません",
		language.English:  "{0} must be greater than {1}",
	},
	string(GtItems): {
		language.Japanese: "{0}は{1}よりも多い項目を含まなければなりません",
		language.English:  "{0} must contain more than {1}",
	},
	string(GteString): {
		language.Japanese: "{0}の長さは少なくとも{1}以上はなければなりません",
		language.English:  "{0} must be at least {1} in length",
	},
	string(GteNumber): {
		language.Japanese: "{0}は{1}かより大きくなければなりません",
		language.English:  "{0} must be {1} or greater",
	},
	string(GteItems): {
		language.Japanese: "{0}は少なくとも{1}を含まなければなりません",
		language.English:  "{0} must contain at least {1}",
	},
	string(Time): {
		language.Japanese: "{0}はhh:mm形式で入力してください",
		language.English:  "{0} should be hh:mm format",
	},
	string(Date): {
		language.Japanese: "{0}はyyyy-MM-dd形式で入力してください",
		language.English:  "{0} should be yyyy-MM-dd format",
	},
	string(FacilityAuth): {
		language.Japanese: "施設ID:%vの権限がありません",
		language.English:  "accessing to the facility:%v not allowed",
	},
}
