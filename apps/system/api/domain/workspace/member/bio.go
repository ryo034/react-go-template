package member

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

const MaxLength = 2000
const InvalidMemberBio domainError.MessageKey = "invalid.member.bio"

type Bio struct {
	v string
}

func NewBio(v string) (Bio, error) {
	errs := validation.NewErrors()
	if len(v) > MaxLength {
		errs.Append(InvalidMemberBio, v)
	}
	if errs.IsNotEmpty() {
		return Bio{}, errs
	}
	return Bio{v}, nil
}

func (b Bio) ToString() string {
	return b.v
}

func NewAsEmptyBio() Bio {
	return Bio{""}
}
