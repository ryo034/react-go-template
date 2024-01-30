package member

type IDNumber struct {
	v string
}

func NewIDNumber(v string) (IDNumber, error) {
	return IDNumber{v}, nil
}

func (in IDNumber) ToString() string {
	return in.v
}
