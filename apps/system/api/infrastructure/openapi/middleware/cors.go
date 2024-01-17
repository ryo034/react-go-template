package middleware

import (
	"github.com/ogen-go/ogen/middleware"
	middl "github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
)

func Cors(info *middl.CORSInfo) middleware.Middleware {
	return func(req middleware.Request, next func(req middleware.Request) (middleware.Response, error)) (middleware.Response, error) {
		middl.Cors(info, req.Raw)
		return next(req)
	}
}
