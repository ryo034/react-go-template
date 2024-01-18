package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/multi_factor"
	"github.com/ryo034/react-go-template/apps/system/api/domain/membership"
)

type Me struct {
	emailVerified bool
	multiFactors  multi_factor.MultiFactors
	membership    *membership.Membership
}

func NewMe(emailVerified bool, multiFactors multi_factor.MultiFactors, membership *membership.Membership) *Me {
	return &Me{emailVerified, multiFactors, membership}
}

func (m *Me) EmailVerified() bool {
	return m.emailVerified
}

func (m *Me) NotEmailVerified() bool {
	return !m.emailVerified
}

func (m *Me) MultiFactors() multi_factor.MultiFactors {
	return m.multiFactors
}

func (m *Me) Membership() *membership.Membership {
	return m.membership
}
