package error

type MessageKey string

const (
	PhoneNumberAlreadyInUseMessageKey MessageKey = "PhoneNumberInUse"
	EmailAlreadyInUseMessageKey       MessageKey = "EmailAlreadyInUse"
	InvalidEmailMessageKey            MessageKey = "InvalidEmail"
	EmailNotVerifiedMessageKey        MessageKey = "EmailNotVerified"
	InvalidAddressMessageKey          MessageKey = "InvalidAddress"
	ConflictedMessageKey              MessageKey = "Conflicted"
	ForbiddenMessageKey               MessageKey = "Forbidden"
	ConflictVersionMessageKey         MessageKey = "ConflictVersion"
	NoSuchDataMessageKey              MessageKey = "NoSuchData"
	UnauthenticatedMessageKey         MessageKey = "Unauthenticated"
	NotBelongMessageKey               MessageKey = "NotBelong"
	InternalServerErrorMessageKey     MessageKey = "InternalServerError"
)

type Error interface {
	error
}
