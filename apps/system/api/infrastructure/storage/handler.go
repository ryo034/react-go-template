package storage

import (
	"fmt"
	"net/url"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"
)

type Handler interface {
	CreateAvatarPath(hostingTo media.HostingTo, aID account.ID, id media.ID) (url.URL, error)
}

type handler struct {
	IsLocal    bool
	BucketName string
	Host       string
}

func NewHandler(isLocal bool, bucketName string, host string) Handler {
	return &handler{isLocal, bucketName, host}
}

func (h *handler) createLocalAvatarPath(aID account.ID, id media.ID) url.URL {
	// localhost can only jpg
	return url.URL{
		Scheme: "http",
		Host:   h.Host,
		Path:   fmt.Sprintf("/%s/accounts/%s/avatar/%s.jpg", h.BucketName, aID.ToString(), id.String()),
	}
}

func (h *handler) createMinioAvatarPath(aID account.ID, id media.ID) url.URL {
	return url.URL{
		Scheme: "https",
		Host:   h.Host,
		Path:   fmt.Sprintf("/%s/accounts/%s/avatar/%s.webp", h.BucketName, aID.ToString(), id.String()),
	}
}

func (h *handler) createFirebaseAvatarPath(url url.URL) url.URL {
	return url
}

func (h *handler) CreateAvatarPath(hostingTo media.HostingTo, aID account.ID, id media.ID) (url.URL, error) {
	if hostingTo == media.HostingToFirebase {
		return url.URL{}, fmt.Errorf("unsupported hostingTo: %s", hostingTo)
	}

	if h.IsLocal {
		return h.createLocalAvatarPath(aID, id), nil
	}
	if hostingTo == media.HostingToR2 {
		return h.createMinioAvatarPath(aID, id), nil
	}
	return url.URL{}, fmt.Errorf("unsupported hostingTo: %s", hostingTo)
}
