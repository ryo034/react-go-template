package openapi

import (
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/injector"
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

	inj, err := injector.NewInjector(fb, co, conf)
	if err != nil {
		log.Fatalln(err)
	}

	srv, err := openapi.NewServer(NewService(inj))
	if err != nil {
		log.Fatal(err)
	}
	if err = http.ListenAndServe(endpoint, srv); err != nil {
		log.Fatal(err)
	}
}
