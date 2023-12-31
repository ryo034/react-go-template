package health

import (
	"context"
	"github.com/bufbuild/connect-go"
	health "github.com/ryo034/react-go-template/apps/system/api/schema/pb/health/v1"
	healthconnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/health/v1/v1connect"
)

type Server struct {
	healthconnect.UnimplementedHealthServiceHandler
}

func NewServer() healthconnect.HealthServiceHandler {
	return &Server{
		healthconnect.UnimplementedHealthServiceHandler{},
	}
}

// AuthFuncOverride SkipAuthHealthServer構造体がServiceAuthFuncOverrideインターフェースを実装する
func (*Server) AuthFuncOverride(ctx context.Context, _ string) (context.Context, error) {
	return ctx, nil
}

func (*Server) Check(context.Context, *connect.Request[health.CheckRequest]) (*connect.Response[health.CheckResponse], error) {
	return connect.NewResponse(&health.CheckResponse{
		Status: health.ServingStatus_SERVING_STATUS_SERVING,
	}), nil
}
