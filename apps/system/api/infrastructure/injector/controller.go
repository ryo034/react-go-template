package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/auth"
	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/workspace"
	sharedPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
)

type Controller struct {
	Me        me.Controller
	Auth      auth.Controller
	Workspace workspace.Controller
}

func newControllerInjector(
	ui UseCase,
	mr message.Resource,
	la sharedPresenter.LanguageAdapter,
	co shared.ContextOperator,
	di Driver,
) Controller {
	sr := sharedPresenter.NewResolver(mr, la)
	return Controller{
		me.NewController(ui.Me, sr, co),
		auth.NewController(ui.Auth, sr, co, di.Firebase),
		workspace.NewController(ui.Workspace, di.Firebase, sr, co),
	}
}
