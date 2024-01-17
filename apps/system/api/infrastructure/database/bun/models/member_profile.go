package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type MemberProfile struct {
	bun.BaseModel `bun:"table:member_profiles,alias:mp"`

	MemberID       uuid.UUID `bun:"member_id,pk"`
	MemberIDNumber string    `bun:"member_id_number,notnull"`
	Name           string    `bun:"name,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
