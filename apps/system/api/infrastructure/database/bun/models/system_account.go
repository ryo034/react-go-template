package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SystemAccount struct {
	bun.BaseModel `bun:"table:system_accounts,alias:sa"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Profile       *SystemAccountProfile       `bun:"rel:has-one"`
	AuthProviders []*AuthProvider             `bun:"rel:has-many"`
	Emails        []*SystemAccountEmail       `bun:"rel:has-many"`
	PhoneNumbers  []*SystemAccountPhoneNumber `bun:"rel:has-many"`
}

type AuthProvider struct {
	bun.BaseModel `bun:"table:auth_providers,alias:ap"`

	AuthProviderID  uuid.UUID `bun:"auth_provider_id,pk"`
	SystemAccountID uuid.UUID `bun:"system_account_id,notnull"`
	Provider        string    `bun:"provider,notnull"`
	ProvidedBy      string    `bun:"provided_by,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	SystemAccount *SystemAccount `bun:"sa,rel:belongs-to"`
}

type SystemAccountEmail struct {
	bun.BaseModel `bun:"table:system_account_emails,alias:saes"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	Email           string    `bun:"email,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	SystemAccount *SystemAccount `bun:"sa,rel:belongs-to"`
}

type SystemAccountPhoneNumber struct {
	bun.BaseModel `bun:"table:system_account_phone_numbers,alias:sapns"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	PhoneNumber     string    `bun:"phone_number,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	SystemAccount *SystemAccount `bun:"sa,rel:belongs-to"`
}

type SystemAccountProfile struct {
	bun.BaseModel `bun:"table:system_account_profiles,alias:saps"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	Name            string    `bun:"name"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`

	SystemAccount *SystemAccount `bun:"sa,rel:belongs-to"`
}
