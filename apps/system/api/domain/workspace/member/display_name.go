package member

type DisplayName struct {
	v string
}

func NewDisplayName(v string) *DisplayName {
	if v == "" {
		return nil
	}
	return &DisplayName{v}
}

func (dn DisplayName) ToString() string {
	return dn.v
}

func (dn DisplayName) IsEmpty() bool {
	return dn.v == ""
}

func (dn DisplayName) IsNotEmpty() bool {
	return !dn.IsEmpty()
}
