package config

const (
	gcpProjectID Key = "GCP_PROJECT_ID"

	materialImageStorageSchemaHost       Key = "MATERIAL_IMAGE_STORAGE_PATH_SCHEMA_HOST"
	accountProfileImageStorageSchemaHost Key = "ACCOUNT_PROFILE_IMAGE_STORAGE_PATH_SCHEMA_HOST"
	tankImageStorageSchemaHost           Key = "TANK_IMAGE_STORAGE_PATH_SCHEMA_HOST"
	postImageStorageSchemaHost           Key = "POST_IMAGE_STORAGE_PATH_SCHEMA_HOST"

	subscriptionFollowEndpoint           Key = "SUBSCRIPTION_FOLLOW_ENDPOINT"
	subscriptionCreateTankEndpoint       Key = "SUBSCRIPTION_CREATE_TANK_ENDPOINT"
	subscriptionLikeTankEndpoint         Key = "SUBSCRIPTION_LIKE_TANK_ENDPOINT"
	subscriptionLikePostEndpoint         Key = "SUBSCRIPTION_LIKE_POST_ENDPOINT"
	subscriptionDeleteTankEndpoint       Key = "SUBSCRIPTION_DELETE_TANK_ENDPOINT"
	subscriptionPostCommentEndpoint      Key = "SUBSCRIPTION_POST_COMMENT_ENDPOINT"
	subscriptionReplyPostCommentEndpoint Key = "SUBSCRIPTION_REPLY_POST_COMMENT_ENDPOINT"
	subscriptionDeletePostEndpoint       Key = "SUBSCRIPTION_DELETE_POST_ENDPOINT"

	eventFollowSubscription           Key = "EVENT_FOLLOW_SUBSCRIPTION"
	eventCreateTankSubscription       Key = "EVENT_CREATE_TANK_SUBSCRIPTION"
	eventDeleteTankSubscription       Key = "EVENT_DELETE_TANK_SUBSCRIPTION"
	eventLikeTankSubscription         Key = "EVENT_LIKE_TANK_SUBSCRIPTION"
	eventLikePostSubscription         Key = "EVENT_LIKE_POST_SUBSCRIPTION"
	eventPostCommentSubscription      Key = "EVENT_POST_COMMENT_SUBSCRIPTION"
	eventReplyPostCommentSubscription Key = "EVENT_REPLY_POST_COMMENT_SUBSCRIPTION"
	eventDeletePostSubscription       Key = "EVENT_DELETE_POST_SUBSCRIPTION"

	eventFollowTopic           Key = "EVENT_FOLLOW_TOPIC"
	eventCreateTankTopic       Key = "EVENT_CREATE_TANK_TOPIC"
	eventLikeTankTopic         Key = "EVENT_LIKE_TANK_TOPIC"
	eventLikePostTopic         Key = "EVENT_LIKE_POST_TOPIC"
	eventDeleteTankTopic       Key = "EVENT_DELETE_TANK_TOPIC"
	eventPostCommentTopic      Key = "EVENT_POST_COMMENT_TOPIC"
	eventReplyPostCommentTopic Key = "EVENT_REPLY_POST_COMMENT_TOPIC"
	eventDeletePostTopic       Key = "EVENT_DELETE_POST_TOPIC"
)

func (r *reader) GcpProjectID() string {
	return r.fromEnv(gcpProjectID)
}

func (r *reader) MaterialImageStorageSchemaHost() string {
	return r.fromEnv(materialImageStorageSchemaHost)
}

func (r *reader) AccountProfileImageStorageSchemaHost() string {
	return r.fromEnv(accountProfileImageStorageSchemaHost)
}

func (r *reader) TankImageStorageSchemaHost() string {
	return r.fromEnv(tankImageStorageSchemaHost)
}

func (r *reader) PostImageStorageSchemaHost() string {
	return r.fromEnv(postImageStorageSchemaHost)
}

func (r *reader) EventFollowSubscription() string {
	return r.fromEnv(eventFollowSubscription)
}

func (r *reader) EventCreateTankSubscription() string {
	return r.fromEnv(eventCreateTankSubscription)
}

func (r *reader) EventLikeTankSubscription() string {
	return r.fromEnv(eventLikeTankSubscription)
}

func (r *reader) EventLikePostSubscription() string {
	return r.fromEnv(eventLikePostSubscription)
}

func (r *reader) EventPostCommentSubscription() string {
	return r.fromEnv(eventPostCommentSubscription)
}

func (r *reader) EventReplyPostCommentSubscription() string {
	return r.fromEnv(eventReplyPostCommentSubscription)
}

func (r *reader) EventDeletePostSubscription() string {
	return r.fromEnv(eventDeletePostSubscription)
}

func (r *reader) EventDeleteTankSubscription() string {
	return r.fromEnv(eventDeleteTankSubscription)
}

func (r *reader) EventFollowTopic() string {
	return r.fromEnv(eventFollowTopic)
}

func (r *reader) EventCreateTankTopic() string {
	return r.fromEnv(eventCreateTankTopic)
}

func (r *reader) EventDeleteTankTopic() string {
	return r.fromEnv(eventDeleteTankTopic)
}

func (r *reader) EventDeletePostTopic() string {
	return r.fromEnv(eventDeletePostTopic)
}

func (r *reader) EventLikeTankTopic() string {
	return r.fromEnv(eventLikeTankTopic)
}

func (r *reader) EventLikePostTopic() string {
	return r.fromEnv(eventLikePostTopic)
}

func (r *reader) EventPostCommentTopic() string {
	return r.fromEnv(eventPostCommentTopic)
}

func (r *reader) EventReplyPostCommentTopic() string {
	return r.fromEnv(eventReplyPostCommentTopic)
}

func (r *reader) SubscriptionFollowEndpoint() string {
	return r.fromEnv(subscriptionFollowEndpoint)
}

func (r *reader) SubscriptionCreateTankEndpoint() string {
	return r.fromEnv(subscriptionCreateTankEndpoint)
}

func (r *reader) SubscriptionDeleteTankEndpoint() string {
	return r.fromEnv(subscriptionDeleteTankEndpoint)
}

func (r *reader) SubscriptionDeletePostEndpoint() string {
	return r.fromEnv(subscriptionDeletePostEndpoint)
}

func (r *reader) SubscriptionLikeTankEndpoint() string {
	return r.fromEnv(subscriptionLikeTankEndpoint)
}

func (r *reader) SubscriptionLikePostEndpoint() string {
	return r.fromEnv(subscriptionLikePostEndpoint)
}

func (r *reader) SubscriptionPostCommentEndpoint() string {
	return r.fromEnv(subscriptionPostCommentEndpoint)
}

func (r *reader) SubscriptionReplyPostCommentEndpoint() string {
	return r.fromEnv(subscriptionReplyPostCommentEndpoint)
}
