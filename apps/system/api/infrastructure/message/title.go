package message

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"golang.org/x/text/language"
)

var titleMessages = map[domainError.MessageKey]map[language.Tag]string{
	domainError.UnauthenticatedMessageKey: {
		language.Japanese: "認証情報がありません。認証を行ってください",
		language.English:  "No credentials. please authenticate",
	},
	domainError.EmailAlreadyInUseMessageKey: {
		language.Japanese: "すでにそのメールアドレスは使用されています",
		language.English:  "Email already in use",
	},
	domainError.InvalidEmailMessageKey: {
		language.Japanese: "不正なメールアドレスです",
		language.English:  "Invalid email",
	},
	domainError.ExpiredInviteTokenMessageKey: {
		language.Japanese: "招待トークンの有効期限が切れています",
		language.English:  "Invite token is expired",
	},
	domainError.PhoneNumberAlreadyInUseMessageKey: {
		language.Japanese: "すでにその電話番号は使用されています",
		language.English:  "Phone number already in use",
	},
	domainError.InvalidAddressMessageKey: {
		language.Japanese: "不正な住所です",
		language.English:  "Address is invalid",
	},
	domainError.EmailNotVerifiedMessageKey: {
		language.Japanese: "メールアドレスが認証されていません",
		language.English:  "Email is not verified",
	},
	domainError.NotBelongMessageKey: {
		language.Japanese: "このグループに存在していません",
		language.English:  "You are not in group",
	},
	domainError.ConflictedMessageKey: {
		language.Japanese: "参照している情報が古くなっています。画面を更新して下さい。",
		language.English:  "You are seeing old data. Please reload your current page",
	},
	domainError.ConflictVersionMessageKey: {
		language.Japanese: "参照している情報が古くなっています。画面を更新して下さい。",
		language.English:  "You are seeing old data. Please reload your current page",
	},
	domainError.NoSuchDataMessageKey: {
		language.Japanese: "データがありません。画面を更新して下さい。",
		language.English:  "Your requested data doesn't exist. Please reload your current page",
	},
	datetime.InvalidDate: {
		language.Japanese: "不正な日付です",
		language.English:  "date is invalid",
	},
	datetime.InvalidDatetime: {
		language.Japanese: "不正な日付です",
		language.English:  "datetime is invalid",
	},
}
