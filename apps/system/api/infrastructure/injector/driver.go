package injector

import (
	"github.com/redis/go-redis/v9"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/driver/auth"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/driver/keyvalue"
	"github.com/ryo034/react-go-template/apps/system/api/driver/me"
	"github.com/ryo034/react-go-template/apps/system/api/driver/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/driver/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/mailer"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
)

type Driver struct {
	KeyValue  keyvalue.Store
	Firebase  firebaseDriver.Driver
	Email     email.Driver
	Me        me.Driver
	Auth      auth.Driver
	Workspace workspace.Driver
	Member    member.Driver
}

func newDriverInjector(
	conf config.Reader,
	rc *redis.Client,
	f *firebase.Firebase,
	co shared.ContextOperator,
	mc mailer.Client,
	noreplyEmail account.Email,
) Driver {
	meDr := me.NewDriver()
	return Driver{
		keyvalue.NewRedisDriver(rc),
		firebaseDriver.NewDriver(f),
		email.NewDriver(conf.ServiceName(), co, mc, noreplyEmail),
		meDr,
		auth.NewDriver(),
		workspace.NewDriver(),
		member.NewDriver(),
	}
}
