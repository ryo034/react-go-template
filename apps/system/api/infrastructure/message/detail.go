package message

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"golang.org/x/text/language"
)

var detailMessages = map[domainError.MessageKey]map[language.Tag]string{
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
	invitation.InvalidInviteTokenMessageKey: {
		language.Japanese: "不正な招待トークンです",
		language.English:  "Invalid invite token",
	},
	invitation.ExpiredInviteTokenMessageKey: {
		language.Japanese: "招待トークンの有効期限が切れています",
		language.English:  "Invite token is expired",
	},
	invitation.AlreadyRevokedInvitationMessageKey: {
		language.Japanese: "招待が既に取り消されています",
		language.English:  "Invite is already revoked",
	},
	invitation.AlreadyVerifiedInvitationMessageKey: {
		language.Japanese: "招待が既に承認されています",
		language.English:  "Invite is already verified",
	},
	invitation.AlreadyAcceptedInvitationMessageKey: {
		language.Japanese: "受諾済みの招待です",
		language.English:  "Invite is already accepted",
	},
	invitation.AlreadyExpiredInvitationMessageKey: {
		language.Japanese: "招待が既に期限切れです",
		language.English:  "Invite is already expired",
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
	domainError.BadRequestMessageKey: {
		language.Japanese: "不正なリクエストです",
		language.English:  "Bad request",
	},
	domainError.ConflictedMessageKey: {
		language.Japanese: "参照している情報が古くなっています。画面を更新して下さい。",
		language.English:  "You are seeing old data. Please reload your current page",
	},
	domainError.ForbiddenMessageKey: {
		language.Japanese: "アクセス権限がありません",
		language.English:  "You don't have permission",
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
	domainError.InternalServerErrorMessageKey: {
		language.Japanese: "サーバーエラーが発生しました",
		language.English:  "Internal server error",
	},
}
