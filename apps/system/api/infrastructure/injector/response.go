package injector

import (
	meResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me/response"
	userResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/user/response"
)

type ResponseInjector struct {
	Me meResponse.Adapter
}

func newResponseInjector() ResponseInjector {
	ut := userResponse.NewAdapter()
	return ResponseInjector{
		Me: meResponse.NewAdapter(ut),
	}
}
