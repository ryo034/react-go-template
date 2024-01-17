package injector

import (
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/driver/me"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
)

type driver struct {
	Firebase firebaseDriver.Driver
	Me       me.Driver
}

func newDriverInjector(f *firebase.Firebase) driver {
	meDr := me.NewDriver()
	return driver{
		firebaseDriver.NewDriver(f),
		meDr,
	}
}
