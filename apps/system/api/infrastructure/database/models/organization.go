package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Organization struct {
	bun.BaseModel `bun:"table:organizations,alias:org"`

	OrganizationID uuid.UUID `bun:"organization_id,pk"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Detail    *OrganizationDetail `bun:"rel:has-one"`
	Employees []*Employee         `bun:"rel:has-many"`
}
