package config

import (
	"os"
)

type Reader interface {
	IsLocal() bool
	ApplicationName() string
	GcpProjectID() string

	MaterialImageStorageSchemaHost() string
	AccountProfileImageStorageSchemaHost() string
	TankImageStorageSchemaHost() string
	PostImageStorageSchemaHost() string

	EventFollowSubscription() string
	EventCreateTankSubscription() string
	EventLikeTankSubscription() string
	EventLikePostSubscription() string
	EventPostCommentSubscription() string
	EventReplyPostCommentSubscription() string
	EventDeletePostSubscription() string
	EventDeleteTankSubscription() string

	EventFollowTopic() string
	EventCreateTankTopic() string
	EventLikeTankTopic() string
	EventLikePostTopic() string
	EventPostCommentTopic() string
	EventReplyPostCommentTopic() string
	EventDeletePostTopic() string
	EventDeleteTankTopic() string

	SubscriptionFollowEndpoint() string
	SubscriptionCreateTankEndpoint() string
	SubscriptionLikeTankEndpoint() string
	SubscriptionLikePostEndpoint() string
	SubscriptionPostCommentEndpoint() string
	SubscriptionReplyPostCommentEndpoint() string
	SubscriptionDeletePostEndpoint() string
	SubscriptionDeleteTankEndpoint() string
}

type Env string

const (
	Local Env = "local"
)

func (e Env) isLocal() bool {
	return Local == e
}

type Key string

type reader struct {
	env Env
}

func (r *reader) IsLocal() bool {
	return r.env.isLocal()
}

func (r *reader) ApplicationName() string {
	return "Creatures"
}

func NewReader(env Env) Reader {
	return &reader{env}
}

func (r *reader) fromEnv(key Key) string {
	if r.env.isLocal() {
		return localValues[key]
	}
	return os.Getenv(string(key))
}
