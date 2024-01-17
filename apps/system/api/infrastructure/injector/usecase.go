package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	sharedPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	meUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type UseCase struct {
	Me meUseCase.UseCase
}

func newUseCaseInjector(
	isLocal bool,
	co shared.ContextOperator,
	ri RepositoryInjector,
	di driver,
	pi Presenter,
	p core.Provider,
	txp core.TransactionProvider,
	mr message.Resource,
	la sharedPresenter.LanguageAdapter,
) UseCase {
	sr := sharedPresenter.NewResolver(mr, la)
	return UseCase{
		Me: meUseCase.NewUseCase(txp, p, ri.Me, pi.Me, sr),
	}
}
