package shared

type ImagePathProvider interface {
	BaseImageUrl() string
	BaseVideoUrl() string
	MaterialImageStoragePrefixURL() string
	AccountProfileImageStoragePrefixURL() string
	PostMediaStoragePrefixURL(isImage bool) string
	TankThumbnailStoragePrefixURL() string
}
