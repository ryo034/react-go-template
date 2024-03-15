package error

type MessageKey string

const (
	PhoneNumberAlreadyInUseMessageKey MessageKey = "PhoneNumberInUse"
	EmailAlreadyInUseMessageKey       MessageKey = "EmailAlreadyInUse"
	InvalidEmailMessageKey            MessageKey = "InvalidEmail"
	EmailNotVerifiedMessageKey        MessageKey = "EmailNotVerified"
	InvalidAddressMessageKey          MessageKey = "InvalidAddress"
	ConflictVersionMessageKey         MessageKey = "ConflictVersion"
	NotBelongMessageKey               MessageKey = "NotBelong"

	BadRequestMessageKey          MessageKey = "BadRequest"
	UnauthenticatedMessageKey     MessageKey = "Unauthenticated"
	ForbiddenMessageKey           MessageKey = "Forbidden"
	NoSuchDataMessageKey          MessageKey = "NoSuchData"
	ConflictedMessageKey          MessageKey = "Conflicted"
	GoneMessageKey                MessageKey = "Gone"
	InternalServerErrorMessageKey MessageKey = "InternalServerError"
)

type Error interface {
	error
}
