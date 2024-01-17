package middleware

import (
	"github.com/rs/xid"
)

const RequestIDKey = "request-id"

func GenRequestID() string {
	return xid.New().String()
}
