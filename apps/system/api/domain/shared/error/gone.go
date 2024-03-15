package error

import "fmt"

type Gone struct {
	target string
}

func NewGone(target string) *Gone {
	return &Gone{target}
}

func (e *Gone) Error() string {
	return fmt.Sprintf("Data already gone: %s", e.target)
}

func (e *Gone) Target() string {
	return e.target
}

func (e *Gone) MessageKey() MessageKey {
	return GoneMessageKey
}
