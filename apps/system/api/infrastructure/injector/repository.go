package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	authGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/auth"
	meGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
	workspaceGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace"
	invitationGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace/invitation"
)

type RepositoryInjector struct {
	Me         me.Repository
	Auth       auth.Repository
	Workspace  workspace.Repository
	Invitation invitation.Repository
}

func newRepositoryInjector(di Driver, gw GatewayAdapter) RepositoryInjector {
	return RepositoryInjector{
		meGw.NewGateway(di.Me, di.Firebase, di.Workspace, di.Invitation, gw.Me),
		authGw.NewGateway(di.KeyValue, di.Auth, di.Firebase, gw.Auth),
		workspaceGw.NewGateway(di.Workspace, di.Member, di.Invitation, gw.Workspace, gw.Member, gw.Invitation),
		invitationGw.NewGateway(di.Invitation, gw.Invitation),
	}
}
