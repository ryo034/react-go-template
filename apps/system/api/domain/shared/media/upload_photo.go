package media

import (
	"io"

	"github.com/google/uuid"
)

type HostingTo string

const (
	HostingToR2       HostingTo = "r2"
	HostingToFirebase HostingTo = "firebase"
)

func (h HostingTo) String() string {
	return string(h)
}

type UploadPhoto struct {
	id        ID
	content   io.Reader
	hostingTo HostingTo
}

func newUploadPhoto(id ID, content io.Reader, hostingTo HostingTo) *UploadPhoto {
	return &UploadPhoto{id, content, hostingTo}
}

func NewUploadPhotoToR2(content io.Reader) *UploadPhoto {
	nid, _ := uuid.NewV7()
	return newUploadPhoto(NewIDFromUUID(nid), content, HostingToR2)
}

func (p *UploadPhoto) ID() ID {
	return p.id
}

func (p *UploadPhoto) Content() io.Reader {
	return p.content
}

func (p *UploadPhoto) HostingTo() HostingTo {
	return p.hostingTo
}
