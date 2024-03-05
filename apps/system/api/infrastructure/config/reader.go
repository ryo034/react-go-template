package config

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/cors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/datasource"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/redis"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage/minio"

	"github.com/spf13/cast"
	"golang.org/x/text/language"
)

type Reader interface {
	ServiceName() string
	IsLocal() bool
	IsDebug() bool
	TimeLocation() *time.Location
	FixedTime() string
	AllowOrigins() []string
	Cors() *cors.Cors
	DefaultLanguage() language.Tag
	SourceDataSource() datasource.DataSource
	ReplicaDataSource() datasource.DataSource
	RedisConfig() *redis.Config
	MinioConfig() *minio.Config
	FirebaseStorageBucket() string
	ServerPort() string
	NoReplyEmail() account.Email
	MailHost() string
	MailPort() int
	ResendAPIKey() string
}

type Env string

const (
	Local Env = "local"
)

func (e Env) isLocal() bool {
	return Local == e
}

type Key string

const (
	timeLocation    Key = "TIME_LOCATION"
	defaultLanguage Key = "DEFAULT_LANGUAGE"
	nowOrToday      Key = "NOW_OR_TODAY"
	isDebug         Key = "IS_DEBUG"
	serviceName     Key = "SERVICE_NAME"
)

type reader struct {
	env Env
}

func (r *reader) ServiceName() string {
	return r.fromEnv(serviceName)
}

func (r *reader) IsLocal() bool {
	return r.env.isLocal()
}

func (r *reader) IsNotLocal() bool {
	return !r.IsLocal()
}

func (r *reader) IsDebug() bool {
	if r.IsLocal() {
		return true
	}
	return cast.ToBool(r.fromEnv(isDebug))
}

func NewReader(env Env) Reader {
	return &reader{env}
}

func (r *reader) TimeLocation() *time.Location {
	locStr := property(r.env, timeLocation)
	if result, err := time.LoadLocation(locStr); err != nil {
		panic(err)
	} else {
		return result
	}
}

func (r *reader) FixedTime() string {
	return property(r.env, nowOrToday)
}

func (r *reader) DefaultLanguage() language.Tag {
	lTagStr := property(r.env, defaultLanguage)
	if result, err := language.Parse(lTagStr); err != nil {
		panic(err)
	} else {
		return result
	}
}

func (r *reader) fromEnvUint(key Key) uint {
	if result, err := strconv.Atoi(r.fromEnv(key)); err != nil {
		panic(err)
	} else {
		return uint(result)
	}
}
func (r *reader) fromEnvInt(key Key) int {
	if result, err := strconv.Atoi(r.fromEnv(key)); err != nil {
		panic(err)
	} else {
		return result
	}
}

func (r *reader) fromEnv(key Key) string {
	if r.env.isLocal() {
		return localValues[key]
	}
	return os.Getenv(string(key))
}
