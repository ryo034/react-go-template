package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type EmployeeSeparation struct {
	bun.BaseModel `bun:"table:employee_separations,alias:es"`

	SeparationID   uuid.UUID `bun:"separation_id,pk"`
	EmployeeID     uuid.UUID `bun:"employee_id,notnull"`
	OrganizationID uuid.UUID `bun:"organization_id,notnull"`
	SeparationDate time.Time `bun:"separation_date,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
}
