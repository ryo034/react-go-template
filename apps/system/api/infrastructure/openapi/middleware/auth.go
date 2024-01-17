package middleware

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Middleware struct {
}

func (m *Middleware) HandleBearer(ctx context.Context, operationName string, t openapi.Bearer) (context.Context, error) {
	return ctx, nil
}

func NewSecMiddleware() *Middleware {
	return &Middleware{}
}
