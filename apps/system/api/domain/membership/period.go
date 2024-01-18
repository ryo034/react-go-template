package membership

type Period struct {
	startDate string
	endDate   string
}

func NewPeriod(startDate string, endDate string) *Period {
	return &Period{startDate, endDate}
}

func (p *Period) StartDate() string {
	return p.startDate
}

func (p *Period) EndDate() string {
	return p.endDate
}
