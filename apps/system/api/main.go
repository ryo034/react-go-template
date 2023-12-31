package main

import (
	"os"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/grpc"
)

func main() {
	conf := config.NewReader(env())
	grpc.Start(conf)
}

func env() config.Env {
	e := os.Getenv("ENV")
	if e == "" {
		return config.Local
	}
	return config.Env(e)
}
