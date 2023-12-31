package injector

import (
	userGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
)

type gatewayAdapterInjector struct {
	User userGateway.Adapter
}

func newGatewayAdapterInjector() gatewayAdapterInjector {
	user := userGateway.NewAdapter()
	return gatewayAdapterInjector{
		User: user,
	}
}
