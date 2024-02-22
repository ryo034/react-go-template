package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/auth"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/user"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace/invitation"
	authUc "github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

type Presenter struct {
	Auth      authUc.OutputPort
	Me        meUc.OutputPort
	Workspace workspaceUc.OutputPort
}

func newPresenterInjector() Presenter {
	pa := newPresenterAdapter()
	m := member.NewAdapter(pa.User)
	inv := invitation.NewAdapter()
	meAdapter := me.NewAdapter(pa.User, pa.Member, pa.Workspace, inv)
	mp := me.NewPresenter(meAdapter)
	return Presenter{
		auth.NewPresenter(meAdapter),
		mp,
		workspace.NewPresenter(pa.Workspace, pa.Invitation, m),
	}
}

type PresenterAdapter struct {
	User       user.Adapter
	Member     member.Adapter
	Workspace  workspace.Adapter
	Invitation invitation.Adapter
}

func newPresenterAdapter() PresenterAdapter {
	ua := user.NewAdapter()
	ma := member.NewAdapter(ua)
	return PresenterAdapter{
		ua,
		ma,
		workspace.NewAdapter(ma),
		invitation.NewAdapter(),
	}
}
