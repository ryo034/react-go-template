package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

const (
	InvitationUnitTableName  = "invitation_units"
	InvitationTableName      = "invitations"
	InviteeNameTableName     = "invitee_names"
	InviteeTableName         = "invitees"
	InvitationEventTableName = "invitations_events"
	InvitationTokenTableName = "invitation_tokens"
)

type InvitationUnit struct {
	bun.BaseModel `bun:"table:invitation_units"`

	InvitationUnitID uuid.UUID `bun:"invitation_unit_id,pk"`
	WorkspaceID      uuid.UUID `bun:"workspace_id,notnull"`
	InvitedBy        uuid.UUID `bun:"invited_by,notnull"`
	CreatedAt        time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Workspace   *Workspace    `bun:"rel:belongs-to"`
	Member      *Member       `bun:"rel:belongs-to,join:invited_by=member_id"`
	Invitations []*Invitation `bun:"rel:has-many"`
}

type Invitation struct {
	bun.BaseModel `bun:"table:invitations"`

	InvitationID     uuid.UUID `bun:"invitation_id,pk"`
	InvitationUnitID uuid.UUID `bun:"invitation_unit_id,notnull"`

	InvitationUnit *InvitationUnit    `bun:"rel:belongs-to"`
	InviteeName    *InviteeName       `bun:"rel:has-one"`
	Invitee        *Invitee           `bun:"rel:has-one"`
	Events         []*InvitationEvent `bun:"rel:has-many"`
	Tokens         []*InvitationToken `bun:"rel:has-many"`
}

type Invitations []*Invitation

type InviteeName struct {
	bun.BaseModel `bun:"table:invitee_names"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	DisplayName  string    `bun:"display_name,notnull"`

	Invitation *Invitation `bun:"rel:belongs-to"`
}

type Invitee struct {
	bun.BaseModel `bun:"table:invitees"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	Email        string    `bun:"email,notnull"`

	Invitation *Invitation `bun:"rel:belongs-to"`
}

type InvitationEvent struct {
	bun.BaseModel `bun:"table:invitations_events"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	EventType    string    `bun:"event_type,notnull"`
	CreatedAt    time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Invitation *Invitation `bun:"rel:belongs-to"`
}

type InvitationToken struct {
	bun.BaseModel `bun:"table:invitation_tokens"`

	InvitationID uuid.UUID `bun:"invitation_id,pk"`
	Token        uuid.UUID `bun:"token,notnull"`
	ExpiredAt    time.Time `bun:"expired_at,notnull"`
	CreatedAt    time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Invitation *Invitation `bun:"rel:belongs-to"`
}
