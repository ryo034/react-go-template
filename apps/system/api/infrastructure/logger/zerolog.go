package logger

import (
	"context"
	middleware2 "github.com/ogen-go/ogen/middleware"
	"github.com/rs/zerolog"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"net/http"
	"os"
	"time"
)

type zeroLogger struct {
	logger zerolog.Logger
	conf   Config
}

func NewZeroLogger(conf Config, isLocal bool, serviceName string) Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    !isLocal,
	}
	if isLocal {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	logger := zerolog.New(output).With().Timestamp().Logger()
	logger.Info().Str("service", serviceName).Msg("User service started")
	return &zeroLogger{logger, conf}
}

func (z *zeroLogger) Info(msg string, fields ...interface{}) {
	z.logger.Info().Fields(convertToMap(fields)).Msg(msg)
}

func (z *zeroLogger) Debug(msg string, keysAndValues ...interface{}) {
	z.logger.Debug().Fields(convertToMap(keysAndValues)).Msg(msg)
}

func (z *zeroLogger) Warn(msg string, keysAndValues ...interface{}) {
	z.logger.Warn().Fields(convertToMap(keysAndValues)).Msg(msg)
}

func (z *zeroLogger) Error(msg string, keysAndValues ...interface{}) {
	z.logger.Error().Fields(convertToMap(keysAndValues)).Msg(msg)
}

func (z *zeroLogger) With(fields ...interface{}) Logger {
	return &zeroLogger{logger: z.logger.With().Fields(convertToMap(fields)).Logger()}
}

func (z *zeroLogger) LogRequest(ctx context.Context, req *http.Request) time.Time {
	st := time.Now()
	fields := z.baseFields(ctx, req)
	fields = append(fields, "body", req.Body)
	fields = append(fields, "time", st.Format(z.conf.TimeFormat))
	z.logger.Info().Fields(convertToMap(fields)).Msg("Request")
	return st
}

func (z *zeroLogger) baseLogResponse(ctx context.Context, req *http.Request, st time.Time) []interface{} {
	fields := z.baseFields(ctx, req)
	end := time.Now()
	latency := end.Sub(st)
	if z.conf.UTC {
		end = end.UTC()
	}
	return append(fields, "latency", latency)
}

func (z *zeroLogger) LogResponse(ctx context.Context, req *http.Request, st time.Time, resp middleware2.Response) {
	fields := z.baseLogResponse(ctx, req, st)
	fn := z.logger.Info
	//Status parameter is available in GeneralError and GetStatus() is implemented in generated code
	if tresp, ok := resp.Type.(interface{ GetStatus() openapi.OptInt }); ok {
		sta, stOK := tresp.GetStatus().Get()
		if stOK {
			fields = append(fields, "status", sta)
		}
		if stOK && sta >= 500 {
			fn = z.logger.Error
		} else if stOK && sta >= 400 {
			fn = z.logger.Warn
		}
	}
	fn().Fields(convertToMap(fields)).Msg("Response")
}

func (z *zeroLogger) baseFields(ctx context.Context, req *http.Request) []interface{} {
	reqID := ctx.Value(middleware.RequestIDKey).(string)
	var fields []interface{}
	fields = append(fields, "request-id", reqID)
	fields = append(fields, "method", req.Method)
	fields = append(fields, "path", req.URL.Path)
	fields = append(fields, "url", req.URL.String())
	fields = append(fields, "query", req.URL.RawQuery)
	fields = append(fields, "ip", req.RemoteAddr)
	fields = append(fields, "user-agent", req.UserAgent())
	fields = append(fields, "protocol", req.Proto)
	return fields
}

func convertToMap(fields []interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			key, ok := fields[i].(string)
			if !ok {
				continue
			}
			result[key] = fields[i+1]
		}
	}
	return result
}