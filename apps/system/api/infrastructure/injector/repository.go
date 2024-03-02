package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/notification"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	authGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/auth"
	meGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
	notificationGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/notification"
	workspaceGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace"
	invitationGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace/invitation"
)

type RepositoryInjector struct {
	Notification notification.Repository
	Me           me.Repository
	Auth         auth.Repository
	Workspace    workspace.Repository
	Invitation   invitation.Repository
}

func newRepositoryInjector(co shared.ContextOperator, di Driver, gw GatewayAdapter) RepositoryInjector {
	return RepositoryInjector{
		notificationGw.NewGateway(di.Email),
		meGw.NewGateway(di.Me, di.Firebase, di.Workspace, di.Invitation, gw.Me, co),
		authGw.NewGateway(di.KeyValue, di.Auth, di.Firebase, gw.Auth),
		workspaceGw.NewGateway(di.Workspace, di.Member, di.Invitation, gw.Workspace, gw.Member, gw.Invitation),
		invitationGw.NewGateway(di.Invitation, gw.Invitation),
	}
}
