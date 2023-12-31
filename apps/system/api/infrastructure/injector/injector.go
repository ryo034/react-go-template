package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/sqlboiler/core"
	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/grpc/response"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	healthCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/health"
	meCtrl "github.com/ryo034/react-go-template/apps/system/api/interface/controller/me"
	languageAdapter "github.com/ryo034/react-go-template/apps/system/api/interface/controller/shared/language"
	healthConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/health/v1/v1connect"
	meConnect "github.com/ryo034/react-go-template/apps/system/api/schema/pb/me/v1/v1connect"
	"golang.org/x/text/language"
)

type Injector interface {
	HealthServiceServer() healthConnect.HealthServiceHandler
	MeServiceServer() meConnect.MeServiceHandler
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
