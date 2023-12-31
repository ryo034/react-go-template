package config

var localValues = map[Key]string{
	gcpProjectID: "tankmate-4fe2f",

	materialImageStorageSchemaHost:       "http://localhost:9199",
	accountProfileImageStorageSchemaHost: "http://localhost:9199",
	tankImageStorageSchemaHost:           "http://localhost:9199",
	postImageStorageSchemaHost:           "http://localhost:9199",

	eventFollowTopic:           "event-follow-topic-local",
	eventCreateTankTopic:       "event-create-tank-topic-local",
	eventLikeTankTopic:         "event-like-tank-topic-local",
	eventLikePostTopic:         "event-like-post-topic-local",
	eventDeleteTankTopic:       "event-delete-tank-topic-local",
	eventDeletePostTopic:       "event-delete-post-topic-local",
	eventPostCommentTopic:      "event-post-comment-topic-local",
	eventReplyPostCommentTopic: "event-reply-post-comment-topic-local",

	eventFollowSubscription:           "event-follow-subscription-local",
	eventCreateTankSubscription:       "event-create-tank-subscription-local",
	eventLikeTankSubscription:         "event-like-tank-subscription-local",
	eventLikePostSubscription:         "event-like-post-subscription-local",
	eventDeleteTankSubscription:       "event-delete-tank-subscription-local",
	eventDeletePostSubscription:       "event-delete-post-subscription-local",
	eventPostCommentSubscription:      "event-post-comment-subscription-local",
	eventReplyPostCommentSubscription: "event-reply-post-comment-subscription-local",

	subscriptionFollowEndpoint:           "http://host.docker.internal:19002/v1/push/action/follow",
	subscriptionCreateTankEndpoint:       "http://host.docker.internal:19002/v1/push/action/create-tank",
	subscriptionLikeTankEndpoint:         "http://host.docker.internal:19002/v1/push/action/like-tank",
	subscriptionLikePostEndpoint:         "http://host.docker.internal:19002/v1/push/action/like-post",
	subscriptionDeleteTankEndpoint:       "http://host.docker.internal:19002/v1/push/action/delete-tank",
	subscriptionDeletePostEndpoint:       "http://host.docker.internal:19002/v1/push/action/delete-post",
	subscriptionPostCommentEndpoint:      "http://host.docker.internal:19002/v1/push/action/post-comment",
	subscriptionReplyPostCommentEndpoint: "http://host.docker.internal:19002/v1/push/action/reply-post-comment",
}
