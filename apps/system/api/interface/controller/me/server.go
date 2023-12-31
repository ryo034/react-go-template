package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
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

func (s *Server) SignUp(ctx context.Context, req *connect.Request[mePb.SignUpRequest]) (*connect.Response[mePb.SignUpResponse], error) {
	meID, err := s.co.GetUID(ctx)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	firstName, err := account.NewFirstName(req.Msg.FirstName)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	lastName, err := account.NewLastName(req.Msg.LastName)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	m, err := s.au.SignUp(ctx, meID, firstName, lastName)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	result, err := s.meResponseAdapter.Adapt(m)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	return connect.NewResponse(&mePb.SignUpResponse{Me: result}), nil
}

func (s *Server) RegisterComplete(ctx context.Context, _ *connect.Request[mePb.RegisterCompleteRequest]) (*connect.Response[mePb.RegisterCompleteResponse], error) {
	meID, err := s.co.GetUID(ctx)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	m, err := s.au.RegisterComplete(ctx, meID)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	result, err := s.meResponseAdapter.Adapt(m)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	return connect.NewResponse(&mePb.RegisterCompleteResponse{Me: result}), nil
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

func (s *Server) UpdateEmail(ctx context.Context, req *connect.Request[mePb.UpdateEmailRequest]) (*connect.Response[mePb.UpdateEmailResponse], error) {
	meID, err := s.co.GetUID(ctx)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	em, err := account.NewEmail(req.Msg.Email)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	m, err := s.au.UpdateEmail(ctx, meID, em)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	result, err := s.meResponseAdapter.Adapt(m)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	return connect.NewResponse(&mePb.UpdateEmailResponse{Me: result}), nil
}

func (s *Server) UpdateName(ctx context.Context, req *connect.Request[mePb.UpdateNameRequest]) (*connect.Response[mePb.UpdateNameResponse], error) {
	meID, err := s.co.GetUID(ctx)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	fn, err := account.NewFirstName(req.Msg.FirstName)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	ln, err := account.NewLastName(req.Msg.LastName)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	m, err := s.au.UpdateName(ctx, meID, fn, ln)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	result, err := s.meResponseAdapter.Adapt(m)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	return connect.NewResponse(&mePb.UpdateNameResponse{Me: result}), nil
}

func (s *Server) UpdatePhoneNumber(ctx context.Context, req *connect.Request[mePb.UpdatePhoneNumberRequest]) (*connect.Response[mePb.UpdatePhoneNumberResponse], error) {
	meID, err := s.co.GetUID(ctx)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	ph, err := phone.NewPhoneNumber(req.Msg.PhoneNumber)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	m, err := s.au.UpdatePhoneNumber(ctx, meID, ph)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	result, err := s.meResponseAdapter.Adapt(m)
	if err != nil {
		return nil, s.sharedResResolver.Error(ctx, err)
	}
	return connect.NewResponse(&mePb.UpdatePhoneNumberResponse{Me: result}), nil
}
