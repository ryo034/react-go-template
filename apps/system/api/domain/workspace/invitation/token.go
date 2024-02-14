package invitation

import (
	"github.com/google/uuid"
)

type Token struct {
	uuid.UUID
}

func NewToken(v uuid.UUID) Token {
	return Token{v}
}

func GenerateToken() Token {
	return NewToken(uuid.New())
}

func (t Token) Value() uuid.UUID {
	return t.UUID
}

func (t Token) Equals(v Token) bool {
	return t.UUID == v.UUID
}

func (t Token) NotEquals(v Token) bool {
	return !t.Equals(v)
}
