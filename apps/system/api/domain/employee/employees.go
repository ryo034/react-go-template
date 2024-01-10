package employee

type Employees interface {
	Size() int
	AsSlice() []Employee
	IsEmpty() bool
	IsNotEmpty() bool
}

type employees struct {
	wrapped []Employee
}

func NewEmployees(wrapped []Employee) Employees {
	return &employees{wrapped}
}

func (es *employees) IsEmpty() bool {
	return len(es.wrapped) == 0
}

func (es *employees) IsNotEmpty() bool {
	return !es.IsEmpty()
}

func (es *employees) Size() int {
	return len(es.wrapped)
}

func (es *employees) AsSlice() []Employee {
	return append(make([]Employee, 0, es.Size()), es.wrapped...)
}
