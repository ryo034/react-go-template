package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts,alias:sa"`

	AccountID uuid.UUID `bun:"account_id,pk"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`

	AuthProviders []*AuthProvider           `bun:"ap,rel:has-many"`
	Name          *AccountLatestName        `bun:"aln,rel:has-one"`
	PhotoEvent    *AccountLatestPhotoEvent  `bun:"alphoe,rel:has-one"`
	PhoneNumber   *AccountLatestPhoneNumber `bun:"alpn,rel:has-one"`
	Email         *AccountLatestEmail       `bun:"alem,rel:has-one"`

	// Relation exists but is not used in the application
	//Names        []*AccountName        `bun:"an,rel:has-many"`
	//PhotoEvents  []*AccountPhotoEvent  `bun:"aphoe,rel:has-many"`
	//PhoneNumbers []*AccountPhoneNumber `bun:"apn,rel:has-many"`
	//Emails       []*AccountEmail       `bun:"ae,rel:has-many"`
}

type AuthProvider struct {
	bun.BaseModel `bun:"table:auth_providers,alias:ap"`

	AuthProviderID uuid.UUID `bun:"auth_provider_id,pk"`
	AccountID      uuid.UUID `bun:"account_id,notnull"`
	Provider       string    `bun:"provider,notnull"`
	PhotoURL       string    `bun:"photo_url,notnull"`
	ProviderUID    string    `bun:"provider_uid,notnull"`
	ProvidedBy     string    `bun:"provided_by,notnull"`
	RegisteredAt   time.Time `bun:"registered_at,notnull,default:current_timestamp"`

	Account *Account `bun:"sa,rel:belongs-to"`
}

type AccountEmail struct {
	bun.BaseModel `bun:"table:account_emails,alias:ae"`

	AccountEmailID uuid.UUID `bun:"account_email_id,pk"`
	AccountID      uuid.UUID `bun:"account_id,notnull"`
	Email          string    `bun:"email,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Account            *Account            `bun:"sa,rel:belongs-to"`
	AccountLatestEmail *AccountLatestEmail `bun:"alem,rel:has-one"`
}

type AccountLatestEmail struct {
	bun.BaseModel `bun:"table:account_latest_emails,alias:alem"`

	AccountEmailID uuid.UUID `bun:"account_email_id,pk"`
	AccountID      uuid.UUID `bun:"account_id,notnull,unique"`

	Account      *Account      `bun:"sa,rel:belongs-to"`
	AccountEmail *AccountEmail `bun:"ae,rel:has-one"`
}

type AccountPhoneNumber struct {
	bun.BaseModel `bun:"table:account_phone_numbers,alias:apn"`

	AccountPhoneNumberID uuid.UUID `bun:"account_phone_number_id,pk"`
	AccountID            uuid.UUID `bun:"account_id,notnull"`
	PhoneNumber          string    `bun:"phone_number,notnull"`
	CountryCode          string    `bun:"country_code,notnull"`
	CreatedAt            time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Account                  *Account                  `bun:"sa,rel:belongs-to"`
	AccountLatestPhoneNumber *AccountLatestPhoneNumber `bun:"alpn,rel:has-one"`
}

type AccountLatestPhoneNumber struct {
	bun.BaseModel `bun:"table:account_latest_phone_numbers,alias:alpn"`

	AccountPhoneNumberID uuid.UUID `bun:"account_phone_number_id,pk"`
	AccountID            uuid.UUID `bun:"account_id,notnull,unique"`

	Account            *Account            `bun:"sa,rel:belongs-to"`
	AccountPhoneNumber *AccountPhoneNumber `bun:"apn,rel:belongs-to"`
}

type AccountName struct {
	bun.BaseModel `bun:"table:account_names,alias:an"`

	AccountNameID uuid.UUID `bun:"account_name_id,pk"`
	AccountID     uuid.UUID `bun:"account_id,notnull"`
	Name          string    `bun:"name,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Account           *Account           `bun:"sa,rel:belongs-to"`
	AccountLatestName *AccountLatestName `bun:"aln,rel:has-one"`
}

type AccountLatestName struct {
	bun.BaseModel `bun:"table:account_latest_names,alias:aln"`

	AccountNameID uuid.UUID `bun:"account_name_id,pk"`
	AccountID     uuid.UUID `bun:"account_id,notnull,unique"`

	Account     *Account     `bun:"sa,rel:belongs-to"`
	AccountName *AccountName `bun:"an,rel:belongs-to"`
}

type AccountPhotoEvent struct {
	bun.BaseModel `bun:"table:account_photo_events,alias:aphoe"`

	AccountPhotoEventID uuid.UUID `bun:"account_photo_event_id,pk"`
	AccountID           uuid.UUID `bun:"account_id,notnull"`
	EventType           string    `bun:"event_type,notnull"`
	CreatedAt           time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Account *Account      `bun:"sa,rel:belongs-to"`
	Photo   *AccountPhoto `bun:"apho,rel:has-one"`
}

type AccountPhoto struct {
	bun.BaseModel `bun:"table:account_photos,alias:apho"`

	AccountPhotoEventID uuid.UUID `bun:"account_photo_event_id,pk"`
	PhotoID             uuid.UUID `bun:"photo_id,notnull"`
	HostingTo           string    `bun:"hosting_to,notnull"`

	AccountPhotoEvent *AccountPhotoEvent `bun:"sape,rel:belongs-to"`
}

type AccountLatestPhotoEvent struct {
	bun.BaseModel `bun:"table:account_latest_photo_events,alias:alphoe"`

	AccountPhotoEventID uuid.UUID `bun:"account_photo_event_id,pk"`
	AccountID           uuid.UUID `bun:"account_id,notnull,unique"`

	AccountPhotoEvent *AccountPhotoEvent `bun:"aphoe,rel:belongs-to"`
	Account           *Account           `bun:"sa,rel:belongs-to"`
}
