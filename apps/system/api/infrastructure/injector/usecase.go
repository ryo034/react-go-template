package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
	"github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type UseCase struct {
	Me   me.UseCase
	Auth auth.UseCase
}

func newUseCaseInjector(
	isLocal bool,
	co shared.ContextOperator,
	ri RepositoryInjector,
	di driver,
	pi Presenter,
	p core.Provider,
	txp core.TransactionProvider,
) UseCase {
	return UseCase{
		Me:   me.NewUseCase(txp, p, ri.Me, pi.Me),
		Auth: auth.NewUseCase(txp, p, ri.Auth, di.Email, di.Firebase),
	}
}
