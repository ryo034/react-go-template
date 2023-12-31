package injector

import (
	meRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me/request"
	mediaRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/media"
	paginationRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/pagination/request"
	itemRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/item/request"
	staffRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/staff/request"
	transactionRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/transaction/request"
	userRequest "github.com/ryo034/react-go-template/apps/system/api/interface/controller/user/request"
)

type RequestInjector struct {
	Me          meRequest.Adapter
	Staff       staffRequest.Adapter
	Item        itemRequest.Adapter
	Pagination  paginationRequest.Adapter
	Transaction transactionRequest.Adapter
}

func newRequestInjector() RequestInjector {
	ut := userRequest.NewAdapter()
	st := staffRequest.NewAdapter(ut)
	med := mediaRequest.NewAdapter()
	//org := organizationRequest.NewAdapter()
	return RequestInjector{
		Me:          meRequest.NewAdapter(ut),
		Staff:       st,
		Item:        itemRequest.NewAdapter(med),
		Pagination:  paginationRequest.NewAdapter(),
		Transaction: transactionRequest.NewAdapter(),
	}
}
