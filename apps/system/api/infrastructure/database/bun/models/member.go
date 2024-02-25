package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

const (
	MemberTableName             = "members"
	MemberProfileTableName      = "member_profiles"
	MemberLoginHistoryTableName = "member_login_histories"
	MemberAddressTableName      = "member_addresses"
)

type Member struct {
	bun.BaseModel `bun:"table:members,alias:ms"`

	MemberID        uuid.UUID `bun:"member_id,pk"`
	WorkspaceID     uuid.UUID `bun:"workspace_id,notnull"`
	SystemAccountID uuid.UUID `bun:"system_account_id,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	SystemAccount *SystemAccount `bun:"sa,rel:belongs-to"`
	Profile       *MemberProfile `bun:"mp,rel:has-one"`
	Workspace     *Workspace     `bun:"ws,rel:belongs-to"`
}

type Members []*Member

type MemberProfile struct {
	bun.BaseModel `bun:"table:member_profiles,alias:mp"`

	MemberID       uuid.UUID `bun:"member_id,pk"`
	MemberIDNumber string    `bun:"member_id_number"`
	DisplayName    string    `bun:"display_name,notnull"`
	Bio            string    `bun:"bio,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:"updated_at,notnull,default:current_timestamp"`

	Member *Member `bun:"rel:has-one"`
}

type MemberLoginHistory struct {
	bun.BaseModel `bun:"table:member_login_histories,alias:llw"`

	MemberLoginHistoryID uuid.UUID `bun:"member_login_history_id,pk"`
	MemberID             uuid.UUID `bun:"member_id,notnull"`
	LoginAt              time.Time `bun:"login_at,notnull,default:current_timestamp"`

	Member *Member `bun:"rel:belongs-to"`
}

type MemberAddress struct {
	bun.BaseModel `bun:"table:member_addresses,alias:ma"`

	MemberID                 uuid.UUID `bun:"member_id,pk"`
	PostalCode               string    `bun:"postal_code"`
	BuildingComponentID      uuid.UUID `bun:"building_component_id"`
	StreetAddressComponentID uuid.UUID `bun:"street_address_component_id,notnull"`
	CityComponentID          uuid.UUID `bun:"city_component_id,notnull"`
	StateComponentID         uuid.UUID `bun:"state_component_id,notnull"`
	CountryComponentID       uuid.UUID `bun:"country_component_id,notnull"`
	CreatedAt                time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Member   *Member           `bun:"rel:belongs-to"`
	Country  *AddressComponent `bun:"rel:belongs-to,join:country_component_id=component_id"`
	State    *AddressComponent `bun:"rel:belongs-to,join:state_component_id=component_id"`
	City     *AddressComponent `bun:"rel:belongs-to,join:city_component_id=component_id"`
	Street   *AddressComponent `bun:"rel:belongs-to,join:street_address_component_id=component_id"`
	Building *AddressComponent `bun:"rel:belongs-to,join:building_component_id=component_id"`
}
