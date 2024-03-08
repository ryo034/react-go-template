package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Workspace struct {
	bun.BaseModel `bun:"table:workspaces,alias:ws"`

	WorkspaceID uuid.UUID `bun:"workspace_id,pk"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Detail  *WorkspaceLatestDetail `bun:"wld,rel:has-one"`
	Members []*Member              `bun:"ms,rel:has-many"`
}

type Workspaces []*Workspace

type WorkspaceDetail struct {
	bun.BaseModel `bun:"table:workspace_details,alias:wd"`

	WorkspaceDetailID uuid.UUID `bun:"workspace_detail_id,pk"`
	WorkspaceID       uuid.UUID `bun:"workspace_id,notnull,unique"`
	Name              string    `bun:"name,notnull"`
	Subdomain         string    `bun:"subdomain,notnull,unique"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Workspace       *Workspace             `bun:"ws,rel:belongs-to"`
	WorkspaceDetail *WorkspaceLatestDetail `bun:"wld,rel:has-one"`
}

type WorkspaceLatestDetail struct {
	bun.BaseModel `bun:"table:workspace_latest_details,alias:wld"`

	WorkspaceDetailID uuid.UUID `bun:"workspace_detail_id,pk"`
	WorkspaceID       uuid.UUID `bun:"workspace_id,notnull,unique"`

	WorkspaceDetail *WorkspaceDetail `bun:"wd,rel:belongs-to"`
	Workspace       *Workspace       `bun:"ws,rel:belongs-to"`
}
