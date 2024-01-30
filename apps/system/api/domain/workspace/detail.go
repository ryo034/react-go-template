package workspace

type Detail struct {
	name Name
}

func NewDetail(name Name) Detail {
	return Detail{name}
}

func (d *Detail) Name() Name {
	return d.name
}
