package provider

type Kind string

const (
	Google Kind = "google"
	Email  Kind = "email"
)

type Provider struct {
	id         ID
	kind       Kind
	providedBy ProvidedBy
}

func NewProvider(id ID, kind Kind, providedBy ProvidedBy) *Provider {
	return &Provider{id, kind, providedBy}
}

func NewProviderAsGoogleOnFirebase() (*Provider, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewProvider(id, Google, ProvidedByFirebase), nil
}

func NewProviderAsEmailOnFirebase() (*Provider, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewProvider(id, Email, ProvidedByFirebase), nil

}

func (p *Provider) ID() ID {
	return p.id
}

func (p *Provider) Kind() Kind {
	return p.kind
}

func (p *Provider) ProvidedBy() ProvidedBy {
	return p.providedBy
}
