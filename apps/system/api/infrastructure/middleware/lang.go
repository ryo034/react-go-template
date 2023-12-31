package middleware

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"google.golang.org/grpc"
)

func LangUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(shared.SetLang(ctx), req)
	}
}

func LangStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		w := newStreamContextWrapper(ss)
		w.SetContext(shared.SetLang(w.Context()))
		return handler(srv, w)
	}
}

type StreamContextWrapper interface {
	grpc.ServerStream
	SetContext(context.Context)
}

type wrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrapper) Context() context.Context {
	return w.ctx
}

func (w *wrapper) SetContext(ctx context.Context) {
	w.ctx = ctx
}

func newStreamContextWrapper(inner grpc.ServerStream) StreamContextWrapper {
	ctx := inner.Context()
	return &wrapper{inner, ctx}
}
