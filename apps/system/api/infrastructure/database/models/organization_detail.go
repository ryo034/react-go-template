package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type OrganizationDetail struct {
	bun.BaseModel `bun:"table:organization_details,alias:od"`

	OrganizationID uuid.UUID `bun:"organization_id,pk"`
	Name           string    `bun:"name,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
