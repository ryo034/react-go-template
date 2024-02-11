package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	authGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/auth"
	meGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
	workspaceGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace"
)

type RepositoryInjector struct {
	Me        me.Repository
	Auth      auth.Repository
	Workspace workspace.Repository
}

func newRepositoryInjector(di Driver, gw GatewayAdapter) RepositoryInjector {
	return RepositoryInjector{
		meGw.NewGateway(di.Me, di.Firebase, di.Workspace, gw.Me),
		authGw.NewGateway(di.KeyValue, di.Auth, gw.Auth),
		workspaceGw.NewGateway(di.Workspace, di.Member, gw.Workspace, gw.Member),
	}
}
