package openapi

import (
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/redis"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/injector"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
	openapiMiddleware "github.com/ryo034/react-go-template/apps/system/api/infrastructure/openapi/middleware"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/openapi/service"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"log"
	"net/http"
	"time"
)

func Start(conf config.Reader) {
	endpoint := fmt.Sprintf(":%s", conf.ServerPort())

	fb, err := firebase.NewFirebase(conf.IsLocal(), conf.FirebaseStorageBucket())
	if err != nil {
		log.Fatalln(err)
	}
	co := shared.NewContextOperator()

	dbCloseFn, p, txp := core.Initialize(conf.SourceDataSource(), conf.ReplicaDataSource(), conf.IsDebug())
	defer dbCloseFn()

	rc := redis.NewRedisClient(conf.RedisConfig())

	inj, err := injector.NewInjector(fb, p, txp, co, conf, rc)
	if err != nil {
		log.Fatalln(err)
	}

	h, err := openapi.NewServer(
		service.NewService(inj),
		openapiMiddleware.NewSecMiddleware(),
	)

	if err != nil {
		log.Fatal(err)
	}

	zl := logger.NewZeroLogger(logger.Config{TimeFormat: time.RFC3339, UTC: true}, conf.IsLocal(), conf.ServiceName())

	if err = http.ListenAndServe(
		endpoint,
		middleware.NewMiddlewares().Global(h, conf, zl, rc),
	); err != nil {
		log.Fatal(err)
	}
}
