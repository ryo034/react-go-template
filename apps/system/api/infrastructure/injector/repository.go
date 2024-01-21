package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	authGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/auth"
	meGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
)

type RepositoryInjector struct {
	Me   me.Repository
	Auth auth.Repository
}

func newRepositoryInjector(di driver, gw GatewayAdapter) RepositoryInjector {
	return RepositoryInjector{
		Me:   meGw.NewGateway(di.Me, di.Firebase, gw.Me),
		Auth: authGw.NewGateway(di.KeyValue, di.Auth, gw.Auth),
	}
}
