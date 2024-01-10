package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type SystemAccount struct {
	bun.BaseModel `bun:"table:system_accounts,alias:sa"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	PhoneNumber *SystemAccountPhoneNumber `bun:"rel:has-one"`
	Profile     *SystemAccountProfile     `bun:"rel:has-one"`
}
