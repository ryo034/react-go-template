package error

import "fmt"

type ConflictVersion struct {
	v              uint16
	currentVersion uint16
}

func NewConflictVersion(v uint16, currentVersion uint16) *ConflictVersion {
	return &ConflictVersion{v, currentVersion}
}

func (c *ConflictVersion) Version() uint16 {
	return c.v
}
func (c *ConflictVersion) CurrentVersion() uint16 {
	return c.currentVersion
}

func (c *ConflictVersion) Error() string {
	return fmt.Sprintf("version:%d is conflicted with current version:%d", c.Version(), c.CurrentVersion())
}

func (c *ConflictVersion) MessageKey() MessageKey {
	return ConflictVersionMessageKey
}

func (c *ConflictVersion) Code() string {
	return "400-" + string(ConflictVersionCodeKey)
}
