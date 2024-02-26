package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MembershipPeriod struct {
	bun.BaseModel `bun:"table:membership_periods,alias:msp"`

	MemberID    uuid.UUID `bun:"member_id,pk"`
	WorkspaceID uuid.UUID `bun:"workspace_id,pk"`
	StartDate   time.Time `bun:"start_date,notnull"`
	EndDate     time.Time `bun:"end_date,notnull"`
	Status      string    `bun:"status,notnull"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`
}
