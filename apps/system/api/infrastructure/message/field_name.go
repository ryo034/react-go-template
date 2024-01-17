package message

import "golang.org/x/text/language"

type FieldNameTag string

const (
	DeviceToken FieldNameTag = "DeviceToken"
)

var filedNames = map[string]map[language.Tag]string{
	string(DeviceToken): {
		language.Japanese: "デバイストークン",
		language.English:  "Device Token",
	},
}
