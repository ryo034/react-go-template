package error

type MessageKey string

const (
	PhoneNumberAlreadyInUseMessageKey MessageKey = "PhoneNumberInUse"
	EmailAlreadyInUseMessageKey       MessageKey = "EmailAlreadyInUse"
	InvalidEmailMessageKey            MessageKey = "InvalidEmail"
	EmailNotVerifiedMessageKey        MessageKey = "EmailNotVerified"
	InvalidAddressMessageKey          MessageKey = "InvalidAddress"
	ConflictedMessageKey              MessageKey = "Conflicted"
	ConflictVersionMessageKey         MessageKey = "ConflictVersion"
	NoSuchDataMessageKey              MessageKey = "NoSuchData"
	UnauthenticatedMessageKey         MessageKey = "Unauthenticated"
	NotBelongMessageKey               MessageKey = "NotBelong"
)

type Error interface {
	error
}
