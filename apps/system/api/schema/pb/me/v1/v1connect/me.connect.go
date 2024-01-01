// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: me/v1/me.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ryo034/react-go-template/apps/system/api/schema/pb/me/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// MeServiceName is the fully-qualified name of the MeService service.
	MeServiceName = "me.v1.MeService"
)

// MeServiceClient is a client for the me.v1.MeService service.
type MeServiceClient interface {
	Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.LoginResponse], error)
	SignUp(context.Context, *connect_go.Request[v1.SignUpRequest]) (*connect_go.Response[v1.SignUpResponse], error)
	RegisterComplete(context.Context, *connect_go.Request[v1.RegisterCompleteRequest]) (*connect_go.Response[v1.RegisterCompleteResponse], error)
	Find(context.Context, *connect_go.Request[v1.FindRequest]) (*connect_go.Response[v1.FindResponse], error)
	UpdateName(context.Context, *connect_go.Request[v1.UpdateNameRequest]) (*connect_go.Response[v1.UpdateNameResponse], error)
	UpdateEmail(context.Context, *connect_go.Request[v1.UpdateEmailRequest]) (*connect_go.Response[v1.UpdateEmailResponse], error)
	UpdatePhoneNumber(context.Context, *connect_go.Request[v1.UpdatePhoneNumberRequest]) (*connect_go.Response[v1.UpdatePhoneNumberResponse], error)
}

// NewMeServiceClient constructs a client for the me.v1.MeService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMeServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &meServiceClient{
		login: connect_go.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+"/me.v1.MeService/Login",
			opts...,
		),
		signUp: connect_go.NewClient[v1.SignUpRequest, v1.SignUpResponse](
			httpClient,
			baseURL+"/me.v1.MeService/SignUp",
			opts...,
		),
		registerComplete: connect_go.NewClient[v1.RegisterCompleteRequest, v1.RegisterCompleteResponse](
			httpClient,
			baseURL+"/me.v1.MeService/RegisterComplete",
			opts...,
		),
		find: connect_go.NewClient[v1.FindRequest, v1.FindResponse](
			httpClient,
			baseURL+"/me.v1.MeService/Find",
			opts...,
		),
		updateName: connect_go.NewClient[v1.UpdateNameRequest, v1.UpdateNameResponse](
			httpClient,
			baseURL+"/me.v1.MeService/UpdateName",
			opts...,
		),
		updateEmail: connect_go.NewClient[v1.UpdateEmailRequest, v1.UpdateEmailResponse](
			httpClient,
			baseURL+"/me.v1.MeService/UpdateEmail",
			opts...,
		),
		updatePhoneNumber: connect_go.NewClient[v1.UpdatePhoneNumberRequest, v1.UpdatePhoneNumberResponse](
			httpClient,
			baseURL+"/me.v1.MeService/UpdatePhoneNumber",
			opts...,
		),
	}
}

// meServiceClient implements MeServiceClient.
type meServiceClient struct {
	login             *connect_go.Client[v1.LoginRequest, v1.LoginResponse]
	signUp            *connect_go.Client[v1.SignUpRequest, v1.SignUpResponse]
	registerComplete  *connect_go.Client[v1.RegisterCompleteRequest, v1.RegisterCompleteResponse]
	find              *connect_go.Client[v1.FindRequest, v1.FindResponse]
	updateName        *connect_go.Client[v1.UpdateNameRequest, v1.UpdateNameResponse]
	updateEmail       *connect_go.Client[v1.UpdateEmailRequest, v1.UpdateEmailResponse]
	updatePhoneNumber *connect_go.Client[v1.UpdatePhoneNumberRequest, v1.UpdatePhoneNumberResponse]
}

// Login calls me.v1.MeService.Login.
func (c *meServiceClient) Login(ctx context.Context, req *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// SignUp calls me.v1.MeService.SignUp.
func (c *meServiceClient) SignUp(ctx context.Context, req *connect_go.Request[v1.SignUpRequest]) (*connect_go.Response[v1.SignUpResponse], error) {
	return c.signUp.CallUnary(ctx, req)
}

// RegisterComplete calls me.v1.MeService.RegisterComplete.
func (c *meServiceClient) RegisterComplete(ctx context.Context, req *connect_go.Request[v1.RegisterCompleteRequest]) (*connect_go.Response[v1.RegisterCompleteResponse], error) {
	return c.registerComplete.CallUnary(ctx, req)
}

// Find calls me.v1.MeService.Find.
func (c *meServiceClient) Find(ctx context.Context, req *connect_go.Request[v1.FindRequest]) (*connect_go.Response[v1.FindResponse], error) {
	return c.find.CallUnary(ctx, req)
}

// UpdateName calls me.v1.MeService.UpdateName.
func (c *meServiceClient) UpdateName(ctx context.Context, req *connect_go.Request[v1.UpdateNameRequest]) (*connect_go.Response[v1.UpdateNameResponse], error) {
	return c.updateName.CallUnary(ctx, req)
}

// UpdateEmail calls me.v1.MeService.UpdateEmail.
func (c *meServiceClient) UpdateEmail(ctx context.Context, req *connect_go.Request[v1.UpdateEmailRequest]) (*connect_go.Response[v1.UpdateEmailResponse], error) {
	return c.updateEmail.CallUnary(ctx, req)
}

// UpdatePhoneNumber calls me.v1.MeService.UpdatePhoneNumber.
func (c *meServiceClient) UpdatePhoneNumber(ctx context.Context, req *connect_go.Request[v1.UpdatePhoneNumberRequest]) (*connect_go.Response[v1.UpdatePhoneNumberResponse], error) {
	return c.updatePhoneNumber.CallUnary(ctx, req)
}

// MeServiceHandler is an implementation of the me.v1.MeService service.
type MeServiceHandler interface {
	Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.LoginResponse], error)
	SignUp(context.Context, *connect_go.Request[v1.SignUpRequest]) (*connect_go.Response[v1.SignUpResponse], error)
	RegisterComplete(context.Context, *connect_go.Request[v1.RegisterCompleteRequest]) (*connect_go.Response[v1.RegisterCompleteResponse], error)
	Find(context.Context, *connect_go.Request[v1.FindRequest]) (*connect_go.Response[v1.FindResponse], error)
	UpdateName(context.Context, *connect_go.Request[v1.UpdateNameRequest]) (*connect_go.Response[v1.UpdateNameResponse], error)
	UpdateEmail(context.Context, *connect_go.Request[v1.UpdateEmailRequest]) (*connect_go.Response[v1.UpdateEmailResponse], error)
	UpdatePhoneNumber(context.Context, *connect_go.Request[v1.UpdatePhoneNumberRequest]) (*connect_go.Response[v1.UpdatePhoneNumberResponse], error)
}

// NewMeServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMeServiceHandler(svc MeServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/me.v1.MeService/Login", connect_go.NewUnaryHandler(
		"/me.v1.MeService/Login",
		svc.Login,
		opts...,
	))
	mux.Handle("/me.v1.MeService/SignUp", connect_go.NewUnaryHandler(
		"/me.v1.MeService/SignUp",
		svc.SignUp,
		opts...,
	))
	mux.Handle("/me.v1.MeService/RegisterComplete", connect_go.NewUnaryHandler(
		"/me.v1.MeService/RegisterComplete",
		svc.RegisterComplete,
		opts...,
	))
	mux.Handle("/me.v1.MeService/Find", connect_go.NewUnaryHandler(
		"/me.v1.MeService/Find",
		svc.Find,
		opts...,
	))
	mux.Handle("/me.v1.MeService/UpdateName", connect_go.NewUnaryHandler(
		"/me.v1.MeService/UpdateName",
		svc.UpdateName,
		opts...,
	))
	mux.Handle("/me.v1.MeService/UpdateEmail", connect_go.NewUnaryHandler(
		"/me.v1.MeService/UpdateEmail",
		svc.UpdateEmail,
		opts...,
	))
	mux.Handle("/me.v1.MeService/UpdatePhoneNumber", connect_go.NewUnaryHandler(
		"/me.v1.MeService/UpdatePhoneNumber",
		svc.UpdatePhoneNumber,
		opts...,
	))
	return "/me.v1.MeService/", mux
}

// UnimplementedMeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMeServiceHandler struct{}

func (UnimplementedMeServiceHandler) Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.LoginResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("me.v1.MeService.Login is not implemented"))
}

func (UnimplementedMeServiceHandler) SignUp(context.Context, *connect_go.Request[v1.SignUpRequest]) (*connect_go.Response[v1.SignUpResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("me.v1.MeService.SignUp is not implemented"))
}

func (UnimplementedMeServiceHandler) RegisterComplete(context.Context, *connect_go.Request[v1.RegisterCompleteRequest]) (*connect_go.Response[v1.RegisterCompleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("me.v1.MeService.RegisterComplete is not implemented"))
}

func (UnimplementedMeServiceHandler) Find(context.Context, *connect_go.Request[v1.FindRequest]) (*connect_go.Response[v1.FindResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("me.v1.MeService.Find is not implemented"))
}

func (UnimplementedMeServiceHandler) UpdateName(context.Context, *connect_go.Request[v1.UpdateNameRequest]) (*connect_go.Response[v1.UpdateNameResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("me.v1.MeService.UpdateName is not implemented"))
}

func (UnimplementedMeServiceHandler) UpdateEmail(context.Context, *connect_go.Request[v1.UpdateEmailRequest]) (*connect_go.Response[v1.UpdateEmailResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("me.v1.MeService.UpdateEmail is not implemented"))
}

func (UnimplementedMeServiceHandler) UpdatePhoneNumber(context.Context, *connect_go.Request[v1.UpdatePhoneNumberRequest]) (*connect_go.Response[v1.UpdatePhoneNumberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("me.v1.MeService.UpdatePhoneNumber is not implemented"))
}