package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/user"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Presenter struct {
	Me meUc.OutputPort
}

func newPresenterInjector() Presenter {
	pa := newPresenterAdapter()
	return Presenter{
		Me: me.NewPresenter(pa.User, pa.Member, pa.Workspace),
	}
}

type PresenterAdapter struct {
	User      user.Adapter
	Workspace workspace.Adapter
	Member    member.Adapter
}

func newPresenterAdapter() PresenterAdapter {
	ua := user.NewAdapter()
	return PresenterAdapter{
		User:      ua,
		Workspace: workspace.NewAdapter(),
		Member:    member.NewAdapter(ua),
	}
}
