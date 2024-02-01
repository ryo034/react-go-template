package workspace

type Detail struct {
	name      Name
	subdomain Subdomain
}

func NewDetail(name Name, subdomain Subdomain) Detail {
	return Detail{name, subdomain}
}

func (d *Detail) Name() Name {
	return d.name
}

func (d *Detail) Subdomain() Subdomain {
	return d.subdomain
}
