package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Workspace struct {
	bun.BaseModel `bun:"table:workspaces,alias:org"`

	WorkspaceID uuid.UUID `bun:"workspace_id,pk"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Detail  *WorkspaceDetail `bun:"rel:has-one"`
	Members []*Member        `bun:"rel:has-many"`
}
