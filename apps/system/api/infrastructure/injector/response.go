package injector

import (
	employeeResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/business_entity/employee/response"
	businessEntityResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/business_entity/response"
	creatureResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/creature/response"
	meResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me/response"
	addressResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/address"
	paginationResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/pagination/response"
	storeCreatureResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/creature/response"
	itemResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/item/response"
	storeResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/response"
	staffResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/staff/response"
	transactionResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/transaction/response"
	userResponse "github.com/ryo034/react-go-template/apps/system/api/interface/controller/user/response"
)

type ResponseInjector struct {
	Me             meResponse.Adapter
	Staff          staffResponse.Adapter
	Pagination     paginationResponse.Adapter
	Employee       employeeResponse.Adapter
	BusinessEntity businessEntityResponse.Adapter
	Store          storeResponse.Adapter
	Item           itemResponse.Adapter
	StoreCreature  storeCreatureResponse.Adapter
	Transaction    transactionResponse.Adapter
}

func newResponseInjector() ResponseInjector {
	ut := userResponse.NewAdapter()
	st := staffResponse.NewAdapter(ut)
	aa := addressResponse.NewAdapter()
	er := employeeResponse.NewAdapter(ut)
	crr := creatureResponse.NewAdapter()
	return ResponseInjector{
		Me:             meResponse.NewAdapter(ut, er),
		Staff:          st,
		Pagination:     paginationResponse.NewAdapter(),
		Employee:       er,
		BusinessEntity: businessEntityResponse.NewAdapter(er, aa),
		Store:          storeResponse.NewAdapter(aa),
		Item:           itemResponse.NewAdapter(),
		StoreCreature:  storeCreatureResponse.NewAdapter(crr),
		Transaction:    transactionResponse.NewAdapter(),
	}
}
