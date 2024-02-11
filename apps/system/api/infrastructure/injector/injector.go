package injector

import (
	"github.com/redis/go-redis/v9"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/mailer"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/message"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	sharedPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
)

type Injector struct {
	f          *fb.Firebase
	conf       config.Reader
	co         shared.ContextOperator
	driverInj  Driver
	useCaseInj UseCase
	ctrl       Controller
}

func NewInjector(
	f *fb.Firebase,
	d core.Provider,
	txp core.TransactionProvider,
	co shared.ContextOperator,
	conf config.Reader,
	rc *redis.Client,
	mc mailer.Client,
) (*Injector, error) {
	defaultLang := conf.DefaultLanguage()
	di := newDriverInjector(rc, f, co, mc, conf.NoReplyEmail())
	ri := newRepositoryInjector(di, newGatewayAdapterInjector())
	pi := newPresenterInjector()
	la := sharedPresenter.NewLanguageAdapter(defaultLang)
	messageResource := message.NewResource(defaultLang)
	ui := newUseCaseInjector(conf.IsLocal(), co, ri, di, pi, d, txp)
	ctrl := newControllerInjector(ui, messageResource, la, co)
	return &Injector{
		f,
		conf,
		co,
		di,
		ui,
		ctrl,
	}, nil
}

func (i *Injector) UseCase() UseCase {
	return i.useCaseInj
}
func (i *Injector) Controller() Controller {
	return i.ctrl
}

func (i *Injector) Driver() Driver {
	return i.driverInj
}
