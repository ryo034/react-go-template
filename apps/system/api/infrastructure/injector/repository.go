package injector

import (
	beDomain "github.com/ryo034/react-go-template/apps/system/api/domain/business_entity"
	employeeDomain "github.com/ryo034/react-go-template/apps/system/api/domain/business_entity/employee"
	meDomain "github.com/ryo034/react-go-template/apps/system/api/domain/me"
	storeDomain "github.com/ryo034/react-go-template/apps/system/api/domain/store"
	itemDomain "github.com/ryo034/react-go-template/apps/system/api/domain/store/item"
	staffDomain "github.com/ryo034/react-go-template/apps/system/api/domain/store/staff"
	meStaffDomain "github.com/ryo034/react-go-template/apps/system/api/domain/store/staff/me"
	transactionDomain "github.com/ryo034/react-go-template/apps/system/api/domain/store/transaction"
	beGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/business_entity"
	employeeGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/business_entity/employee"
	meGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
	storeGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store"
	storeCreatureGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/creature"
	itemGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/item"
	staffGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/staff"
	meStaffGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/staff/me"
	transactionGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/transaction"
	creatureGateway "github.com/ryo034/react-go-template/packages/go/interface/gateway/creature"
	mediaGateway "github.com/ryo034/react-go-template/packages/go/interface/gateway/shared/media"
)

type RepositoryInjector struct {
	Me             meDomain.Repository
	Store          storeDomain.Repository
	Item           itemDomain.Repository
	BusinessEntity beDomain.Repository
	Employee       employeeDomain.Repository
	Staff          staffDomain.Repository
	MeStaff        meStaffDomain.Repository
	Transaction    transactionDomain.Repository
}

func newRepositoryInjector(di driverInjector) RepositoryInjector {
	gai := newGatewayAdapterInjector()
	mga := mediaGateway.NewAdapter()
	cr := creatureGateway.NewAdapter(mga)
	scr := storeCreatureGateway.NewAdapter(cr)
	return RepositoryInjector{
		Me:             meGateway.NewRepository(di.Me, di.Firebase, gai.Me, gai.Store, gai.BusinessEntity),
		Store:          storeGateway.NewRepository(di.Store, gai.Store, scr),
		Item:           itemGateway.NewRepository(di.Item, gai.Item),
		BusinessEntity: beGateway.NewRepository(di.BusinessEntity, gai.BusinessEntity),
		Employee:       employeeGateway.NewRepository(di.Employee, gai.Employee),
		Staff:          staffGateway.NewRepository(di.Staff, gai.Staff),
		MeStaff:        meStaffGateway.NewRepository(di.Staff, di.Firebase, gai.MeStaff),
		Transaction:    transactionGateway.NewRepository(di.Transaction, gai.Transaction),
	}
}
