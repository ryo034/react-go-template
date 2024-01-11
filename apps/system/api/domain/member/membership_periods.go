package member

type MembershipPeriods interface {
	Size() int
	AsSlice() []MembershipPeriod
	IsEmpty() bool
	IsNotEmpty() bool
}

type membershipPeriods struct {
	wrapped []MembershipPeriod
}

func NewMembershipPeriods(wrapped []MembershipPeriod) MembershipPeriods {
	return &membershipPeriods{wrapped}
}

func (ms *membershipPeriods) IsEmpty() bool {
	return len(ms.wrapped) == 0
}

func (ms *membershipPeriods) IsNotEmpty() bool {
	return !ms.IsEmpty()
}

func (ms *membershipPeriods) Size() int {
	return len(ms.wrapped)
}

func (ms *membershipPeriods) AsSlice() []MembershipPeriod {
	return append(make([]MembershipPeriod, 0, ms.Size()), ms.wrapped...)
}
