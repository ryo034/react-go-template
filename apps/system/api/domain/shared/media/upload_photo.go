package media

import (
	"fmt"
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
	size      int64
	ext       string
	hostingTo HostingTo
}

func newUploadPhoto(id ID, content io.Reader, size int64, ext string, hostingTo HostingTo) *UploadPhoto {
	return &UploadPhoto{id, content, size, ext, hostingTo}
}

func NewUploadPhotoToR2(content io.Reader, size int64, ext string) *UploadPhoto {
	nid, _ := uuid.NewV7()
	return newUploadPhoto(NewIDFromUUID(nid), content, size, ext, HostingToR2)
}

func (p *UploadPhoto) ID() ID {
	return p.id
}

func (p *UploadPhoto) Content() io.Reader {
	return p.content
}

func (p *UploadPhoto) Size() int64 {
	return p.size
}

func (p *UploadPhoto) Ext() string {
	return p.ext
}

func (p *UploadPhoto) FileName() string {
	return fmt.Sprintf("%s%s", p.id.String(), p.ext)
}

func (p *UploadPhoto) HostingTo() HostingTo {
	return p.hostingTo
}
