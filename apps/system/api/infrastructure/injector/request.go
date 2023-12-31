package injector

import (
	meRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me/request"
	userRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/user/request"
)

type RequestInjector struct {
	Me meRequest.Adapter
}

func newRequestInjector() RequestInjector {
	ut := userRequest.NewAdapter()
	return RequestInjector{
		Me: meRequest.NewAdapter(ut),
	}
}
