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

	AuthProviders []*AuthProvider                 `bun:"aps,rel:has-many"`
	Name          *SystemAccountLatestName        `bun:"salns,rel:has-one"`
	PhotoEvent    *SystemAccountLatestPhotoEvent  `bun:"salaphoes,rel:has-one"`
	PhoneNumber   *SystemAccountLatestPhoneNumber `bun:"salpns,rel:has-one"`
	Email         *SystemAccountLatestEmail       `bun:"salems,rel:has-one"`

	// Relation exists but is not used in the application
	//Names        []*SystemAccountName        `bun:"sans,rel:has-many"`
	//PhotoEvents  []*SystemAccountPhotoEvent  `bun:"saphoes,rel:has-many"`
	//PhoneNumbers []*SystemAccountPhoneNumber `bun:"sapns,rel:has-many"`
	//Emails       []*SystemAccountEmail       `bun:"saes,rel:has-many"`
}

type AuthProvider struct {
	bun.BaseModel `bun:"table:auth_providers,alias:aps"`

	AuthProviderID  uuid.UUID `bun:"auth_provider_id,pk"`
	SystemAccountID uuid.UUID `bun:"system_account_id,notnull"`
	Provider        string    `bun:"provider,notnull"`
	ProviderUID     string    `bun:"provider_uid,notnull"`
	ProvidedBy      string    `bun:"provided_by,notnull"`
	RegisteredAt    time.Time `bun:"registered_at,notnull,default:current_timestamp"`

	SystemAccount *SystemAccount `bun:"sa,rel:belongs-to"`
}

type SystemAccountEmail struct {
	bun.BaseModel `bun:"table:system_account_emails,alias:saes"`

	SystemAccountEmailID uuid.UUID `bun:"system_account_email_id,pk"`
	SystemAccountID      uuid.UUID `bun:"system_account_id,notnull"`
	Email                string    `bun:"email,notnull"`
	CreatedAt            time.Time `bun:"created_at,notnull,default:current_timestamp"`

	SystemAccount            *SystemAccount            `bun:"sa,rel:belongs-to"`
	SystemAccountLatestEmail *SystemAccountLatestEmail `bun:"salems,rel:has-one"`
}

type SystemAccountLatestEmail struct {
	bun.BaseModel `bun:"table:system_account_latest_emails,alias:salems"`

	SystemAccountLatestEmailID uuid.UUID `bun:"system_account_latest_email_id,pk"`
	SystemAccountID            uuid.UUID `bun:"system_account_id,notnull,unique"`

	SystemAccount      *SystemAccount      `bun:"sa,rel:belongs-to"`
	SystemAccountEmail *SystemAccountEmail `bun:"saes,rel:belongs-to"`
}

type SystemAccountPhoneNumber struct {
	bun.BaseModel `bun:"table:system_account_phone_numbers,alias:sapns"`

	SystemAccountPhoneNumberID uuid.UUID `bun:"system_account_phone_number_id,pk"`
	SystemAccountID            uuid.UUID `bun:"system_account_id,notnull"`
	PhoneNumber                string    `bun:"phone_number,notnull"`
	CountryCode                string    `bun:"country_code,notnull"`
	CreatedAt                  time.Time `bun:"created_at,notnull,default:current_timestamp"`

	SystemAccount                  *SystemAccount                  `bun:"sa,rel:belongs-to"`
	SystemAccountLatestPhoneNumber *SystemAccountLatestPhoneNumber `bun:"salpns,rel:has-one"`
}

type SystemAccountLatestPhoneNumber struct {
	bun.BaseModel `bun:"table:system_account_latest_phone_numbers,alias:salpns"`

	SystemAccountPhoneNumberID uuid.UUID `bun:"system_account_phone_number_id,pk"`
	SystemAccountID            uuid.UUID `bun:"system_account_id,notnull,unique"`

	SystemAccount            *SystemAccount            `bun:"sa,rel:belongs-to"`
	SystemAccountPhoneNumber *SystemAccountPhoneNumber `bun:"sapns,rel:belongs-to"`
}

type SystemAccountName struct {
	bun.BaseModel `bun:"table:system_account_names,alias:sans"`

	SystemAccountNameID uuid.UUID `bun:"system_account_name_id,pk"`
	SystemAccountID     uuid.UUID `bun:"system_account_id,notnull"`
	Name                string    `bun:"name,notnull"`
	CreatedAt           time.Time `bun:"created_at,notnull,default:current_timestamp"`

	SystemAccount           *SystemAccount           `bun:"sa,rel:belongs-to"`
	SystemAccountLatestName *SystemAccountLatestName `bun:"salns,rel:has-one"`
}

type SystemAccountLatestName struct {
	bun.BaseModel `bun:"table:system_account_latest_names,alias:salns"`

	SystemAccountNameID uuid.UUID `bun:"system_account_name_id,pk"`
	SystemAccountID     uuid.UUID `bun:"system_account_id,notnull,unique"`

	SystemAccount     *SystemAccount     `bun:"sa,rel:belongs-to"`
	SystemAccountName *SystemAccountName `bun:"sans,rel:belongs-to"`
}

type SystemAccountPhotoEvent struct {
	bun.BaseModel `bun:"table:system_account_photo_events,alias:saphoes"`

	SystemAccountPhotoEventID uuid.UUID `bun:"system_account_photo_event_id,pk"`
	SystemAccountID           uuid.UUID `bun:"system_account_id,notnull"`
	EventType                 string    `bun:"event_type,notnull"`
	CreatedAt                 time.Time `bun:"created_at,notnull,default:current_timestamp"`

	SystemAccount *SystemAccount      `bun:"sa,rel:belongs-to"`
	Photo         *SystemAccountPhoto `bun:"saphos,rel:has-one"`
}

type SystemAccountPhoto struct {
	bun.BaseModel `bun:"table:system_account_photos,alias:saphos"`

	SystemAccountPhotoEventID uuid.UUID `bun:"system_account_photo_event_id,pk"`
	PhotoPath                 string    `bun:"photo_path,notnull"`

	SystemAccountPhotoEvent *SystemAccountPhotoEvent `bun:"sape,rel:belongs-to"`
}

type SystemAccountLatestPhotoEvent struct {
	bun.BaseModel `bun:"table:system_account_latest_photo_events,alias:salaphoes"`

	SystemAccountPhotoEventID uuid.UUID `bun:"system_account_photo_event_id,pk"`
	SystemAccountID           uuid.UUID `bun:"system_account_id,notnull,unique"`

	SystemAccountPhotoEvent *SystemAccountPhotoEvent `bun:"saphoes,rel:belongs-to"`
	SystemAccount           *SystemAccount           `bun:"sa,rel:belongs-to"`
}
