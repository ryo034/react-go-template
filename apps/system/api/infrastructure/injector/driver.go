package injector

import (
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
)

type driverInjector struct {
	Firebase firebaseDriver.Driver
}

func newDriverInjector(f *firebase.Firebase) driverInjector {
	return driverInjector{
		Firebase: firebaseDriver.NewDriver(f),
	}
}
