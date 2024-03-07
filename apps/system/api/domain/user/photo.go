package user

import (
	"net/url"

	"github.com/google/uuid"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"
)

type Photo struct {
	id        media.ID
	hostingTo media.HostingTo
	url       *url.URL
}

func NewPhoto(id media.ID, hostingTo media.HostingTo, url *url.URL) *Photo {
	return &Photo{id, hostingTo, url}
}

func NewPhotoFromFirebase(url *url.URL) *Photo {
	gid, _ := uuid.NewV7()
	return &Photo{media.NewIDFromUUID(gid), media.HostingToFirebase, url}
}

func (p *Photo) ID() media.ID {
	return p.id
}

func (p *Photo) HostingTo() media.HostingTo {
	return p.hostingTo
}

func (p *Photo) IsFirebase() bool {
	return p.hostingTo == media.HostingToFirebase
}

func (p *Photo) IsR2() bool {
	return p.hostingTo == media.HostingToR2
}

func (p *Photo) URL() *url.URL {
	return p.url
}

type AvatarExt string

const (
	AvatarExtJpeg AvatarExt = "jpeg"
	AvatarExtJpg  AvatarExt = "jpg"
	AvatarExtPng  AvatarExt = "png"
)
