package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/multi_factor"
)

type Me struct {
	emailVerified bool
	multiFactors  multi_factor.MultiFactors
}

func NewMe(emailVerified bool, multiFactors multi_factor.MultiFactors) *Me {
	return &Me{emailVerified, multiFactors}
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
