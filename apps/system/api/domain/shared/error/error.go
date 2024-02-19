package error

type MessageKey string

const (
	PhoneNumberAlreadyInUseMessageKey MessageKey = "PhoneNumberInUse"
	EmailAlreadyInUseMessageKey       MessageKey = "EmailAlreadyInUse"
	InvalidEmailMessageKey            MessageKey = "InvalidEmail"
	InvalidInviteTokenMessageKey      MessageKey = "InvalidInviteToken"
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
