package me

import (
	"context"
	"github.com/bufbuild/connect-go"
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/grpc/response"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	meRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me/request"
	meResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me/response"
	mePb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/me/v1"
	meconnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/me/v1/v1connect"
	meUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Server struct {
	meconnect.UnimplementedMeServiceHandler
	co                shared.ContextOperator
	fd                firebaseDriver.Driver
	au                meUseCase.UseCase
	sharedResResolver response.Resolver
	meRequestAdapter  meRequest.Adapter
	meResponseAdapter meResponse.Adapter
}

func NewServer(
	co shared.ContextOperator,
	fd firebaseDriver.Driver,
	au meUseCase.UseCase,
	resolver response.Resolver,
	meRequestAdapter meRequest.Adapter,
	meResponseAdapter meResponse.Adapter,
) meconnect.MeServiceHandler {
	return &Server{
		meconnect.UnimplementedMeServiceHandler{},
		co,
		fd,
		au,
		resolver,
		meRequestAdapter,
		meResponseAdapter,
	}
}

func (s *Server) Login(ctx context.Context, _ *connect.Request[mePb.LoginRequest]) (*connect.Response[mePb.LoginResponse], error) {
	meID, err := s.co.GetUID(ctx)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	m, err := s.au.Find(ctx, meID, true)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	result, err := s.meResponseAdapter.Adapt(m)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	return connect.NewResponse(&mePb.LoginResponse{Me: result}), nil
}

func (s *Server) Find(ctx context.Context, _ *connect.Request[mePb.FindRequest]) (*connect.Response[mePb.FindResponse], error) {
	meID, err := s.co.GetUID(ctx)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	m, err := s.au.Find(ctx, meID, true)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	result, err := s.meResponseAdapter.Adapt(m)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	return connect.NewResponse(&mePb.FindResponse{Me: result}), nil
}
