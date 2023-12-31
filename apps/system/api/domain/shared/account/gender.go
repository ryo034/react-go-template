package account

type Gender int

const (
	GenderMan Gender = iota + 1
	GenderWoman
	GenderUnknown
)

func (g Gender) EnumIndex() int {
	return int(g)
}
