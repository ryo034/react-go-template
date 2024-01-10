package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Employee struct {
	bun.BaseModel `bun:"table:employees,alias:e"`

	EmployeeID      uuid.UUID `bun:"employee_id,pk"`
	SystemAccountID uuid.UUID `bun:"system_account_id,notnull"`
	OrganizationID  uuid.UUID `bun:"organization_id,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Organization  *Organization         `bun:"rel:belongs-to"`
	SystemAccount *SystemAccount        `bun:"rel:belongs-to"`
	Profile       *EmployeeProfile      `bun:"rel:has-one"`
	Hires         []*EmployeeHire       `bun:"rel:has-many"`
	Separations   []*EmployeeSeparation `bun:"rel:has-many"`
}
