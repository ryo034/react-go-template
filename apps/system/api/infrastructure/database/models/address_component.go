package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type AddressComponent struct {
	bun.BaseModel `bun:"table:address_components,alias:ac"`

	ComponentID   uuid.UUID `bun:"component_id,pk"`
	ComponentType string    `bun:"component_type,notnull"`
	ComponentName string    `bun:"component_name,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
}
