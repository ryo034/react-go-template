package member

type MembershipPeriod struct {
	startDate string
	endDate   string
}

func NewMembershipPeriod(startDate string, endDate string) *MembershipPeriod {
	return &MembershipPeriod{startDate, endDate}
}

func (mp *MembershipPeriod) StartDate() string {
	return mp.startDate
}

func (mp *MembershipPeriod) EndDate() string {
	return mp.endDate
}
