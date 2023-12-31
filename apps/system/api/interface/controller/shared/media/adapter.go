package media

import (
	"fmt"
	mediaPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/shared/media/v1"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/account"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/media"
	"github.com/ryo034/react-go-template/packages/go/infrastructure/cloudinary"
)

type Adapter interface {
	Adapt(meID account.ID, m *mediaPb.MediaBaseInfo) media.Media
}

type adapter struct {
	isLocal bool
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) isVideo(t mediaPb.MediaType) bool {
	switch t {
	case mediaPb.MediaType_MEDIA_TYPE_MP4:
		return true
	case mediaPb.MediaType_MEDIA_TYPE_MOV:
		return true
	}
	return false
}

func (a *adapter) AdaptVideo(meID account.ID, m *mediaPb.MediaBaseInfo) media.Media {
	path := fmt.Sprintf("/%s/%s", cloudinary.UsersPath, meID.ToString())
	if a.isLocal {
		path = fmt.Sprintf("/%s%s", cloudinary.LocalPath, path)
	}
	p, err := media.NewPath(path)
	if err != nil {
		return nil
	}
	return media.NewVideo(media.GenID(), m.GetName(), "", p, 0, 0, 0, m.GetOrder())
}

func (a *adapter) AdaptAsPhoto(meID account.ID, m *mediaPb.MediaBaseInfo) media.Media {
	path := fmt.Sprintf("/%s/%s", cloudinary.UsersPath, meID.ToString())
	if a.isLocal {
		path = fmt.Sprintf("/%s%s", cloudinary.LocalPath, path)
	}
	p, err := media.NewPath(path)
	if err != nil {
		return nil
	}
	return media.NewPhoto(media.GenID(), m.GetName(), "", p, 0, 0, m.GetOrder())
}

func (a *adapter) Adapt(meID account.ID, m *mediaPb.MediaBaseInfo) media.Media {
	if a.isVideo(m.GetMediaType()) {
		return a.AdaptVideo(meID, m)
	}
	return a.AdaptAsPhoto(meID, m)
}
