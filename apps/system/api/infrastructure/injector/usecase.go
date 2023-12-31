package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	meUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type useCaseInjector struct {
	Me meUseCase.UseCase
}

func newUseCaseInjector(
	isLocal bool,
	co shared.ContextOperator,
	ri RepositoryInjector,
	di driverInjector,
) useCaseInjector {
	return useCaseInjector{
		Me: meUseCase.NewUseCase(isLocal, ri.Me),
	}
}
