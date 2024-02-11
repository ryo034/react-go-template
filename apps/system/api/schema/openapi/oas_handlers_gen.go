// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleAPIV1AuthOAuthPostRequest handles POST /api/v1/auth/oauth operation.
//
// Auth by OAuth.
//
// POST /api/v1/auth/oauth
func (s *Server) handleAPIV1AuthOAuthPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/auth/oauth"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1AuthOAuthPost",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIV1AuthOAuthPost",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "APIV1AuthOAuthPost", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				recordError("Security:Bearer", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response APIV1AuthOAuthPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1AuthOAuthPost",
			OperationSummary: "Auth by OAuth",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = APIV1AuthOAuthPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1AuthOAuthPost(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1AuthOAuthPost(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1AuthOAuthPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAPIV1AuthOtpPostRequest handles POST /api/v1/auth/otp operation.
//
// One Time Password (OTP) to user.
//
// POST /api/v1/auth/otp
func (s *Server) handleAPIV1AuthOtpPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/auth/otp"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1AuthOtpPost",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIV1AuthOtpPost",
			ID:   "",
		}
	)
	request, close, err := s.decodeAPIV1AuthOtpPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1AuthOtpPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1AuthOtpPost",
			OperationSummary: "Send OTP",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *APIV1AuthOtpPostReq
			Params   = struct{}
			Response = APIV1AuthOtpPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1AuthOtpPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1AuthOtpPost(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1AuthOtpPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAPIV1AuthOtpVerifyPostRequest handles POST /api/v1/auth/otp/verify operation.
//
// Verify OTP sent by user.
//
// POST /api/v1/auth/otp/verify
func (s *Server) handleAPIV1AuthOtpVerifyPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/auth/otp/verify"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1AuthOtpVerifyPost",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIV1AuthOtpVerifyPost",
			ID:   "",
		}
	)
	request, close, err := s.decodeAPIV1AuthOtpVerifyPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1AuthOtpVerifyPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1AuthOtpVerifyPost",
			OperationSummary: "Verify OTP",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *APIV1AuthOtpVerifyPostReq
			Params   = struct{}
			Response = APIV1AuthOtpVerifyPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1AuthOtpVerifyPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1AuthOtpVerifyPost(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1AuthOtpVerifyPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAPIV1MeGetRequest handles GET /api/v1/me operation.
//
// Returns the admin user.
//
// GET /api/v1/me
func (s *Server) handleAPIV1MeGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/me"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1MeGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIV1MeGet",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "APIV1MeGet", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				recordError("Security:Bearer", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response APIV1MeGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1MeGet",
			OperationSummary: "Get Admin User",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = APIV1MeGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1MeGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1MeGet(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1MeGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAPIV1MeProfilePutRequest handles PUT /api/v1/me/profile operation.
//
// Updates the user profile.
//
// PUT /api/v1/me/profile
func (s *Server) handleAPIV1MeProfilePutRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/api/v1/me/profile"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1MeProfilePut",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIV1MeProfilePut",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "APIV1MeProfilePut", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				recordError("Security:Bearer", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeAPIV1MeProfilePutRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1MeProfilePutRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1MeProfilePut",
			OperationSummary: "Update Profile",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *APIV1MeProfilePutReq
			Params   = struct{}
			Response = APIV1MeProfilePutRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1MeProfilePut(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1MeProfilePut(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1MeProfilePutResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAPIV1MembersGetRequest handles GET /api/v1/members operation.
//
// Returns the members of the workspace.
//
// GET /api/v1/members
func (s *Server) handleAPIV1MembersGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/members"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1MembersGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIV1MembersGet",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "APIV1MembersGet", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				recordError("Security:Bearer", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response APIV1MembersGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1MembersGet",
			OperationSummary: "Get Members",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = APIV1MembersGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1MembersGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1MembersGet(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1MembersGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAPIV1PingGetRequest handles GET /api/v1/ping operation.
//
// Checks if the server is running.
//
// GET /api/v1/ping
func (s *Server) handleAPIV1PingGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/ping"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1PingGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err error
	)

	var response APIV1PingGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1PingGet",
			OperationSummary: "Checks if the server is running",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = APIV1PingGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1PingGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1PingGet(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1PingGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAPIV1WorkspacesGetRequest handles GET /api/v1/workspaces operation.
//
// Returns the workspaces the user is a member of.
//
// GET /api/v1/workspaces
func (s *Server) handleAPIV1WorkspacesGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/workspaces"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1WorkspacesGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIV1WorkspacesGet",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "APIV1WorkspacesGet", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				recordError("Security:Bearer", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response APIV1WorkspacesGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1WorkspacesGet",
			OperationSummary: "Get Joined Workspaces",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = APIV1WorkspacesGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1WorkspacesGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1WorkspacesGet(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1WorkspacesGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAPIV1WorkspacesPostRequest handles POST /api/v1/workspaces operation.
//
// Creates a new workspace.
//
// POST /api/v1/workspaces
func (s *Server) handleAPIV1WorkspacesPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/workspaces"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIV1WorkspacesPost",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIV1WorkspacesPost",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "APIV1WorkspacesPost", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				recordError("Security:Bearer", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeAPIV1WorkspacesPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1WorkspacesPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "APIV1WorkspacesPost",
			OperationSummary: "Create Workspace",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *APIV1WorkspacesPostReq
			Params   = struct{}
			Response = APIV1WorkspacesPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1WorkspacesPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1WorkspacesPost(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1WorkspacesPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInviteMultipleUsersToWorkspaceRequest handles inviteMultipleUsersToWorkspace operation.
//
// Invite multiple users to the workspace by email.
//
// POST /api/v1/members/invitations/bulk
func (s *Server) handleInviteMultipleUsersToWorkspaceRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteMultipleUsersToWorkspace"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/members/invitations/bulk"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InviteMultipleUsersToWorkspace",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "InviteMultipleUsersToWorkspace",
			ID:   "inviteMultipleUsersToWorkspace",
		}
	)
	request, close, err := s.decodeInviteMultipleUsersToWorkspaceRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response InviteMultipleUsersToWorkspaceRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InviteMultipleUsersToWorkspace",
			OperationSummary: "Invite multiple users to the workspace by email",
			OperationID:      "inviteMultipleUsersToWorkspace",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *InviteMultipleUsersToWorkspaceReq
			Params   = struct{}
			Response = InviteMultipleUsersToWorkspaceRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.InviteMultipleUsersToWorkspace(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.InviteMultipleUsersToWorkspace(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInviteMultipleUsersToWorkspaceResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleLoginRequest handles login operation.
//
// Login.
//
// POST /api/v1/login
func (s *Server) handleLoginRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("login"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/login"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Login",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err error
	)

	var response LoginRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "Login",
			OperationSummary: "Login",
			OperationID:      "login",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = LoginRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Login(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.Login(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeLoginResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleVerifyInvitationRequest handles verifyInvitation operation.
//
// Verify invitation.
//
// POST /api/v1/members/invitations/verify
func (s *Server) handleVerifyInvitationRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("verifyInvitation"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/members/invitations/verify"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "VerifyInvitation",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "VerifyInvitation",
			ID:   "verifyInvitation",
		}
	)
	request, close, err := s.decodeVerifyInvitationRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response VerifyInvitationRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "VerifyInvitation",
			OperationSummary: "Verify invitation",
			OperationID:      "verifyInvitation",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *VerifyInvitationReq
			Params   = struct{}
			Response = VerifyInvitationRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.VerifyInvitation(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.VerifyInvitation(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeVerifyInvitationResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
