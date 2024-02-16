package error

type MessageKey string

const (
	PhoneNumberAlreadyInUseMessageKey MessageKey = "PhoneNumberInUse"
	EmailAlreadyInUseMessageKey       MessageKey = "EmailAlreadyInUse"
	InvalidEmailMessageKey            MessageKey = "InvalidEmail"
	InvalidInviteTokenMessageKey      MessageKey = "InvalidInviteToken"
	ExpiredInviteTokenMessageKey      MessageKey = "ExpiredInviteToken"
	EmailNotVerifiedMessageKey        MessageKey = "EmailNotVerified"
	InvalidAddressMessageKey          MessageKey = "InvalidAddress"
	ConflictedMessageKey              MessageKey = "Conflicted"
	ConflictVersionMessageKey         MessageKey = "ConflictVersion"
	NoSuchDataMessageKey              MessageKey = "NoSuchData"
	UnauthenticatedMessageKey         MessageKey = "Unauthenticated"
	NotBelongMessageKey               MessageKey = "NotBelong"
	InternalServerErrorMessageKey     MessageKey = "InternalServerError"
)

type Error interface {
	error
}

type Code string

const (
	BasicCodeKey              Code = "000"
	ExpiredInviteTokenCodeKey Code = "001"
	EmailNotVerifiedCodeKey   Code = "002"
	EmailAlreadyInUseCodeKey  Code = "003"
	PhoneNumberInUseCodeKey   Code = "004"
	InvalidEmailCodeKey       Code = "005"
	InvalidAddressCodeKey     Code = "006"
	ConflictVersionCodeKey    Code = "007"
	NotBelongCodeKey          Code = "008"
)
