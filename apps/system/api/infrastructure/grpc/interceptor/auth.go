package interceptor

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
)

func NewAuthInterceptor(ma *middleware.Authentication) connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			newCtx, err := ma.Authenticate(ctx, req)
			if err != nil {
				return nil, err
			}
			return next(newCtx, req)
		}
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}
