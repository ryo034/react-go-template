package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/auth"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace/invitation"
)

type GatewayAdapter struct {
	User       user.Adapter
	Member     member.Adapter
	Me         me.Adapter
	Workspace  workspace.Adapter
	Auth       auth.Adapter
	Invitation invitation.Adapter
}

func newGatewayAdapterInjector() GatewayAdapter {
	authGw := auth.NewAdapter()
	userGw := user.NewAdapter()
	memberGw := member.NewAdapter(userGw)
	inv := invitation.NewAdapter()
	workspaceGw := workspace.NewAdapter(memberGw)
	meGw := me.NewAdapter(userGw, workspaceGw, memberGw, inv)
	return GatewayAdapter{
		userGw,
		memberGw,
		meGw,
		workspaceGw,
		authGw,
		inv,
	}
}
