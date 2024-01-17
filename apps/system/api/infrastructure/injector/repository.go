package injector

import (
	meDomain "github.com/ryo034/react-go-template/apps/system/api/domain/me"
	meGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
)

type RepositoryInjector struct {
	Me meDomain.Repository
}

func newRepositoryInjector(di driver, gw GatewayAdapter) RepositoryInjector {
	return RepositoryInjector{
		Me: meGateway.NewGateway(di.Me, di.Firebase, gw.Me),
	}
}
