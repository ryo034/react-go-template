package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type SystemAccountProfile struct {
	bun.BaseModel `bun:"table:system_account_profiles,alias:saps"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	Name            string    `bun:"name,notnull"`
	Email           string    `bun:"email,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
