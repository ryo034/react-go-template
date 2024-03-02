package provider

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

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

func NewProviderAsEmailOnFirebase(aID account.ID) (*Provider, error) {
	apUID, err := NewUID(aID.Value().String())
	if err != nil {
		return nil, err
	}
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewProvider(id, Email, ProvidedByFirebase, apUID), nil
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
