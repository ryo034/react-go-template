package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type SystemAccountPhoneNumber struct {
	bun.BaseModel `bun:"table:system_account_phone_numbers,alias:sapn"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	PhoneNumber     string    `bun:"phone_number,notnull,unique"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
