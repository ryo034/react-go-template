package error

import "fmt"

type NoSuchData struct {
	target string
}

func NewNoSuchData(target string) *NoSuchData {
	return &NoSuchData{target}
}

func (n *NoSuchData) Error() string {
	return fmt.Sprintf("no such data:%s", n.Target())
}

func (n *NoSuchData) Target() string {
	return n.target
}

func (n *NoSuchData) MessageKey() MessageKey {
	return NoSuchDataMessageKey
}

func (n *NoSuchData) Code() string {
	return "404-000"
}
