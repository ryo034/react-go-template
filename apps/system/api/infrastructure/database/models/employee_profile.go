package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type EmployeeProfile struct {
	bun.BaseModel `bun:"table:employee_profiles,alias:ep"`

	EmployeeID       uuid.UUID `bun:"employee_id,pk"`
	EmployeeIDNumber string    `bun:"employee_id_number,notnull"`
	Name             string    `bun:"name,notnull"`
	CreatedAt        time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt        time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
