package injector

import (
	"github.com/redis/go-redis/v9"
	"github.com/ryo034/react-go-template/apps/system/api/driver/auth"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/driver/keyvalue"
	"github.com/ryo034/react-go-template/apps/system/api/driver/me"
	"github.com/ryo034/react-go-template/apps/system/api/driver/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/driver/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
)

type driver struct {
	KeyValue  keyvalue.Store
	Firebase  firebaseDriver.Driver
	Email     email.Driver
	Me        me.Driver
	Auth      auth.Driver
	Workspace workspace.Driver
	Member    member.Driver
}

func newDriverInjector(rc *redis.Client, f *firebase.Firebase) driver {
	meDr := me.NewDriver()
	return driver{
		keyvalue.NewRedisDriver(rc),
		firebaseDriver.NewDriver(f),
		email.NewDriver(),
		meDr,
		auth.NewDriver(),
		workspace.NewDriver(),
		member.NewDriver(),
	}
}
