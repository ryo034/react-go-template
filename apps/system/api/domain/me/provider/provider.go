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
	uid        UID
}

func NewProvider(id ID, kind Kind, providedBy ProvidedBy, uid UID) *Provider {
	return &Provider{id, kind, providedBy, uid}
}

func NewProviderAsGoogleOnFirebase(uid UID) (*Provider, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewProvider(id, Google, ProvidedByFirebase, uid), nil
}

func NewProviderAsEmailOnFirebase(uid UID) (*Provider, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewProvider(id, Email, ProvidedByFirebase, uid), nil

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

func (p *Provider) UID() UID {
	return p.uid
}
