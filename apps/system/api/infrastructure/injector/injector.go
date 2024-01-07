package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
)

type Injector struct {
	f          *fb.Firebase
	conf       config.Reader
	co         shared.ContextOperator
	driverInj  driverInjector
	useCaseInj UseCaseInjector
	resi       ResponseInjector
	reqi       RequestInjector
}

func NewInjector(
	f *fb.Firebase,
	co shared.ContextOperator,
	conf config.Reader,
) (*Injector, error) {
	di := newDriverInjector(f)
	ri := newRepositoryInjector(di)
	ui := newUseCaseInjector(conf.IsLocal(), co, ri, di)
	resi := newResponseInjector()
	reqi := newRequestInjector()
	return &Injector{
		f,
		conf,
		co,
		di,
		ui,
		resi,
		reqi,
	}, nil
}

func (i *Injector) UseCase() UseCaseInjector {
	return i.useCaseInj
}
