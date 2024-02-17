package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

const (
	SystemAccountTableName            = "system_accounts"
	SystemAccountPhoneNumberTableName = "system_account_phone_numbers"
	SystemAccountProfileTableName     = "system_account_profiles"
)

type SystemAccount struct {
	bun.BaseModel `bun:"table:system_accounts,alias:sa"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`

	PhoneNumber *SystemAccountPhoneNumber `bun:"rel:has-one"`
	Profile     *SystemAccountProfile     `bun:"rel:has-one"`
}

type SystemAccountPhoneNumber struct {
	bun.BaseModel `bun:"table:system_account_phone_numbers,alias:sapn"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	PhoneNumber     string    `bun:"phone_number,notnull,unique"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

type SystemAccountProfile struct {
	bun.BaseModel `bun:"table:system_account_profiles,alias:saps"`

	SystemAccountID uuid.UUID `bun:"system_account_id,pk"`
	Name            string    `bun:"name"`
	Email           string    `bun:"email,notnull"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
