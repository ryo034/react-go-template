package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

const (
	InvitationUnitTableName       = "invitation_units"
	InvitationUnitTableAliasName  = "invus"
	InvitationTableName           = "invitations"
	InvitationTableAliasName      = "invs"
	InviteeNameTableName          = "invitee_names"
	InviteeNameTableAliasName     = "invns"
	InviteeTableName              = "invitees"
	InvitationEventTableName      = "invitation_events"
	InvitationEventTableAliasName = "inves"
	InvitationTokenTableName      = "invitation_tokens"
	InvitationTokenTableAliasName = "invts"
)

type InvitationUnit struct {
	bun.BaseModel `bun:"table:invitation_units,alias:invus"`

	InvitationUnitID uuid.UUID `bun:"invitation_unit_id,pk"`
	WorkspaceID      uuid.UUID `bun:"workspace_id,notnull"`
	InvitedBy        uuid.UUID `bun:"invited_by,notnull"`
	CreatedAt        time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Workspace   *Workspace    `bun:"ws,rel:belongs-to"`
	Member      *Member       `bun:"ms,rel:belongs-to,join:invited_by=member_id"`
	Invitations []*Invitation `bun:"invs,rel:has-many"`
}

type Invitation struct {
	bun.BaseModel `bun:"table:invitations,alias:invs"`

	InvitationID     uuid.UUID `bun:"invitation_id,pk"`
	InvitationUnitID uuid.UUID `bun:"invitation_unit_id,notnull"`

	InvitationUnit *InvitationUnit    `bun:"invus,rel:belongs-to"`
	InviteeName    *InviteeName       `bun:"invns,rel:has-one"`
	Invitee        *Invitee           `bun:"rel:has-one"`
	Events         []*InvitationEvent `bun:"inves,rel:has-many"`
	Tokens         []*InvitationToken `bun:"invts,rel:has-many"`
}

type Invitations []*Invitation

type InviteeName struct {
	bun.BaseModel `bun:"table:invitee_names,alias:invns"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	DisplayName  string    `bun:"display_name,notnull"`

	Invitation *Invitation `bun:"rel:belongs-to"`
}

type Invitee struct {
	bun.BaseModel `bun:"table:invitees"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	Email        string    `bun:"email,notnull"`

	Invitation *Invitation `bun:"invs,rel:belongs-to"`
}

type InvitationEvent struct {
	bun.BaseModel `bun:"table:invitation_events,alias:inves"`

	InvitationEventID uuid.UUID `bun:"invitation_event_id,pk"`
	InvitationID      uuid.UUID `bun:"invitation_id,notnull"`
	EventType         string    `bun:"event_type,notnull"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Invitation *Invitation `bun:"invs,rel:belongs-to"`
}

type InvitationToken struct {
	bun.BaseModel `bun:"table:invitation_tokens,alias:invts"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	Token        uuid.UUID `bun:"token,pk"`
	ExpiredAt    time.Time `bun:"expired_at,notnull"`
	CreatedAt    time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Invitation *Invitation `bun:"invs,rel:belongs-to"`
}
