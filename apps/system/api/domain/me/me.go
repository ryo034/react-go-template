package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/multi_factor"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

type Me struct {
	multiFactors multi_factor.MultiFactors
	member       *member.Member
}

func NewMe(multiFactors multi_factor.MultiFactors, member *member.Member) *Me {
	return &Me{multiFactors, member}
}

func (m *Me) MultiFactors() multi_factor.MultiFactors {
	return m.multiFactors
}

func (m *Me) Member() *member.Member {
	return m.member
}
