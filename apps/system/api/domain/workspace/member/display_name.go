package member

type DisplayName struct {
	v string
}

func NewDisplayName(v string) (DisplayName, error) {
	return DisplayName{v}, nil
}

func (dn DisplayName) ToString() string {
	return dn.v
}

func (dn DisplayName) IsEmpty() bool {
	return dn.v == ""
}
