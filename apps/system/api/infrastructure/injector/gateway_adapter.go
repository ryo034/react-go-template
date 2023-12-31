package injector

import (
	meGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
	userGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
)

type gatewayAdapterInjector struct {
	User userGateway.Adapter
	Me   meGateway.Adapter
}

func newGatewayAdapterInjector() gatewayAdapterInjector {
	user := userGateway.NewAdapter()
	return gatewayAdapterInjector{
		User: user,
		Me:   meGateway.NewAdapter(user),
	}
}
