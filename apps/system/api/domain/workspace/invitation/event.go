package invitation

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
)

type EventType int

const (
	Verified EventType = iota // 承認済み
	Accepted                  // 受け入れ済み
	Reissued                  // 再発行
	Revoked                   // 取り消し
)

type Event struct {
	eventType  EventType
	occurredAt datetime.Datetime
}

func NewEvent(eventType EventType, occurredAt datetime.Datetime) Event {
	return Event{eventType, occurredAt}

}

func NewAsVerified(occurredAt datetime.Datetime) Event {
	return NewEvent(Verified, occurredAt)
}

func NewAsReissued(occurredAt datetime.Datetime) Event {
	return NewEvent(Reissued, occurredAt)
}

func NewAsRevoked(occurredAt datetime.Datetime) Event {
	return NewEvent(Revoked, occurredAt)
}

func NewAsAccepted(occurredAt datetime.Datetime) Event {
	return NewEvent(Accepted, occurredAt)

}

func (e *Event) EventType() EventType {
	return e.eventType
}

func (e *Event) OccurredAt() datetime.Datetime {
	return e.occurredAt
}

func (e *Event) IsVerified() bool {
	return e.eventType == Verified
}

func (e *Event) IsAccepted() bool {
	return e.eventType == Accepted
}

func (e *Event) IsReissued() bool {
	return e.eventType == Reissued
}

func (e *Event) IsRevoked() bool {
	return e.eventType == Revoked
}

func (e *Event) IsActive() bool {
	return e.eventType != Verified
}
