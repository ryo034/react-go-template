package openapi

import (
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/injector"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/openapi/middleware"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/openapi/service"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"log"
	"net/http"
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

	inj, err := injector.NewInjector(fb, p, txp, co, conf)
	if err != nil {
		log.Fatalln(err)
	}

	srv, err := openapi.NewServer(
		service.NewService(inj),
		middleware.NewSecMiddleware(),
		openapi.WithMiddleware(middleware.NewMiddlewares().Global(conf)...),
	)
	if err != nil {
		log.Fatal(err)
	}
	corsHandler := conf.Cors().Handler(srv)
	if err = http.ListenAndServe(endpoint, corsHandler); err != nil {
		log.Fatal(err)
	}
}
