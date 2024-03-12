package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Member struct {
	bun.BaseModel `bun:"table:members,alias:ms"`

	MemberID    uuid.UUID `bun:"member_id,pk"`
	WorkspaceID uuid.UUID `bun:"workspace_id,notnull"`
	AccountID   uuid.UUID `bun:"account_id,notnull"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Account         *Account               `bun:"sa,rel:belongs-to"`
	Profile         *MemberLatestProfile   `bun:"mlp,rel:has-one"`
	Workspace       *Workspace             `bun:"ws,rel:belongs-to"`
	Role            *MemberLatestRole      `bun:"mlr,rel:has-one"`
	MembershipEvent *LatestMembershipEvent `bun:"lmshi,rel:has-one"`
}

type Members []*Member

type MemberProfile struct {
	bun.BaseModel `bun:"table:member_profiles,alias:mp"`

	MemberProfileID uuid.UUID `bun:"member_profile_id,pk"`
	MemberID        uuid.UUID `bun:"member_id,notnull"`
	MemberIDNumber  string    `bun:"member_id_number"`
	DisplayName     string    `bun:"display_name,notnull"`
	Bio             string    `bun:"bio,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Member        *Member              `bun:"ms,rel:belongs-to"`
	MemberProfile *MemberLatestProfile `bun:"mlp,rel:has-one"`
}

type MemberLatestProfile struct {
	bun.BaseModel `bun:"table:member_latest_profiles,alias:mlp"`

	MemberProfileID uuid.UUID `bun:"member_profile_id,pk"`
	MemberID        uuid.UUID `bun:"member_id,notnull"`

	MemberProfile *MemberProfile `bun:"mp,rel:has-one"`
	Member        *Member        `bun:"ms,rel:belongs-to"`
}

type MemberRole struct {
	bun.BaseModel `bun:"table:member_roles,alias:mr"`

	MemberRoleID uuid.UUID `bun:"member_role_id,pk"`
	MemberID     uuid.UUID `bun:"member_id,notnull"`
	Role         string    `bun:"role,notnull"`
	AssignedAt   time.Time `bun:"assigned_at,notnull,default:current_timestamp"`
	AssignedBy   uuid.UUID `bun:"assigned_by,notnull"`

	Member           *Member           `bun:"ms,rel:belongs-to"`
	MemberLatestRole *MemberLatestRole `bun:"mlr,rel:has-one"`
	Assignor         *Member           `bun:"ms,rel:belongs-to,join:assigned_by=member_id"`
}

type MemberLatestRole struct {
	bun.BaseModel `bun:"table:member_latest_roles,alias:mlr"`

	MemberRoleID uuid.UUID `bun:"member_role_id,pk"`
	MemberID     uuid.UUID `bun:"member_id,notnull"`

	MemberRole *MemberRole `bun:"mr,rel:has-one"`
	Member     *Member     `bun:"ms,rel:belongs-to"`
}

type MemberLoginHistory struct {
	bun.BaseModel `bun:"table:member_login_histories,alias:mllhs"`

	MemberLoginHistoryID uuid.UUID `bun:"member_login_history_id,pk"`
	MemberID             uuid.UUID `bun:"member_id,notnull"`
	LoginAt              time.Time `bun:"login_at,notnull,default:current_timestamp"`

	Member          *Member                   `bun:"ms,rel:belongs-to"`
	MemberLastLogin *MemberLatestLoginHistory `bun:"mllhs,rel:has-one"`
}

type MemberLatestLoginHistory struct {
	bun.BaseModel `bun:"table:member_latest_login_histories,alias:mllhs"`

	MemberLoginHistoryID uuid.UUID `bun:"member_login_history_id,pk"`
	MemberID             uuid.UUID `bun:"member_id,notnull"`

	Member             *Member             `bun:"ms,rel:belongs-to"`
	MemberLoginHistory *MemberLoginHistory `bun:"mllhs,rel:belongs-to"`
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

type MembershipEvent struct {
	bun.BaseModel `bun:"table:membership_events,alias:mshi"`

	MembershipEventID uuid.UUID `bun:"membership_event_id,pk"`
	MemberID          uuid.UUID `bun:"member_id,notnull"`
	EventType         string    `bun:"event_type,notnull"`
	CreatedBy         uuid.UUID `bun:"created_by,notnull"`
	EventAt           time.Time `bun:"event_at,notnull,default:current_timestamp"`

	Member          *Member                `bun:"ms,rel:belongs-to"`
	Creator         *Member                `bun:"ms,rel:belongs-to,join:created_by=member_id"`
	MembershipEvent *LatestMembershipEvent `bun:"lmshi,rel:has-one"`
}

type LatestMembershipEvent struct {
	bun.BaseModel `bun:"table:latest_membership_events,alias:lmshi"`

	MembershipEventID uuid.UUID `bun:"membership_event_id,pk"`
	MemberID          uuid.UUID `bun:"member_id,notnull"`

	MembershipEvent *MembershipEvent `bun:"mshi,rel:has-one"`
	Member          *Member          `bun:"ms,rel:belongs-to"`
}
