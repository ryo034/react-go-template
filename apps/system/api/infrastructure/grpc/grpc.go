package grpc

import (
	"fmt"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/grpc/interceptor"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/injector"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	healthConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/health/v1/v1connect"
	meConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/me/v1/v1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func newServeMuxWithReflection() *http.ServeMux {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		healthConnect.HealthServiceName,
		meConnect.MeServiceName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

func Start(conf config.Reader) {
	grpcEndpoint := fmt.Sprintf(":%s", conf.ServerPort())

	fb, err := firebase.NewFirebase(conf.IsLocal(), conf.FirebaseStorageBucket())
	if err != nil {
		log.Fatalln(err)
	}

	co := shared.NewContextOperator()

	auth := middleware.NewAuthentication(fb, co)

	//dbCloseFn := core.Initialize(conf.SourceDataSource(), conf.ReplicaDataSource(), conf.IsDebug())
	//defer dbCloseFn()

	inj, err := injector.NewInjector(
		fb,
		co,
		conf,
	)
	if err != nil {
		log.Fatalln(err)
	}

	mux := newServeMuxWithReflection()
	interceptors := interceptor.NewAuthInterceptor(auth)

	path, handler := healthConnect.NewHealthServiceHandler(inj.HealthServiceServer())
	mux.Handle(path, handler)
	path, handler = meConnect.NewMeServiceHandler(inj.MeServiceServer(), interceptors)
	mux.Handle(path, handler)

	log.Printf("Listening on port %s", conf.ServerPort())
	corsHandler := conf.Cors().Handler(h2c.NewHandler(mux, &http2.Server{}))
	if err = http.ListenAndServe(grpcEndpoint, corsHandler); err != nil {
		log.Fatal(err)
	}
}
