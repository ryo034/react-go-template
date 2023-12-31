package injector

import (
	businessEntityGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/business_entity"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/business_entity/employee"
	employeeGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/business_entity/employee"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/business_entity/employee/permission"
	meGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
	storeGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store"
	itemGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/item"
	staffGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/staff"
	meStaffGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/staff/me"
	transactionGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/store/transaction"
	userGateway "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
	sharedGateway "github.com/ryo034/react-go-template/packages/go/interface/gateway/shared"
	"github.com/ryo034/react-go-template/packages/go/interface/gateway/shared/media"
)

type gatewayAdapterInjector struct {
	User           userGateway.Adapter
	Me             meGateway.Adapter
	Store          storeGateway.Adapter
	Item           itemGateway.Adapter
	Staff          staffGateway.Adapter
	MeStaff        meStaffGateway.Adapter
	BusinessEntity businessEntityGateway.Adapter
	Employee       employeeGateway.Adapter
	Transaction    transactionGateway.Adapter
}

func newGatewayAdapterInjector() gatewayAdapterInjector {
	shared := sharedGateway.NewAdapter()
	user := userGateway.NewAdapter()
	staff := staffGateway.NewAdapter(user)
	st := storeGateway.NewAdapter(shared)
	ep := permission.NewAdapter()
	emp := employee.NewAdapter(user, ep)
	med := media.NewAdapter()
	return gatewayAdapterInjector{
		User:           user,
		Me:             meGateway.NewAdapter(user, emp, st, staff),
		Store:          st,
		Item:           itemGateway.NewAdapter(med),
		Staff:          staff,
		MeStaff:        meStaffGateway.NewAdapter(staff),
		BusinessEntity: businessEntityGateway.NewAdapter(shared, emp),
		Employee:       emp,
		Transaction:    transactionGateway.NewAdapter(),
	}
}
