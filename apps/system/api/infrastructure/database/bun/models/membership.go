package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Membership struct {
	bun.BaseModel `bun:"table:memberships,alias:mship"`

	MemberID    uuid.UUID `bun:"member_id,pk"`
	WorkspaceID uuid.UUID `bun:"workspace_id,notnull"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Workspace *Workspace `bun:"rel:belongs-to"`
	Member    *Member    `bun:"rel:belongs-to"`
}
