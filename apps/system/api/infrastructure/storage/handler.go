package storage

import (
	"fmt"
	"net/url"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
)

type Handler interface {
	CreateAvatarPath(hostingTo media.HostingTo, aID account.ID, id media.ID) (url.URL, error)
}

type handler struct {
	conf config.Reader
}

func NewHandler(conf config.Reader) Handler {
	return &handler{conf}
}

func (h *handler) createLocalAvatarPath(aID account.ID, id media.ID) url.URL {
	return url.URL{
		Scheme: "http",
		Host:   h.conf.MinioConfig().Host,
		Path:   fmt.Sprintf("/accounts/%s/avatar/%s", aID.ToString(), id.String()),
	}
}

func (h *handler) createMinioAvatarPath(aID account.ID, id media.ID) url.URL {
	return url.URL{
		Scheme: "https",
		Host:   h.conf.MinioConfig().Host,
		Path:   fmt.Sprintf("/accounts/%s/avatar/%s", aID.ToString(), id.String()),
	}
}

func (h *handler) createFirebaseAvatarPath(url url.URL) url.URL {
	return url
}

func (h *handler) CreateAvatarPath(hostingTo media.HostingTo, aID account.ID, id media.ID) (url.URL, error) {
	if hostingTo == media.HostingToFirebase {
		return url.URL{}, fmt.Errorf("unsupported hostingTo: %s", hostingTo)
	}

	if h.conf.IsLocal() {
		return h.createLocalAvatarPath(aID, id), nil
	}
	if hostingTo == media.HostingToR2 {
		return h.createMinioAvatarPath(aID, id), nil
	}
	return url.URL{}, fmt.Errorf("unsupported hostingTo: %s", hostingTo)
}
