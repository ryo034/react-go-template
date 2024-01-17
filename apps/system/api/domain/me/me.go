package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/multi_factor"
	"github.com/ryo034/react-go-template/apps/system/api/domain/member"
)

type Me struct {
	emailVerified bool
	multiFactors  multi_factor.MultiFactors
	e             *member.Member
}

func NewMe(emailVerified bool, multiFactors multi_factor.MultiFactors, e *member.Member) *Me {
	return &Me{emailVerified, multiFactors, e}
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

func (m *Me) Member() *member.Member {
	return m.e
}
