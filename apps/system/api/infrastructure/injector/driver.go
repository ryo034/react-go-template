package injector

import (
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/driver/auth"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/driver/keyvalue"
	"github.com/ryo034/react-go-template/apps/system/api/driver/me"
	"github.com/ryo034/react-go-template/apps/system/api/driver/media"
	"github.com/ryo034/react-go-template/apps/system/api/driver/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/driver/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/driver/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/mailer"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage"
)

type Driver struct {
	KeyValue   keyvalue.Store
	Firebase   firebaseDriver.Driver
	Email      email.Driver
	Me         me.Driver
	Auth       auth.Driver
	Workspace  workspace.Driver
	Member     member.Driver
	Invitation invitation.Driver
	Media      media.Driver
}

func newDriverInjector(
	conf config.Reader,
	logger logger.Logger,
	rc *redis.Client,
	f *firebase.Firebase,
	co shared.ContextOperator,
	mc mailer.Client,
	minioClient *minio.Client,
	sh storage.Handler,
	noreplyEmail account.Email,
) Driver {
	invDr := invitation.NewDriver()
	meDr := me.NewDriver(invDr)
	return Driver{
		keyvalue.NewRedisDriver(rc),
		firebaseDriver.NewDriver(f, co, sh),
		email.NewDriver(conf.ServiceName(), co, mc, noreplyEmail, logger),
		meDr,
		auth.NewDriver(),
		workspace.NewDriver(),
		member.NewDriver(),
		invDr,
		media.NewDriver(minioClient, conf.MinioConfig()),
	}
}
