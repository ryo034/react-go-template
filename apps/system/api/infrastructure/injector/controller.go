package injector

import (
	meCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me"
)

type Controller struct {
	Me meCtrl.Controller
}

func newControllerInjector(ui UseCase) Controller {
	return Controller{
		Me: meCtrl.NewController(ui.Me),
	}
}
