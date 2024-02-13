package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type InvitedMember struct {
	bun.BaseModel `bun:"table:invited_members,alias:ims"`

	InvitedMemberID uuid.UUID `bun:"invited_member_id,pk"`
	WorkspaceID     uuid.UUID `bun:"workspace_id,notnull"`
	Email           string    `bun:"email,notnull"`
	DisplayName     string    `bun:"display_name,nullzero"`
	Used            bool      `bun:"used,notnull,default:false"`
	Token           uuid.UUID `bun:"token,notnull"`
	ExpiredAt       time.Time `bun:"expired_at,notnull"`
	InvitedBy       uuid.UUID `bun:"invited_by,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`

	Workspace *Workspace `bun:"rel:belongs-to"`
	Member    *Member    `bun:"rel:belongs-to,join:invited_by=member_id"`
}

type InvitedMembers []*InvitedMember
