package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type WorkspaceDetail struct {
	bun.BaseModel `bun:"table:workspace_details,alias:od"`

	WorkspaceID uuid.UUID `bun:"workspace_id,pk"`
	Name        string    `bun:"name,notnull"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
