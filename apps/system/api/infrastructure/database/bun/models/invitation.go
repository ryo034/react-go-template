package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

const (
	InvitationUnitTableAliasName = "invus"
	InvitationTableAliasName     = "invs"
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

	InvitationUnit *InvitationUnit        `bun:"invus,rel:belongs-to"`
	InviteeName    *InviteeName           `bun:"invns,rel:has-one"`
	Invitee        *Invitee               `bun:"invte,rel:has-one"`
	Event          *LatestInvitationEvent `bun:"linve,rel:has-one"`
	Token          *LatestInvitationToken `bun:"linvt,rel:has-one"`
}

type Invitations []*Invitation

type InviteeName struct {
	bun.BaseModel `bun:"table:invitee_names,alias:invns"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	DisplayName  string    `bun:"display_name,notnull"`

	Invitation *Invitation `bun:"rel:belongs-to"`
}

type Invitee struct {
	bun.BaseModel `bun:"table:invitees,alias:invte"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	Email        string    `bun:"email,notnull"`

	Invitation *Invitation `bun:"invs,rel:belongs-to"`
}

type InvitationEvent struct {
	bun.BaseModel `bun:"table:invitation_events,alias:inve"`

	InvitationEventID uuid.UUID `bun:"invitation_event_id,pk"`
	InvitationID      uuid.UUID `bun:"invitation_id,notnull"`
	EventType         string    `bun:"event_type,notnull"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Invitation            *Invitation            `bun:"invs,rel:belongs-to"`
	LatestInvitationEvent *LatestInvitationEvent `bun:"linve,rel:has-one"`
}

type LatestInvitationEvent struct {
	bun.BaseModel `bun:"table:latest_invitation_events,alias:linve"`

	InvitationEventID uuid.UUID `bun:"invitation_event_id,pk"`
	InvitationID      uuid.UUID `bun:"invitation_id,notnull"`

	Invitation      *Invitation      `bun:"invs,rel:belongs-to"`
	InvitationEvent *InvitationEvent `bun:"inve,rel:has-one"`
}

type InvitationToken struct {
	bun.BaseModel `bun:"table:invitation_tokens,alias:invt"`

	InvitationTokenID uuid.UUID `bun:"invitation_token_id,pk"`
	InvitationID      uuid.UUID `bun:"invitation_id,notnull"`
	Token             uuid.UUID `bun:"token,notnull"`
	ExpiredAt         time.Time `bun:"expired_at,notnull"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Invitation *Invitation            `bun:"invs,rel:belongs-to"`
	Latest     *LatestInvitationToken `bun:"linvt,rel:has-one"`
}

type LatestInvitationToken struct {
	bun.BaseModel `bun:"table:latest_invitation_tokens,alias:linvt"`

	InvitationTokenID uuid.UUID `bun:"invitation_token_id,pk"`
	InvitationID      uuid.UUID `bun:"invitation_id,notnull"`

	InvitationToken *InvitationToken `bun:"invt,rel:has-one"`
	Invitation      *Invitation      `bun:"invs,rel:belongs-to"`
}
