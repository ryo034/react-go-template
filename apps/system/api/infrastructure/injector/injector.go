package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	sharedPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
)

type Injector struct {
	f          *fb.Firebase
	conf       config.Reader
	co         shared.ContextOperator
	driverInj  driver
	useCaseInj UseCase
	ctrl       Controller
	resi       ResponseInjector
	reqi       RequestInjector
}

func NewInjector(
	f *fb.Firebase,
	d core.Provider,
	txp core.TransactionProvider,
	co shared.ContextOperator,
	conf config.Reader,
) (*Injector, error) {
	defaultLang := conf.DefaultLanguage()
	di := newDriverInjector(f)
	ri := newRepositoryInjector(di, newGatewayAdapterInjector())
	pi := newPresenterInjector()
	la := sharedPresenter.NewLanguageAdapter(defaultLang)
	messageResource := message.NewResource(defaultLang)
	ui := newUseCaseInjector(conf.IsLocal(), co, ri, di, pi, d, txp, messageResource, la)
	resi := newResponseInjector()
	reqi := newRequestInjector()
	ctrl := newControllerInjector(ui)
	return &Injector{
		f,
		conf,
		co,
		di,
		ui,
		ctrl,
		resi,
		reqi,
	}, nil
}

func (i *Injector) UseCase() UseCase {
	return i.useCaseInj
}
func (i *Injector) Controller() Controller {
	return i.ctrl
}
