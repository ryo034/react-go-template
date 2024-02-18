package invitation

import (
	"sort"
)

type Events interface {
	Size() int
	AsSlice() []Event
	IsEmpty() bool
	IsNotEmpty() bool
	Latest() *Event
}

type events struct {
	wrapped []Event
}

func NewEvents(wrapped []Event) Events {
	return &events{wrapped}
}

func (ms *events) IsEmpty() bool {
	return len(ms.wrapped) == 0
}

func (ms *events) IsNotEmpty() bool {
	return !ms.IsEmpty()
}

func (ms *events) Size() int {
	return len(ms.wrapped)
}

func (ms *events) AsSlice() []Event {
	return append(make([]Event, 0, ms.Size()), ms.wrapped...)
}

func (ms *events) Latest() *Event {
	//OccurredAtが最新のもの,sortを使う
	if ms.IsEmpty() {
		return nil
	}
	tmp := ms.AsSlice()
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].OccurredAt().ToTime().After(tmp[j].OccurredAt().ToTime())
	})
	return &tmp[0]
}
