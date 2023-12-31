package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/grpc/response"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	beCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/business_entity"
	healthCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/health"
	meCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me"
	languageAdapter "github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/language"
	storeCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store"
	itemCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/item"
	transactionCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/store/transaction"
	beConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/business_entity/v1/v1connect"
	healthConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/health/v1/v1connect"
	meConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/me/v1/v1connect"
	itemConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/store/item/v1/v1connect"
	traConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/store/transaction/v1/v1connect"
	stConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/store/v1/v1connect"
	"github.com/ryo034/react-go-template/packages/go/infrastructure/database/sqlboiler/core"
	"github.com/ryo034/react-go-template/packages/go/infrastructure/message"
	"golang.org/x/text/language"
)

type Injector interface {
	HealthServiceServer() healthConnect.HealthServiceHandler
	MeServiceServer() meConnect.MeServiceHandler
	BusinessEntityServiceServer() beConnect.BusinessEntityServiceHandler
	StoreServiceServer() stConnect.StoreServiceHandler
	ItemServiceServer() itemConnect.ItemServiceHandler
	TransactionServiceServer() traConnect.TransactionServiceHandler
}

type injector struct {
	conf       config.Reader
	co         shared.ContextOperator
	driverInj  driverInjector
	useCaseInj useCaseInjector
	resi       ResponseInjector
	reqi       RequestInjector
	reslv      response.Resolver
}

func NewInjector(
	txp core.Provider,
	f *fb.Firebase,
	co shared.ContextOperator,
	conf config.Reader,
) (Injector, error) {
	di := newDriverInjector(f)
	ri := newRepositoryInjector(di)
	ui := newUseCaseInjector(conf.IsLocal(), txp, co, ri, di)
	resi := newResponseInjector()
	reqi := newRequestInjector()
	reslv := response.NewResolver(message.NewResource(conf.DefaultLanguage()), languageAdapter.NewAdapter(language.Japanese, co))
	return &injector{
		conf,
		co,
		di,
		ui,
		resi,
		reqi,
		reslv,
	}, nil
}

func (i *injector) HealthServiceServer() healthConnect.HealthServiceHandler {
	return healthCtrl.NewServer()
}

func (i *injector) MeServiceServer() meConnect.MeServiceHandler {
	return meCtrl.NewServer(i.co, i.driverInj.Firebase, i.useCaseInj.Me, i.reslv, i.reqi.Me, i.resi.Me, i.resi.BusinessEntity, i.resi.Store)
}

func (i *injector) BusinessEntityServiceServer() beConnect.BusinessEntityServiceHandler {
	return beCtrl.NewServer(i.co, i.driverInj.Firebase, i.useCaseInj.BusinessEntity, i.reslv, i.resi.BusinessEntity)
}

func (i *injector) StoreServiceServer() stConnect.StoreServiceHandler {
	return storeCtrl.NewServer(i.co, i.driverInj.Firebase, i.useCaseInj.Store, i.reslv, i.resi.Store, i.resi.StoreCreature)
}

func (i *injector) ItemServiceServer() itemConnect.ItemServiceHandler {
	return itemCtrl.NewServer(i.co, i.useCaseInj.Item, i.reslv, i.reqi.Item, i.resi.Item, i.reqi.Pagination, i.resi.Pagination)
}

func (i *injector) TransactionServiceServer() traConnect.TransactionServiceHandler {
	return transactionCtrl.NewServer(i.co, i.useCaseInj.Transaction, i.reslv, i.reqi.Transaction, i.resi.Transaction, i.reqi.Pagination, i.resi.Pagination)
}
