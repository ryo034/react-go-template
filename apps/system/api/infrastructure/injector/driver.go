package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/driver/auth"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/driver/keyvalue"
	"github.com/ryo034/react-go-template/apps/system/api/driver/me"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
)

type driver struct {
	Firebase firebaseDriver.Driver
	Me       me.Driver
	Auth     auth.Driver
	Email    email.Driver
	KeyValue keyvalue.Store
}

func newDriverInjector(f *firebase.Firebase) driver {
	meDr := me.NewDriver()
	return driver{
		firebaseDriver.NewDriver(f),
		meDr,
		auth.NewDriver(),
		email.NewDriver(),
		keyvalue.NewRedisDriver("", "", 1),
	}
}
