package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type EmployeeHire struct {
	bun.BaseModel `bun:"table:employee_hires,alias:eh"`

	HireID         uuid.UUID `bun:"hire_id,pk"`
	EmployeeID     uuid.UUID `bun:"employee_id,notnull"`
	OrganizationID uuid.UUID `bun:"organization_id,notnull"`
	HireDate       time.Time `bun:"hire_date,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
}
