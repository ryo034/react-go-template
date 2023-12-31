package id

import (
	"github.com/oklog/ulid"
	domainError "github.com/ryo034/react-go-template/packages/go/domain/shared/error"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/validation"
	"math/rand"
	"time"
)

type SortableID struct {
	v string
}

const InvalidSortableID domainError.MessageKey = "invalid.sortable_id"

func NewSortableID(v string) (SortableID, error) {
	errs := validation.NewErrors()
	ul, err := ulid.Parse(v)
	if err != nil {
		errs.Append(InvalidSortableID, v)
	}
	if errs.IsNotEmpty() {
		return SortableID{}, errs
	}
	return SortableID{ul.String()}, nil
}

func GenStringID() SortableID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return SortableID{id.String()}
}

func (s SortableID) ToString() string {
	return s.v
}
