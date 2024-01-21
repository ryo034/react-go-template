package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/auth"
	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/me"
	sharedPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
)

type Controller struct {
	Me   me.Controller
	Auth auth.Controller
}

func newControllerInjector(
	ui UseCase,
	mr message.Resource,
	la sharedPresenter.LanguageAdapter,
) Controller {
	sr := sharedPresenter.NewResolver(mr, la)
	return Controller{
		Me:   me.NewController(ui.Me, sr),
		Auth: auth.NewController(ui.Auth, sr),
	}
}
