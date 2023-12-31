package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	beUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/business_entity"
	meUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
	storeUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/store"
	itemUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/store/item"
	meStaffUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/store/staff/me"
	transactionUseCase "github.com/ryo034/react-go-template/apps/system/api/usecase/store/transaction"
	"github.com/ryo034/react-go-template/packages/go/infrastructure/database/sqlboiler/core"
)

type useCaseInjector struct {
	Me             meUseCase.UseCase
	Store          storeUseCase.UseCase
	Item           itemUseCase.UseCase
	BusinessEntity beUseCase.UseCase
	MeStaff        meStaffUseCase.UseCase
	Transaction    transactionUseCase.UseCase
}

func newUseCaseInjector(
	isLocal bool,
	txp core.Provider,
	co shared.ContextOperator,
	ri RepositoryInjector,
	di driverInjector,
) useCaseInjector {
	return useCaseInjector{
		Me:             meUseCase.NewUseCase(isLocal, txp, ri.Me),
		Store:          storeUseCase.NewUseCase(txp, co, ri.Store, ri.Staff, ri.Me, ri.BusinessEntity, ri.Employee, di.Firebase),
		Item:           itemUseCase.NewUseCase(txp, co, ri.Item, ri.Store, ri.Staff),
		BusinessEntity: beUseCase.NewUseCase(txp, co, ri.BusinessEntity),
		MeStaff:        meStaffUseCase.NewUseCase(txp, ri.MeStaff),
		Transaction:    transactionUseCase.NewUseCase(txp, co, ri.Transaction, ri.Staff),
	}
}
