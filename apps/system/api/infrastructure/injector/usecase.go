package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
	"github.com/ryo034/react-go-template/apps/system/api/usecase/me"
	"github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

type UseCase struct {
	Me        me.UseCase
	Auth      auth.UseCase
	Workspace workspace.UseCase
}

func newUseCaseInjector(
	isLocal bool,
	co shared.ContextOperator,
	ri RepositoryInjector,
	pi Presenter,
	p core.Provider,
	txp core.TransactionProvider,
) UseCase {
	return UseCase{
		me.NewUseCase(txp, p, ri.Me, ri.Workspace, pi.Me),
		auth.NewUseCase(txp, p, ri.Auth, ri.Me, ri.Invitation, ri.Workspace, ri.Notification, pi.Auth),
		workspace.NewUseCase(txp, p, ri.Workspace, ri.Me, ri.Invitation, ri.Notification, pi.Workspace),
	}
}
