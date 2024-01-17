package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/me"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Presenter struct {
	Me meUc.OutputPort
}

func newPresenterInjector() Presenter {
	return Presenter{
		Me: me.NewPresenter(),
	}
}
