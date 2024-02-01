package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/user"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

type Presenter struct {
	Me        meUc.OutputPort
	Workspace workspaceUc.OutputPort
}

func newPresenterInjector() Presenter {
	pa := newPresenterAdapter()
	return Presenter{
		me.NewPresenter(pa.User, pa.Member, pa.Workspace),
		workspace.NewPresenter(pa.Workspace),
	}
}

type PresenterAdapter struct {
	User      user.Adapter
	Member    member.Adapter
	Workspace workspace.Adapter
}

func newPresenterAdapter() PresenterAdapter {
	ua := user.NewAdapter()
	return PresenterAdapter{
		ua,
		member.NewAdapter(ua),
		workspace.NewAdapter(),
	}
}
