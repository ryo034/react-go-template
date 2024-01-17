package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Member struct {
	bun.BaseModel `bun:"table:members,alias:m"`

	MemberID        uuid.UUID `bun:"member_id,pk"`
	SystemAccountID uuid.UUID `bun:"system_account_id,notnull"`
	WorkspaceID     uuid.UUID `bun:"workspace_id,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Workspace     *Workspace     `bun:"rel:belongs-to"`
	SystemAccount *SystemAccount `bun:"rel:belongs-to"`
	Profile       *MemberProfile `bun:"rel:has-one"`
}
