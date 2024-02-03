package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type MemberLoginHistory struct {
	bun.BaseModel `bun:"table:member_login_histories,alias:llw"`

	MemberLoginHistoryID uuid.UUID `bun:"member_login_history_id,pk"`
	MemberID             uuid.UUID `bun:"member_id,notnull"`
	LoginAt              time.Time `bun:"login_at,notnull,default:current_timestamp"`

	Member *Member `bun:"rel:belongs-to"`
}
