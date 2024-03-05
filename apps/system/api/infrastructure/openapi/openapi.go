package openapi

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage/minio"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/redis"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/injector"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/mailer"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
	openapiMiddleware "github.com/ryo034/react-go-template/apps/system/api/infrastructure/openapi/middleware"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/openapi/service"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
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

	minioClient, err := minio.NewMinioClient(conf.MinioConfig())
	if err != nil {
		log.Fatalln(err)
		return
	}

	var mc mailer.Client
	if conf.IsLocal() {
		mc = mailer.NewMailhogMailer(conf.MailHost(), conf.MailPort())
	} else {
		mc = mailer.NewResendMailer(conf.ResendAPIKey())
	}

	zl := logger.NewZeroLogger(logger.Config{TimeFormat: time.RFC3339, UTC: true}, conf.IsLocal(), conf.ServiceName())

	inj, err := injector.NewInjector(fb, p, txp, co, conf, rc, mc, zl, minioClient)
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

	server := &http.Server{
		Addr:         endpoint,
		Handler:      middleware.NewMiddlewares(co).Global(h, conf, zl, rc, p, inj.Driver()),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
